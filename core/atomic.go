package core

import "sync/atomic"

type atomicBool int32

func (b *atomicBool) isSet() bool { return atomic.LoadInt32((*int32)(b)) != 0 }
func (b *atomicBool) setTrue()    { atomic.StoreInt32((*int32)(b), 1) }

type AtomicInt64 int64

func (b *AtomicInt64) AddAndGet(v int64) int64 {

	return atomic.AddInt64((*int64)(b), v)
}
