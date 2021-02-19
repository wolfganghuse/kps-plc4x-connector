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
package interceptors

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/model"
	apiModel "github.com/apache/plc4x/plc4go/pkg/plc4go/model"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/values"
)

type SingleItemRequestInterceptor struct {
}

func NewSingleItemRequestInterceptor() SingleItemRequestInterceptor {
	return SingleItemRequestInterceptor{}
}

func (m SingleItemRequestInterceptor) InterceptReadRequest(readRequest apiModel.PlcReadRequest) []apiModel.PlcReadRequest {
	// If this request just has one field, go the shortcut
	if len(readRequest.GetFieldNames()) == 1 {
		return []apiModel.PlcReadRequest{readRequest}
	}
	// In all other cases, create a new read request containing only one item
	defaultReadRequest := readRequest.(model.DefaultPlcReadRequest)
	var readRequests []apiModel.PlcReadRequest
	for _, fieldName := range readRequest.GetFieldNames() {
		field := readRequest.GetField(fieldName)
		subReadRequest := model.NewDefaultPlcReadRequest(
			map[string]apiModel.PlcField{fieldName: field},
			defaultReadRequest.Reader,
			defaultReadRequest.ReadRequestInterceptor)
		readRequests = append(readRequests, subReadRequest)
	}
	return readRequests
}

func (m SingleItemRequestInterceptor) ProcessReadResponses(readRequest apiModel.PlcReadRequest, readResults []apiModel.PlcReadRequestResult) apiModel.PlcReadRequestResult {
	if len(readResults) == 1 {
		return readResults[0]
	}

	responseCodes := map[string]apiModel.PlcResponseCode{}
	val := map[string]values.PlcValue{}
	var err error
	for _, readResult := range readResults {
		if readResult.Err != nil {
			// TODO: Handle the case that multiple requests had errors that are different
			err = readResult.Err
		} else if readResult.Response != nil {
			for _, fieldName := range readResult.Response.GetRequest().GetFieldNames() {
				responseCodes[fieldName] = readResult.Response.GetResponseCode(fieldName)
				val[fieldName] = readResult.Response.GetValue(fieldName)
			}
		}
	}
	return apiModel.PlcReadRequestResult{
		Request:  readRequest,
		Response: model.NewDefaultPlcReadResponse(readRequest, responseCodes, val),
		Err:      err,
	}
}

func (m SingleItemRequestInterceptor) InterceptWriteRequest(writeRequest apiModel.PlcWriteRequest) []apiModel.PlcWriteRequest {
	return []apiModel.PlcWriteRequest{writeRequest}
}

func (m SingleItemRequestInterceptor) ProcessWriteResponses(writeRequest apiModel.PlcWriteRequest, writeResponses []apiModel.PlcWriteRequestResult) apiModel.PlcWriteRequestResult {
	return apiModel.PlcWriteRequestResult{}
}
