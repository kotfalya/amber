package example

import (
	"github.com/kotfalya/store/store"
)

type MyStore struct {
	*store.Store
}

func NewMyStore() *MyStore {
	return &MyStore{
		store.NewStore(),
	}
}

func (ms *MyStore) MyGet(keyName string) (string, error) {
	cmd := NewBKGetCmd(keyName)
	cmd.Process(ms.Store)

	res := cmd.Res().(*store.StrCmdRes)

	return res.Val(), res.Err()
}

func (ms *MyStore) Get(keyName string) (string, error) {
	cmd := store.NewGetCmd(keyName)
	cmd.Process(ms.Store)

	res := cmd.Res().(*store.StrCmdRes)

	return res.Val(), res.Err()
}

func (ms *MyStore) Set(keyName string, value string) (bool, error) {
	cmd := store.NewSetCmd(keyName, value)
	cmd.Process(ms.Store)

	res := cmd.Res().(*store.BoolCmdRes)

	return res.Val(), res.Err()
}
