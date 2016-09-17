package key

type Key interface {
	Name() interface{}
	Value() interface{}
	SetValue(interface{}) error
}

type BaseKey struct {
	name string
}

func (bk *BaseKey) Name() string {
	return bk.name
}
