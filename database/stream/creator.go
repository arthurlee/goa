package stream

const (
	CREATE_RECORD = iota
	CREATE_LIST
)

type CreatorType int

type Creator func(creatorType CreatorType) interface{}
