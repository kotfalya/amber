package types

type Key interface {
	Name() string
	Value() interface{}
	SetValue(interface{}) error
}

type BaseKey struct {
	name string
}

func (bk *BaseKey) Name() string {
	return bk.name
}
