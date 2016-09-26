package example

import (
	"errors"
	"github.com/kotfalya/store/store"
)

var (
	_ store.Key = (*BestKey)(nil)
	_ store.Cmd = (*BKGetCmd)(nil)
)

type BestKey struct {
	*store.StringKey
}

func NewBestKey(name string) *BestKey {
	return &BestKey{
		store.NewStringKey(name),
	}
}

type BKGetCmd struct {
	*store.BaseCmd
	keyName string
}

func NewBKGetCmd(keyName string) *BKGetCmd {
	cmd := &BKGetCmd{
		&store.BaseCmd{},
		keyName,
	}

	return cmd
}

func (bkc *BKGetCmd) Process(st *store.Store) error {
	var res *store.StrCmdRes

	if key, err := st.GetKey(bkc.keyName); err != nil {
		res = store.NewStrCmdRes("", err)
	} else {
		switch bestKey := key.(type) {
		case BestKey:
			res = store.NewStrCmdRes(bestKey.StrValue(), nil)
		default:
			return errors.New(store.ErrInvalidKeyType)
		}
	}

	bkc.SetRes(res)

	return nil
}

func (bkc *BKGetCmd) New(name string) store.Key {
	return NewBestKey(name)
}
