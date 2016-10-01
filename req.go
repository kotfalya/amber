package db

const (
	RequestDBHandler = iota
	RequestNetHandler
	RequestKeyHandler
)

type Req struct {
	handler int
	cmd     string
	args    []interface{}
	res     chan Res
	stop    chan struct{}
}

func newReq(handler int, cmd string, args ...interface{}) *Req {
	req := &Req{
		handler: handler,
		cmd:     cmd,
		args:    args,
		res:     make(chan Res),
		stop:    make(chan struct{}),
	}
	go req.start()

	return req
}

func (r *Req) start() {
	<-r.stop
	close(r.res)
}

func (r *Req) Stop() {
	close(r.stop)
}

func (r *Req) Done() Res {
	defer r.Stop()
	return <-r.res
}
