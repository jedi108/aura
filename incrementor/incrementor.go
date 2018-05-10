//Package incrementor provides primitives for increment unsigned integer value with concurrency
package incrementor

import (
	"sync/atomic"
)

type Incrementor struct {
	maxNumber     uint32
	currentNubmer uint32
}

// New object-value of Incrementor
func NewIncrementor() *Incrementor {
	return &Incrementor{maxNumber: 1<<32 - 1}
}

// Getter number
func (incr *Incrementor) GetNumber() uint32 {
	return atomic.LoadUint32(&incr.currentNubmer)
}

// Increment number delta of 1
//  if number > maxNumber then number set 0
func (incr *Incrementor) IncrementNumber() {
	if atomic.LoadUint32(&incr.currentNubmer) < atomic.LoadUint32(&incr.maxNumber) {
		atomic.AddUint32(&incr.currentNubmer, 1)
	} else {
		atomic.StoreUint32(&incr.currentNubmer, 0)
	}
}

// Set max number
// 	Default max nubmer is 4294967295 (uint32)
func (incr *Incrementor) SetMaximumValue(maxNum uint32) {
	atomic.StoreUint32(&incr.maxNumber, maxNum)
	if atomic.LoadUint32(&incr.currentNubmer) > maxNum {
		atomic.StoreUint32(&incr.currentNubmer, 0)
	}
}
