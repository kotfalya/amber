package db

import (
	"errors"
	"github.com/golang/glog"
	"github.com/kotfalya/store/utils"
)

type DB struct {
	config *Config
	name   string
	data   map[string]Key
	req    chan *Req
	stop   chan struct{}
}

func NewDB(name string, config *Config) *DB {
	db := &DB{
		config: config,
		name:   name,
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
			req.master = db.name
			sem.Acquire()
			go func(req *Req) {
				defer sem.Release()
				switch req.handler {
				case RequestDBHandler:
					DBHandle(db, req)
				case RequestKeyHandler:
					KeyHandler(db, req)
				case RequestNetHandler:
					NetHandler(db, req)
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

func (db *DB) load(keyName string, level int) (Key, error) {
	if key, ok := db.data[keyName]; ok {
		return key, nil
	} else {
		return nil, errors.New(ErrUndefinedKey)
	}
}

func (db *DB) add(keyName string, key Key, level int) error {
	db.data[keyName] = key

	return nil
}

func DBHandle(db *DB, req *Req) {

}
