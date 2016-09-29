package db

type Req struct {
	name string
	args []interface{}
	res  chan Res
	stop chan struct{}
}

func newReq(name string, args ...interface{}) *Req {
	req := &Req{
		name: name,
		args: args,
		res:  make(chan Res),
		stop: make(chan struct{}),
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

func NewAddReq(keyName string, readLevel int) *Req {
	return newReq("get", keyName, readLevel)
}
