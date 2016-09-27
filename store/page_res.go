package store

type PageRes interface {
	Err() error
}

type EmptyPageRes struct {
	err error
}

func (epr *EmptyPageRes) Err() error {
	return epr.err
}

func NewEmptyPageRes(err error) *EmptyPageRes {
	return &EmptyPageRes{err: err}
}

type StrPageRes struct {
	*EmptyPageRes
	val string
}

func NewStrPageRes(val string, err error) *StrPageRes {
	return &StrPageRes{
		&EmptyPageRes{err: err},
		val,
	}
}

func (spr *StrPageRes) Val() string {
	return spr.val
}

type KeyPageRes struct {
	*EmptyPageRes
	val Key
}

func NewKeyPageRes(val Key, err error) *KeyPageRes {
	return &KeyPageRes{
		&EmptyPageRes{err: err},
		val,
	}
}

func (kpr *KeyPageRes) Val() Key {
	return kpr.val
}
