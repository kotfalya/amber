package db

type GetCmd struct {
	*BaseCmd
	keyName string
}

func NewGetCmd(keyName string) *GetCmd {
	cmd := &GetCmd{
		&BaseCmd{},
		keyName,
	}

	return cmd
}

func (gc *GetCmd) Process(db *DB) {
	db.req <- NewAddReq(gc.keyName)
}
