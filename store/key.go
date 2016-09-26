package store

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
