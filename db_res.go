package db

type Res interface {
	Err() error
}

type EmptyRes struct {
	err error
}

func (er *EmptyRes) Err() error {
	return er.err
}

func NewEmptyRes(err error) *EmptyRes {
	return &EmptyRes{err: err}
}

type StrRes struct {
	*EmptyRes
	val string
}

func NewStrRes(val string, err error) *StrRes {
	return &StrRes{
		&EmptyRes{err: err},
		val,
	}
}

func (sr *StrRes) Val() string {
	return sr.val
}
