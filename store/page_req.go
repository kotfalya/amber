package store

type PageReq struct {
	name string
	args []interface{}
	res  chan PageRes
	stop chan struct{}
}

func NewPageReq(name string, args ...interface{}) *PageReq {
	return &PageReq{
		name: name,
		args: args,
		res:  make(chan PageRes),
		stop: make(chan struct{}),
	}
}

func (r *PageReq) Name() string {
	return r.name
}

func (r *PageReq) Args() []interface{} {
	return r.args
}

func (r *PageReq) start() {
	go func() {
		<-r.stop
		close(r.res)
	}()
}

func (r *PageReq) Done() PageRes {
	defer r.Stop()
	return <-r.res
}

func (r *PageReq) AddRes(res PageRes) {
	r.res <- res
}

func (r *PageReq) Stop() {
	close(r.stop)
}
