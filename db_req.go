package db

type Req struct {
	name string
	args []interface{}
	res  chan Res
	stop chan struct{}
}

func newReq(name string, args ...interface{}) *Req {
	return &Req{
		name: name,
		args: args,
		res:  make(chan Res),
		stop: make(chan struct{}),
	}
}

func (r *Req) start() {
	go func() {
		<-r.stop
		close(r.res)
	}()
}

func (r *Req) Stop() {
	close(r.stop)
}

func (r *Req) Done() Res {
	defer r.Stop()
	return <-r.res
}

func NewAddReq(keyName string) *Req {
	return newReq("get", keyName)
}
