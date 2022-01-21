package gochannelfunc

import (
	"fmt"
)


type HandlerChannel[K any] struct {
	C chan K
}

// func (hc HandlerChannel[K]) New(c chan K) (HandlerChannel, error) {
// 	return HandlerChannel[K]{ c: c }, nil
// }

func (hc HandlerChannel[K]) New() HandlerChannel[K] {
	return HandlerChannel[K]{ C: make(chan K) }
}

/*
	Reqs:
		Inputs:
			- channel type T
			- function with sig: func (x: T) T { }

		Behavior:
			- TLDR: data is passed into this custom channel, and at somepoint
				before the receiver consumes the data, the function is applied
				to the data and the modified version comes out
			- should probably have an option to do eager vs lazy
			- receiver gets the customized data
			- this is like an iterator tbh

		Details:
			- to get a generic type T, we can use interface{}, which'll accent any type
				- An empty interface may hold values of any type. (Every type implements at least zero methods.)
			- probably have to use the types library to do type assertion stuff

		Challenges:
			- how do i get the channel type, function input + output to be the same type?
				- types lib?


		usage

		isEven := func(x int) bool {
			if x % 2 == 0 {
				return true;
			}
			return false;
		}

		timesTwo := func(x int) int {
			return x * 2;
		}

		divisibleByfour := func(b bool) bool {
			if x % 4 == 0 {
				return true;
			}

			return false;
		} 

		hc := HandlerChannel.New(int)

		hc.AddFunc(isEven)
		hc.AddFunc(timesTwo, divisibleBy4) // doesnt perform this step if second func result is falsey

		hc.c <- 2
		
		fmt.Printf("result: %d\n", <-hc.c) // should print 4
*/

func Work() string {
	fmt.Println("working...")

	c := make(chan int)
	hc := HandlerChannel[int]{}


	// isEven := func(x int) bool {
	// 	if x % 2 == 0 {
	// 		return true;
	// 	}
	// 	return false;
	// }

	// timesTwo := func(x int) int {
	// 	return x * 2;
	// }

	// divisibleByfour := func(x int) bool {
	// 	if x % 4 == 0 {
	// 		return true;
	// 	}

	// 	return false;
	// } 


	return "done working"
}


type EvenIter struct {
	max       int
	currValue int
}

func (e *EvenIter) Value() int {
	return e.currValue
}

func (e *EvenIter) Next() bool {
	newValue := e.currValue + 2

	if newValue > e.max {
		return false
	}

	e.currValue = newValue
	return true
}
