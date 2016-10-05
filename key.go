package db

import "github.com/golang/glog"

type Key interface {
	Deleted() bool
}

type BaseKey struct {
	deleted bool
}

func (bk *BaseKey) Deleted() bool {
	return bk.deleted
}

func KeyHandler(db *DB, req *Req) {
	glog.V(2).Infof("req:  %s, args: %v", req.cmd, req.args)

	key := NewStrKey()
	key.SetVal("hihi")

	req.res <- NewStrRes(key.StrVal(), nil)
}
