package db

import "github.com/golang/glog"

type GetCmd struct {
	*BaseCmd
	keyName string
	level   int
}

func NewGetCmd(keyName string, options ...string) *GetCmd {
	return &GetCmd{
		&BaseCmd{},
		keyName,
		parseLevel(options),
	}
}

func (gc *GetCmd) Process(db *DB) {
	level, err := readOption(gc.level, db.config.readLevel)
	if err != nil {
		glog.Errorln(err)
	}

	req := newReq(RequestKeyHandler, "get", gc.keyName, level)
	db.req <- req

	if keyRes, err := ToKeyRes(req.Done()); err != nil {
		gc.SetRes(NewStrCmdRes("", err))
	} else {
		strKey, err := ToStrKey(keyRes.Val())
		gc.SetRes(NewStrCmdRes(strKey.StrVal(), err))
	}
}
