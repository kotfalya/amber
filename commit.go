package amber

type Commit struct {
	id     uint32
	parent uint32
}

type Diff struct {
	commitId  uint32
	payload   interface{}
	timestamp int64
	applied   bool
}
