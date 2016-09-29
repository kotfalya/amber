package db

func (db *DB) Get(keyName string) (string, error) {
	cmd := NewGetCmd(keyName)
	cmd.Process(db)

	res := cmd.Res().(*StrCmdRes)

	return res.Val(), res.Err()
}

func (db *DB) Set(keyName string, value string) (string, error) {
	cmd := NewGetCmd(keyName)
	cmd.Process(db)

	res := cmd.Res().(*StrCmdRes)

	return res.Val(), res.Err()
}
