package store

import (
	"github.com/kotfalya/store/types"
)

func (s *Store) Exists(keyName string) (bool, error) {
	if _, err := s.page.Load(keyName, getBalancingIndex(keyName)); err.Error() == ErrUndefinedKey {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (s *Store) Get(keyName string) (string, error) {
	if key, err := s.page.Load(keyName, getBalancingIndex(keyName)); err != nil {
		return "", err
	} else {
		return key.Value().(string), nil
	}
}

func (s *Store) Set(keyName string, value interface{}) error {
	var (
		index int = getBalancingIndex(keyName)
		key   types.Key
		err   error
	)

	key, err = s.page.Load(keyName, index)
	if err.Error() == ErrUndefinedKey {
		key = types.NewStringKey(keyName)
		if err = s.page.Add(key, index); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return key.SetValue(value)
}

func getBalancingIndex(keyName string) int {
	return 5
}
