package db

import "errors"

type GetCmd struct {
	*BaseCmd
	keyName string
	level   int
}

func NewGetCmd(keyName string, options ...string) *GetCmd {
	return &GetCmd{
		&BaseCmd{},
		keyName,
		ParseReadLevel(options),
	}
}

func (gc *GetCmd) Process(db *DB) {
	var level int
	if gc.level > -1 {
		level = gc.level
	} else {
		level = db.config.readLevel
	}

	req := NewAddReq(gc.keyName, level)
	db.req <- req
	res := req.Done()

	switch r := res.(type) {
	case *StrRes:
		gc.SetRes(NewStrCmdRes(r.Val(), nil))
	default:
		gc.SetRes(NewStrCmdRes("", errors.New(ErrInvalidResType)))
	}
}
