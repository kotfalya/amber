package db

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

	if keyRes, err := GetReq(db, gc.keyName, level); err != nil {
		gc.SetRes(NewStrCmdRes("", err))
	} else {
		strKey, err := ToStrKey(keyRes.Val())
		gc.SetRes(NewStrCmdRes(strKey.StrVal(), err))
	}
}
