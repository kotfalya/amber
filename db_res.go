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

type KeyRes struct {
	*EmptyRes
	val Key
}

func NewKeyRes(val Key, err error) *KeyRes {
	return &KeyRes{
		&EmptyRes{err: err},
		val,
	}
}

func (kr *KeyRes) Val() Key {
	return kr.val
}

func ToKeyRes(res Res) (*KeyRes, error) {
	switch kr := res.(type) {
	case *KeyRes:
		return kr, nil
	default:
		return nil, errors.New(ErrInvalidResType)
	}
}
