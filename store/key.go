package store

var (
	_ CmdRes = (*StrCmdRes)(nil)
)

type Key interface {
	Name() string
	Enabled() bool
	Value() interface{}
	SetValue(interface{}) error
}

type BaseKey struct {
	enabled bool
	name    string
}

func (bk *BaseKey) Name() string {
	return bk.name
}

func (bk *BaseKey) Enabled() bool {
	return bk.enabled
}

type Cmd interface {
	Process(st *Store)
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
