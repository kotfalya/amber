package db

type Key interface {
	Deleted() bool
}

type BaseKey struct {
	deleted bool
}

func (bk *BaseKey) Deleted() bool {
	return bk.deleted
}
