package gochannelfunc_test

import (
	"fmt"
	"testing"
	"github.com/gochannelfunc"
)

// func Test(t *testing.T) {
// 	want := "done working"
// 	got := Work()

// 	if strings.Compare(got, want) != 0 {
// 		t.Errorf("Want: %s . Got: %s", want, got)
// 	}
// }

// func TestEvenIter(t *testing.T) {
// 	iter := gochannelfunc.EvenIter{Max: 30, CurrValue: 0}
// 	i := 2

// 	for iter.Next() {
// 		assertEqual(t, i, iter.Value(), "not the same value")
// 		i += 2
// 	}
// }

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("\nExpected: %v\nGot: %v", a, b)
	}
	t.Fatal(message)
}

func TestBasicHc(t *testing.T) {
	timesTwo := func(x int) int {
		return x * 2;
	}

	hc := gochannelfunc.HandlerChannel[int]{}
	hc.New()
	hc.AddFunc(timesTwo)
	// hc.AddFunc(timesTwo)

	go func() {
		// hc.In() <- 2
		hc.Send(2) 
		hc.Send(3) 
	}()


	// t.Logf("x1: %d\n", <-hc.Out)
	// t.Logf("x2: %d\n", <-hc.Out)


	assertEqual(t, <-hc.Out(), 4, "")
	assertEqual(t, <-hc.Out(), 6, "")


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