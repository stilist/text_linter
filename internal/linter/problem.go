package linter

import (
	"fmt"
)

type Problem struct {
	Text string
	Position
	RuleMetadata
	Replacement string
}

func (p *Problem) Describe() {
	fmt.Printf("%s: %s @[%d,%d):\n-> %s\n",
		p.RuleMetadata.ID,
		p.RuleMetadata.Description,
		p.Position.Start,
		p.Position.End,
		p.Text[p.Position.Start:p.Position.End])
}
