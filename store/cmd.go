package store

var (
	_ CmdRes = (*StrCmdRes)(nil)
)

type Cmd interface {
	New(name string) Key
	Process(st *Store) error
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

type CmdRes interface {
	Err() error
}

type StrCmdRes struct {
	err error
	val string
}

func NewStrCmdRes(val string, err error) *StrCmdRes {
	return &StrCmdRes{
		val: val,
		err: err,
	}
}

func (scr *StrCmdRes) Err() error {
	return scr.err
}

func (scr *StrCmdRes) Val() string {
	return scr.val
}

type BoolCmdRes struct {
	err error
	val bool
}

func NewBoolCmdRes(val bool, err error) *BoolCmdRes {
	return &BoolCmdRes{
		val: val,
		err: err,
	}
}

func (scr *BoolCmdRes) Err() error {
	return scr.err
}

func (scr *BoolCmdRes) Val() bool {
	return scr.val
}
