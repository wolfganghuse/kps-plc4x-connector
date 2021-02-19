//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package knxnetip

import (
	"errors"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/transports"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

type KnxNetIpExpectation struct {
	expiration     time.Time
	acceptsMessage spi.AcceptsMessage
	handleMessage  spi.HandleMessage
	handleError    spi.HandleError
}

type KnxNetIpMessageCodec struct {
	sequenceCounter               int32
	transportInstance             transports.TransportInstance
	messageInterceptor            func(message interface{})
	defaultIncomingMessageChannel chan interface{}
	expectations                  []KnxNetIpExpectation
}

func NewKnxNetIpMessageCodec(transportInstance transports.TransportInstance, messageInterceptor func(message interface{})) *KnxNetIpMessageCodec {
	codec := &KnxNetIpMessageCodec{
		sequenceCounter:               0,
		transportInstance:             transportInstance,
		messageInterceptor:            messageInterceptor,
		defaultIncomingMessageChannel: make(chan interface{}),
		expectations:                  []KnxNetIpExpectation{},
	}
	// Start a worker that handles processing of responses
	go func() {
		for {
			work(codec)
			fmt.Printf("Restarting Work loop")
		}
	}()
	return codec
}

func (m *KnxNetIpMessageCodec) Connect() error {
	// "connect" to the remote UDP server
	return m.transportInstance.Connect()
}

func (m *KnxNetIpMessageCodec) Disconnect() error {
	return m.transportInstance.Close()
}

func (m *KnxNetIpMessageCodec) Send(message interface{}) error {
	// Cast the message to the correct type of struct
	knxMessage := model.CastKnxNetIpMessage(message)
	// Serialize the request
	wb := utils.NewWriteBuffer()
	err := knxMessage.Serialize(*wb)
	if err != nil {
		return errors.New("error serializing request " + err.Error())
	}

	// Send it to the PLC
	err = m.transportInstance.Write(wb.GetBytes())
	if err != nil {
		return errors.New("error sending request " + err.Error())
	}

	return nil
}

func (m *KnxNetIpMessageCodec) Expect(acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	expectation := KnxNetIpExpectation{
		expiration:     time.Now().Add(ttl),
		acceptsMessage: acceptsMessage,
		handleMessage:  handleMessage,
		handleError:    handleError,
	}
	m.expectations = append(m.expectations, expectation)
	return nil
}

func (m *KnxNetIpMessageCodec) SendRequest(message interface{}, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	// Send the actual message
	err := m.Send(message)
	if err != nil {
		return err
	}
	err = m.Expect(acceptsMessage, handleMessage, handleError, ttl)
	// Throttle the throughput.
	time.Sleep(time.Millisecond * 20)
	return err
}

func (m *KnxNetIpMessageCodec) GetDefaultIncomingMessageChannel() chan interface{} {
	return m.defaultIncomingMessageChannel
}

func (m *KnxNetIpMessageCodec) receive() (interface{}, error) {
	// We need at least 6 bytes in order to know how big the packet is in total
	if num, err := m.transportInstance.GetNumReadableBytes(); (err == nil) && (num >= 6) {
		data, err := m.transportInstance.PeekReadableBytes(6)
		if err != nil {
			fmt.Printf("Got error reading: %s\n", err.Error())
			// TODO: Possibly clean up ...
			return nil, nil
		}
		// Get the size of the entire packet
		packetSize := (uint32(data[4]) << 8) + uint32(data[5])
		if num >= packetSize {
			data, err = m.transportInstance.Read(packetSize)
			if err != nil {
				fmt.Printf("Got error reading: %s\n", err.Error())
				// TODO: Possibly clean up ...
				return nil, nil
			}
			rb := utils.NewReadBuffer(data)
			knxMessage, err := model.KnxNetIpMessageParse(rb)
			if err != nil {
				fmt.Printf("Got error parsing message: %s\n", err.Error())
				// TODO: Possibly clean up ...
				return nil, nil
			}
			return knxMessage, nil
		} else {
			fmt.Printf("Not enough bytes. Got: %d Need: %d\n", num, packetSize)
		}
	} else if err != nil {
		fmt.Printf("Got error reading: %s\n", err.Error())
	}
	return nil, nil
}

func work(m *KnxNetIpMessageCodec) {
	// Start an endless loop
	// TODO: Provide some means to terminate this ...
	for {
		message, err := m.receive()
		if err != nil {
			fmt.Printf("got an error reading from transport %s", err.Error())
		} else if message != nil {
			// If this message is a simple KNXNet/IP UDP Ack, ignore it for now
			tunnelingResponse := model.CastTunnelingResponse(message)
			if tunnelingResponse != nil {
				continue
			}

			// If this is an incoming tunneling request, automatically send a tunneling ACK back to the gateway
			tunnelingRequest := model.CastTunnelingRequest(message)
			if tunnelingRequest != nil {
				response := model.NewTunnelingResponse(
					model.NewTunnelingResponseDataBlock(
						tunnelingRequest.TunnelingRequestDataBlock.CommunicationChannelId,
						tunnelingRequest.TunnelingRequestDataBlock.SequenceCounter,
						model.Status_NO_ERROR),
				)
				err = m.Send(response)
				if err != nil {
					fmt.Printf("got an error sending ACK from transport %s\n", err.Error())
				}
			}

			// Handle the packet itself
			now := time.Now()
			// Give a message interceptor a chance to intercept
			if m.messageInterceptor != nil {
				m.messageInterceptor(message)
			}
			// Go through all expectations
			messageHandled := false
			for index, expectation := range m.expectations {
				// Check if this expectation has expired.
				if now.After(expectation.expiration) {
					// Remove this expectation from the list.
					m.expectations = append(m.expectations[:index], m.expectations[index+1:]...)
					// Call the error handler.
					err := expectation.handleError(errors.New("expectation timed out"))
					if err != nil {
						log.Errorf("Got an error handling error on expectation: %s", err.Error())
					}
					continue
				}

				// Check if the current message matches the expectations
				// If it does, let it handle the message.
				if accepts := expectation.acceptsMessage(message); accepts {
					err = expectation.handleMessage(message)
					if err == nil {
						messageHandled = true
						// If this is the last element of the list remove it differently than if it's before that
						if (index + 1) == len(m.expectations) {
							m.expectations = m.expectations[:index]
						} else if (index + 1) < len(m.expectations) {
							m.expectations = append(m.expectations[:index], m.expectations[index+1:]...)
						}
					}
				}
			}

			// If the message has not been handled and a default handler is provided, call this ...
			if !messageHandled {
				m.defaultIncomingMessageChannel <- message
			}
		} else {
			// Sleep for 10ms
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (m KnxNetIpMessageCodec) GetTransportInstance() transports.TransportInstance {
	return m.transportInstance
}
