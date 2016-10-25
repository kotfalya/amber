package amber

type Cmd interface {
	Process(db *DB)
	Res() Res
	SetRes(res Res)
}

type BaseCmd struct {
	res Res
}

func (bc *BaseCmd) Res() Res {
	return bc.res
}

func (bc *BaseCmd) SetRes(res Res) {
	bc.res = res
}
