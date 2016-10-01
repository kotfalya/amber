package db

import (
	"errors"
	"github.com/golang/glog"
	"github.com/kotfalya/store/utils"
)

type DB struct {
	config *Config
	data   map[string]Key
	req    chan *Req
	stop   chan struct{}
}

func NewDB(config *Config) *DB {
	db := &DB{
		config: config,
		data:   make(map[string]Key),
		req:    make(chan *Req, 10),
		stop:   make(chan struct{}),
	}
	go db.start()

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
				switch req.handler {
				case RequestKeyHandler:
					db.keyHandle(req)
				case RequestDBHandler:
					db.dbHandle(req)
				case RequestNetHandler:
					db.netHandle(req)
				default:
					panic(errors.New(ErrInvalidReqHandler))
				}

			}(req)

		case <-db.stop:
			glog.V(1).Infoln("db:stop")
			return
		}
	}
}

func (db *DB) dbHandle(req *Req) {

}

func (db *DB) netHandle(req *Req) {

}

func (db *DB) keyHandle(req *Req) {
	glog.V(2).Infof("req:  %s, args: %v", req.cmd, req.args)

	key := NewStrKey()
	key.SetVal("hihi")

	req.res <- NewKeyRes(key, nil)
}
