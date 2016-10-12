package db

import "errors"

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

type StopRes struct {
	*EmptyRes
}

func NewStopRes(err error) *StopRes {
	return &StopRes{&EmptyRes{err: err}}
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

func ToStrRes(res Res) *StrRes {
	switch sr := res.(type) {
	case *StrRes:
		return sr
	case *StopRes:
		return NewStrRes("", sr.err)
	case *EmptyRes:
		return NewStrRes("", sr.err)
	default:
		return NewStrRes("", errors.New(ErrInvalidResType))
	}
}

type BoolRes struct {
	*EmptyRes
	val bool
}

func NewBoolRes(val bool, err error) *BoolRes {
	return &BoolRes{
		&EmptyRes{err: err},
		val,
	}
}

func (sr *BoolRes) Val() bool {
	return sr.val
}

func ToBoolRes(res Res) *BoolRes {
	switch sr := res.(type) {
	case *BoolRes:
		return sr
	case *StopRes:
		return NewBoolRes(false, sr.err)
	case *EmptyRes:
		return NewBoolRes(false, sr.err)
	default:
		return NewBoolRes(false, errors.New(ErrInvalidResType))
	}
}
