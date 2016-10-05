package db

import "github.com/golang/glog"

var (
	_ Cmd = (*SetCmd)(nil)
)

type SetCmd struct {
	*BaseCmd
	keyName string
	value   string
	level   int
	persist int
}

func NewSetCmd(keyName string, value string, options ...string) *SetCmd {
	return &SetCmd{
		&BaseCmd{},
		keyName,
		value,
		parseLevel(options),
		parsePersist(options),
	}
}

func (gc *SetCmd) BoolRes() *BoolRes {
	return ToBoolRes(gc.res)
}

func (gc *SetCmd) Process(db *DB) {
	level, err := readOption(gc.level, db.config.readLevel)
	if err != nil {
		glog.Errorln(err)
	}
	persist, err := readOption(gc.persist, db.config.persist)
	if err != nil {
		glog.Errorln(err)
	}

	newKeyFunc := func() Key {
		return NewStrKey()
	}

	req := newReq(RequestKeyHandler, KeyCmdModeUpsert, newKeyFunc, gc.keyName, level, "set", gc.value, persist)
	db.req <- req
	gc.SetRes(req.Done())
}
