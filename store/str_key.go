package store

import (
	"errors"
)

var (
	_ Key = (*StringKey)(nil)
)

type StringKey struct {
	BaseKey
	value string
}

func NewStringKey(name string) *StringKey {
	return &StringKey{
		BaseKey{name: name},
		"",
	}
}

func (sk *StringKey) Value() interface{} {
	return sk.value
}

func (sk *StringKey) SetValue(value interface{}) error {
	sk.value = value.(string)

	return nil
}

func (sk *StringKey) StrValue() string {
	return sk.value
}

type GetCmd struct {
	*BaseCmd
	keyName string
}

func NewGetCmd(keyName string) *GetCmd {
	cmd := &GetCmd{
		&BaseCmd{},
		keyName,
	}

	return cmd
}

func (gc *GetCmd) Process(st *Store) error {
	var res *StrCmdRes

	if key, err := st.FindKey(gc.keyName); err != nil {
		res = NewStrCmdRes("", err)
	} else {
		switch bestKey := key.(type) {
		case *StringKey:
			res = NewStrCmdRes(bestKey.StrValue(), nil)
		default:
			res = NewStrCmdRes("", errors.New(ErrInvalidKeyType))
		}
	}

	gc.SetRes(res)

	return nil
}

type SetCmd struct {
	*BaseCmd
	keyName string
	val     string
}

func NewSetCmd(keyName string, val string) *SetCmd {
	cmd := &SetCmd{
		&BaseCmd{},
		keyName,
		val,
	}

	return cmd
}

// TODO refactoring
func (sc *SetCmd) Process(st *Store) {
	var res *BoolCmdRes

	if key, keyErr := st.FindKey(sc.keyName); keyErr.Error() == ErrUndefinedKey {
		newKey := NewStringKey(sc.keyName)

		if err := st.AddKey(newKey); err != nil {
			res = NewBoolCmdRes(false, err)
			goto END
		}

		if err := newKey.SetValue(sc.val); err != nil {
			res = NewBoolCmdRes(false, err)
			goto END
		}

		res = NewBoolCmdRes(true, nil)
	} else if keyErr != nil {
		res = NewBoolCmdRes(false, keyErr)
	} else {
		switch strKey := key.(type) {
		case *StringKey:
			strKey.SetValue(sc.val)
			res = NewBoolCmdRes(true, nil)
		default:
			res = NewBoolCmdRes(false, errors.New(ErrInvalidKeyType))
		}
	}

END:
	sc.SetRes(res)
}
