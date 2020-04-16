package linter

import (
	"fmt"
)

type Problem struct {
	Text        string
	Position    *Position
	Rule        *Rule
	Replacement string
}

func (p *Problem) Describe() {
	fmt.Printf("%s %s (severity %d) @[%d,%d):\n-> %s\n",
		p.Rule.ID,
		p.Rule.Description,
		p.Rule.Severity,
		p.Position.Start,
		p.Position.End,
		p.Text[p.Position.Start:p.Position.End])
}
