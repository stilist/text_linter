package linter

type Position struct {
	Start int
	End   int
}

func (p *Position) Len() int {
	return p.End - p.Start
}
