package store

type Store struct {
	page *Page
}

func NewStore() *Store {
	return &Store{
		page: NewPage(),
	}
}

func (s *Store) GetKey(keyName string) (Key, error) {
	req := NewPageReq("get", keyName)
	s.page.AddReq(req)
	res := req.Done().(*KeyPageRes)

	return res.Val(), res.Err()
}
