package db

import "github.com/golang/glog"

type Key interface {
	Deleted() bool
	handle(req *Req, cmd string, args ...interface{})
}

type BaseKey struct {
	deleted bool
}

func (bk *BaseKey) Deleted() bool {
	return bk.deleted
}

func KeyHandler(db *DB, req *Req) {
	mode := req.options[0].(int)
	var newKeyFunc func() Key
	if mode == KeyCmdModeUpsert {
		newKeyFunc = req.options[1].(func() Key)
	} else {
		newKeyFunc = nil
	}

	keyName := req.options[2].(string)
	level := req.options[3].(int)
	cmd := req.options[4].(string)
	args := req.options[5:]

	glog.V(2).Infof("mode: %d, level: %d, cmd: %s, keyName: %s, args: %v", mode, level, cmd, keyName, args)

	key, err := db.load(keyName, level)
	if err != nil && err.Error() != ErrUndefinedKey {
		req.res <- NewEmptyRes(err)
	} else if err != nil && mode != KeyCmdModeUpsert {
		req.res <- NewEmptyRes(err)
	} else if err != nil {
		key = newKeyFunc()
		go db.add(keyName, key, level)
	}

	key.handle(req, cmd, args...)
}
