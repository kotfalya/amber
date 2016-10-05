package db

var (
	_ Key = (*StrKey)(nil)
)

type StrKey struct {
	BaseKey
	val string
}

func NewStrKey() *StrKey {
	return &StrKey{
		BaseKey{},
		"",
	}
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
