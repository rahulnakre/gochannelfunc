package gochannelfunc

import (
	"errors"
	"fmt"
	"reflect"
)

// type HandlerChannel struct {
// 	c chan string
// 	f func(string) string
// }

type HandlerChannel struct {
	c chan interface{}
	f func(interface{}) interface{}
}

// func (hc *HandlerChannel) New(c chan interface{}, f func(interface{}) interface{}) HandlerChannel {

func (hc *HandlerChannel) New(c interface{}, f interface{}) (HandlerChannel, error) {

	// types.Identical(c.)

	isChannel := reflect.ValueOf(c).Kind() == reflect.Chan
	if !isChannel {
		msg := fmt.Sprintf("Expected c to be of type: chan T\nGot: %s\n", reflect.TypeOf(c))
		return HandlerChannel{}, errors.New(msg)
	}

	channelDataType := reflect.ValueOf(c).Type().Elem()
	// fDataType := reflect.StructField

	chanType := reflect.TypeOf(c)
	funcType := reflect.TypeOf(f)

	fmt.Printf("isChannel: %v\n", isChannel)
	fmt.Printf("chanType: %v\n", chanType)
	fmt.Printf("inner data of chan Type: %v\n", channelDataType)
	fmt.Printf("funcType: %v\n", funcType)
	// fmt.Printf("fDataType: %v\n", fDataType)

	return HandlerChannel{}, nil
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
*/

func Work() string {
	fmt.Println("working...")

	// hc := HandlerChannel{
	// c: make(chan string),
	// }

	// fmt.Printf("%v\n", hc)
	// c := make(chan string)

	hc := HandlerChannel{}
	c := make(chan int)
	f := func(i int) int {
		return 1
	}

	hc.New(c, f)

	hel(make(chan string))

	// reflect.FuncOf([]reflect.Type{reflect.Int}, []reflect.Type{reflect.Bool}, false)

	return "done working"
}

func hel(i interface{}) bool {

	fmt.Println(i)

	return true
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
