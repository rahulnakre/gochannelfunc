package gochannelfunc

type HandlerChannel[K any] struct {
	c chan K
	funcs []func(K) K
}

func (hc *HandlerChannel[K]) New() {
	hc.c = make(chan K)
}

func (hc *HandlerChannel[K]) NewBuffered(bufsize int) {
	hc.c = make(chan K, bufsize)
}

func (hc *HandlerChannel[K]) In() chan<- K {
	return hc.c
}

func (hc *HandlerChannel[K]) Out() <-chan K {
	return hc.c
}

func (hc *HandlerChannel[K]) AddFunc(f func(K) K) error {
	hc.funcs = append(hc.funcs, f)

	return nil
}

func (hc *HandlerChannel[K]) Send(x K) {
	for _, f := range hc.funcs {
		x = f(x)			
	}

	hc.c <- x	
}

func (hc *HandlerChannel[K]) Receive() K {
	return <-hc.c			
}