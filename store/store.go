package store

import "github.com/kotfalya/store/key"

type Store struct {
	page *Page
}

func NewStore() *Store {
	return &Store{
		page: NewPage(),
	}
}

func (s *Store) Load(keyName string) (*key.Key, error) {
	return s.page.Load(keyName)
}

func (s *Store) Save(keyName string, value interface{}) error {
	return s.page.Save(keyName, value)
}
