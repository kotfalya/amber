package example

import (
	"github.com/kotfalya/db/store"
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
