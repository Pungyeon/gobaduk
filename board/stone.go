package board

type Stone struct {
	liberties int
	player    int
	groupid   int
}

type Group struct {
	liberties int
	player    int
	id        string
}
