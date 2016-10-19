package db

import "errors"

var (
	_ Key = (*StrKey)(nil)
)

type StrKey struct {
	BaseKey
	val string
}

func NewStrKey(master string) *StrKey {
	return &StrKey{
		*NewBaseKey(master),
		"",
	}
}

func (sk *StrKey) handle(req *Req, cmd string, args ...interface{}) {
	var res Res

	switch cmd {
	case "get":
		res = NewStrRes(sk.StrVal(), nil)
	case "set":
		err := sk.SetVal(args[0])
		res = NewBoolRes(err == nil, err)
	default:
		res = NewEmptyRes(errors.New(ErrUndefinedKeyCmd))
	}

	req.res <- res
}

func (sk *StrKey) Val() interface{} {
	return sk.val
}

func (sk *StrKey) SetVal(val interface{}) error {
	sk.val = val.(string)

	return nil
}

func (sk *StrKey) StrVal() string {
	return sk.val
}
