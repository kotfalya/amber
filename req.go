package db

import "errors"

const (
	RequestDBHandler = iota
	RequestNetHandler
	RequestKeyHandler
)

type Req struct {
	master  string
	handler int
	options []interface{}
	res     chan Res
	stop    chan bool
}

func newReq(handler int, options ...interface{}) *Req {
	req := &Req{
		handler: handler,
		options: options,
		res:     make(chan Res),
		stop:    make(chan bool),
	}
	go req.start()

	return req
}

func (r *Req) start() {
	interrupt := <-r.stop
	close(r.stop)
	if interrupt {
		r.res <- NewStopRes(errors.New(ErrStopRequest))
	}
	close(r.res)
}

func (r *Req) Stop() {
	r.stop <- true
}

func (r *Req) Done() Res {
	defer func() {
		r.stop <- false
	}()
	return <-r.res
}
