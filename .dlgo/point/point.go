package point

type Point struct {
	Row int
	Col int
}

func (p *Point) Neighbours() []Point {
	return []Point{
		Point{p.Row - 1, p.Col},
		Point{p.Row + 1, p.Col},
		Point{p.Row, p.Col - 1},
		Point{p.Row, p.Col + 1},
	}
}
