package gochannelfunc_test

import (
	"fmt"
	"testing"
	"github.com/gochannelfunc"
)

func TestBasicHc(t *testing.T) {
	timesTwo := func(x int) int {
		return x * 2;
	}

	hc := gochannelfunc.HandlerChannel[int]{}
	hc.New()
	hc.AddFunc(timesTwo)

	go func() {
		hc.Send(2) 
		hc.Send(3) 
	}()

	assertEqual(t, hc.Receive(), 4, "")
	assertEqual(t, hc.Receive(), 6, "")
}

func TestBasicHc2Funcs(t *testing.T) {
	timesTwo := func(x int) int {
		return x * 2;
	}

	timesThree := func(x int) int {
		return x * 3;
	}

	hc := gochannelfunc.HandlerChannel[int]{}
	hc.New()
	hc.AddFunc(timesTwo)
	hc.AddFunc(timesThree)

	go func() {
		hc.Send(2) 
		hc.Send(3) 
	}()

	assertEqual(t, hc.Receive(), 12, "")
	assertEqual(t, hc.Receive(), 18, "")
}