package apichan_test

import (
	"fmt"
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

	go ac.Poll(stop)

	for {
		select {
		case <-timer.C:
			fmt.Println("sending stop signal")
			// fmt.Println(len(ac.C))
			// stop <- 1
			close(stop)
			// break
			return
		case msg := <-ac.Out():
			fmt.Printf("msg:%s\n", msg)
		default:
		}
	}

	msg := <-ac.C

	fmt.Println(msg)

	// for range <-ac.Out() {
	// fmt.Println("sf")
	// }

}
