package db

type Cmd interface {
	Process(db *DB)
	Res() CmdRes
	SetRes(res CmdRes)
}

type BaseCmd struct {
	res CmdRes
}

func (bc *BaseCmd) Res() CmdRes {
	return bc.res
}

func (bc *BaseCmd) SetRes(res CmdRes) {
	bc.res = res
}
