package constants

type status struct {
	ACTIVE    int
	DELETED   int
	PENDING   int
	COMPLETED int
}

var Status = status{
	ACTIVE:    1,
	DELETED:   2,
	PENDING:   3,
	COMPLETED: 4,
}
