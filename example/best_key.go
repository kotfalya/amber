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

func (bkc *BKGetCmd) Process(st *store.Store) {
	var res *store.StrCmdRes

	if key, err := st.FindKey(bkc.keyName); err != nil {
		res = store.NewStrCmdRes("", err)
	} else {
		switch bestKey := key.(type) {
		case BestKey:
			res = store.NewStrCmdRes(bestKey.StrValue(), nil)
		default:
			res = store.NewStrCmdRes("", errors.New(store.ErrInvalidKeyType))
		}
	}

	bkc.SetRes(res)
}
