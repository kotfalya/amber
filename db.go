package amber

import (
	"errors"

	"github.com/golang/glog"
	"github.com/kotfalya/amber/utils"
)

type DB struct {
	config   *Config
	name     string
	rootPage *Page
	req      chan *Req
	stop     chan struct{}
}

func NewDB(name string, config *Config) *DB {
	db := &DB{
		config:   config,
		name:     name,
		rootPage: createRootPage(),
		req:      make(chan *Req, 10),
		stop:     make(chan struct{}),
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
	key, err := db.rootPage.load(keyName)
	if err != nil {
		return nil, errors.New(ErrUndefinedKey)
	}
	return key, nil
}

func (db *DB) add(keyName string, key Key, level int) (err error) {
	err = db.rootPage.add(keyName, key)

	return
}

func DBHandle(db *DB, req *Req) {

}
