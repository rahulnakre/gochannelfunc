package gochannelfunc

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	want := "done working"
	got := Work()

	if strings.Compare(got, want) != 0 {
		t.Errorf("Want: %s . Got: %s", want, got)
	}
}

func TestEvenIter(t *testing.T) {
	iter := EvenIter{max: 30, currValue: 0}
	i := 2

	for iter.Next() {
		assertEqual(t, i, iter.Value(), "not the same value")
		i += 2
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
