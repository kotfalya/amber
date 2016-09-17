package key

var (
	_ Key = (*StringKey)(nil)
)

type StringKey struct {
	BaseKey
	value string
}

func NewStringKey(name, value string) *Key {
	return &StringKey{
		&BaseKey{name: name},
		value: value,
	}
}

func (sk *StringKey) Value() interface{} {
	return sk.value
}

func (sk *StringKey) SetValue(value interface{}) error {
	sk.value = value.(string)

	return
}
