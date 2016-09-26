package store

type Store struct {
	page *Page
}

func NewStore() *Store {
	return &Store{
		page: NewPage(),
	}
}

func (s *Store) FindKey(keyName string) (Key, error) {
	req := NewPageReq("get", keyName)
	s.page.AddReq(req)
	res := req.Done().(*KeyPageRes)

	return res.Val(), res.Err()
}

func (s *Store) AddKey(key Key) error {
	req := NewPageReq("add", key)
	s.page.AddReq(req)
	res := req.Done()

	return res.Err()
}

func (s *Store) Get(keyName string) (string, error) {
	cmd := NewGetCmd(keyName)
	cmd.Process(s)

	res := cmd.Res().(*StrCmdRes)

	return res.Val(), res.Err()
}

func (s *Store) Set(keyName string, value string) (bool, error) {
	cmd := NewSetCmd(keyName, value)
	cmd.Process(s)

	res := cmd.Res().(*BoolCmdRes)

	return res.Val(), res.Err()
}
