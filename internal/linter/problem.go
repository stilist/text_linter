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

const samMaxLen = 80
const (
	snipBef = "[…] "
	snipAft = " […]"
)

func (p *Problem) Describe() {
	sam := p.Text[p.Position.Start:p.Position.End]
	samLen := len(sam)
	if samLen < samMaxLen {
		if len(p.Text) > p.Position.End && samLen+len(snipAft) <= samMaxLen {
			sam = sam + snipAft
		}
		if p.Position.Start > 0 && samLen+len(snipBef) <= samMaxLen {
			sam = snipBef + sam
		}
	}

	fmt.Printf("%s: %s @[%d,%d):\n-> %s\n\n",
		p.RuleMetadata.ID,
		p.RuleMetadata.Name,
		p.Position.Start,
		p.Position.End,
		sam)
}
