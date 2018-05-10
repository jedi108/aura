package main

import (
	"aura/incrementor"
	"fmt"
)

func main() {
	i := 0
	countIterations := 10
	chanWait := make(chan struct{})
	incr := incrementor.NewIncrementor()
	go func() {
		for {
			i++
			go func(i int) {
				incr.IncrementNumber()
				if i >= countIterations {
					chanWait <- struct{}{}
				}
			}(i)
		}
	}()
	<-chanWait
	fmt.Println(incr.GetNumber())
}
