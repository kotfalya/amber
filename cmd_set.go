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

func (sc *SetCmd) BoolRes() *BoolRes {
	return ToBoolRes(sc.res)
}

func (sc *SetCmd) Process(db *DB) {
	level, err := readOption(sc.level, db.config.readLevel)
	if err != nil {
		glog.Errorln(err)
	}
	persist, err := readOption(sc.persist, db.config.persist)
	if err != nil {
		glog.Errorln(err)
	}

	newKeyFunc := func(master string) Key {
		return NewStrKey(master)
	}

	req := newReq(RequestKeyHandler, KeyCmdModeUpsert, newKeyFunc, sc.keyName, level, "set", sc.value, persist)
	db.req <- req
	sc.SetRes(req.Done())
}
