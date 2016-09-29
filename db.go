package db

import (
	"github.com/golang/glog"
	"github.com/kotfalya/store/utils"
)

type DB struct {
	config *Config
	data   map[string]string
	req    chan *Req
	stop   chan struct{}
}

func NewDB(config *Config) *DB {
	db := &DB{
		config: config,
		data:   make(map[string]string),
		req:    make(chan *Req, 10),
		stop:   make(chan struct{}),
	}

	return db
}

func (db *DB) start() {
	sem := utils.NewSemaphore(10)
	for {
		select {
		case req := <-db.req:
			sem.Acquire()
			go func(req *Req) {
				defer sem.Release()

				db.handle(req)
			}(req)

		case <-db.stop:
			glog.V(1).Infoln("db:stop")
			return
		}
	}
}

func (db *DB) handle(req *Req) {
	glog.V(2).Infof("req:  %s, args: %v", req.name, req.args)

	req.res <- NewEmptyRes(nil)
}
