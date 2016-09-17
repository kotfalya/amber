package store

import "github.com/kotfalya/store/utils"

type Store struct {
	page *Page
}

func NewStore() *Store {
	return &Store{
		page: NewPage(),
	}
}

func getBalancingIndex(keyName string) int {
	return utils.TextToIndex(keyName, pageChildSize)
}
