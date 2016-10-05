package db

import "github.com/golang/glog"

var (
	_ Cmd = (*GetCmd)(nil)
)

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

	req := newReq(RequestKeyHandler, KeyCmdModeRead, "get", gc.keyName, level)
	db.req <- req
	gc.SetRes(req.Done())
}

func (gc *GetCmd) StrRes() *StrRes {
	return ToStrRes(gc.res)
}
