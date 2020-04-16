package linter

import (
	"fmt"
)

type Problem struct {
	Text string
	Position
	Rule
	Replacement string
}

func (p *Problem) Describe() {
	fmt.Printf("%s: %s @[%d,%d):\n-> %s\n",
		p.Rule.ID,
		p.Rule.Description,
		p.Position.Start,
		p.Position.End,
		p.Text[p.Position.Start:p.Position.End])
}
