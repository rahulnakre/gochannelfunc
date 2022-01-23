package apichan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

/*

	- sooo just want to pass it a url, then i can listen for info periodically from that api

	ac := ApiChan{}

	ac.AddUrl("https://jsonplaceholder.typicode.com/comments")

	- keeps polling at url at a given interval, gets the data, and also computes and stores the hash
	- now on the next polls it can just hash and check if it matches the prev hash. if it does, dont
		send on the channel cus its the same data
	-

	issues
	- rate limiting, this is kinda spammy
		- exponential backoff?
	- if the url returns a large amount of data, it can suck w constantly polling it

	- could focus only on websockets


	used like:

	ac := ApiChan{}
	ac.New()
	ac.AddUrl("https://jsonplaceholder.typicode.com/comments")

	var wg sync.WaitGroup

	go ac.Poll()

	count := 0

	for {
		msg := <-ac.Receive()
	}


*/

type ApiChan struct {
	C        chan string
	url      *url.URL
	data     string
	lasthash string
}

func (ac *ApiChan) New() {
	ac.C = make(chan string)
}

func (ac *ApiChan) Out() chan string {
	return ac.C
}

func (ac *ApiChan) AddUrl(_url string) error {
	u, err := url.Parse(_url)
	if err != nil {
		return err
	}

	ac.url = u
	return nil
}

func (ac *ApiChan) Poll(stop <-chan interface{}) {
	limiter := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-stop:
			fmt.Println("[ApiChan.Poll] stopping...")
			limiter.Stop()
			return
		default:
			<-limiter.C

			response, err := http.Get(ac.url.String())
			if err != nil {
				fmt.Println(err.Error())
				ac.C <- err.Error()
			}

			s, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				ac.C <- err.Error()
			} else {
				ac.C <- string(s)
			}
		}
	}
}
