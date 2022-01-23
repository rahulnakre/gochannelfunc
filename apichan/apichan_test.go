package apichan_test

import (
	"testing"
	"time"

	"github.com/gochannelfunc/apichan"
)

func TestBasicHc(t *testing.T) {
	ac := apichan.ApiChan{}
	ac.New()
	ac.AddUrl("https://jsonplaceholder.typicode.com/comments/1")

	stop := make(chan interface{})
	timer := time.NewTimer(2 * time.Second)

	go ac.Poll(stop, 300*time.Millisecond)

	for {
		select {
		case <-timer.C:
			t.Log("sending stop signal")
			close(stop)
			return
		case msg := <-ac.Out():
			t.Logf("msg:%s\n", msg)
		default:
		}
	}
}
