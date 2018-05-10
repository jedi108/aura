package incrementor

import (
	"sync"
	"testing"
)

func TestGetNumberInitMin(t *testing.T) {
	i := NewIncrementor()
	if i.GetNumber() != 0 {
		t.Fatal("min value not equeal")
	}
}

func TestIncrement(t *testing.T) {
	i := NewIncrementor()
	i.IncrementNumber()
	i.IncrementNumber()
	if i.GetNumber() != 2 {
		t.Fatalf("error increment number, not equal 2 != %v", i.GetNumber())
	}
}

func TestSetMaxNumber(t *testing.T) {
	i := NewIncrementor()
	i.SetMaximumValue(1000)
	if i.maxNumber != 1000 {
		t.Fatalf("error in max number, not equal 1000 != %v", i.maxNumber)
	}
}

func TestSetMaxNumberRoundZero(t *testing.T) {
	i := NewIncrementor()
	i.SetMaximumValue(65534)
	i.currentNubmer = 65535
	i.IncrementNumber()
	if i.currentNubmer != 0 {
		t.Fatalf("error increment number, not equal 0 != %v", i.GetNumber())
	}
}

func TestManyGoroutinesIncrement(t *testing.T) {
	countIterations := 32768
	var wg sync.WaitGroup
	wg.Add(countIterations)
	incr := NewIncrementor()
	for i := 0; i < countIterations; i++ {
		go func(i int) {
			incr.IncrementNumber()
			wg.Done()
		}(i)
	}
	wg.Wait()
	if incr.GetNumber() != uint32(countIterations) {
		t.Fatalf("error nubmer, not equal 32768 != %v", incr.GetNumber())
	}
}
