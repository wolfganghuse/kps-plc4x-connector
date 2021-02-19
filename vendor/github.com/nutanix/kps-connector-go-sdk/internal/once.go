package internal

import (
	"sync"
	"sync/atomic"
)

// Once is an object that will perform exactly one action unless the action fails.
type Once struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/x86),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    sync.Mutex
}

// TryDo calls the function f if and only if TryDo is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.TryDo(f) is called multiple times, only the first call will invoke f,
// unless the invocation results in an error A new instance of
// Once is required for each function to execute.
//
// TryDo is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.TryDo(func() error { return config.init(filename) })
//
// Because no call to TryDo returns until the one call to f returns, if f causes
// TryDo to be called, it will deadlock.
//
// If f panics, TryDo considers it to have failed
//
func (o *Once) TryDo(f func() error) error {
	// Note: Here is an incorrect implementation of Do:
	//
	//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//		f()
	//	}
	//
	// TryDo guarantees that when it returns, f has finished.
	// This implementation would not implement that guarantee:
	// given two simultaneous calls, the winner of the cas would
	// call f, and the second would return immediately, without
	// waiting for the first's call to f to complete.
	// This is why the slow path falls back to a mutex, and why
	// the atomic.StoreUint32 must be delayed until after f returns.

	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		return o.doSlow(f)
	}

	return nil
}

func (o *Once) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		if err := f(); err != nil {
			return err
		}
		atomic.StoreUint32(&o.done, 1)
	}
	return nil
}
