package rules

import (
	"github.com/stilist/text_linter/internal/linter"
	"strings"
)

var mdRuleSTE811 = linter.RuleMetadata{
	Name:        "unapproved punctuation",
	Description: "Do not use semicolons.",
	ID:          "STE-8.11",
	Severity:    linter.SevError,
}
var RuleSTE811 = linter.Rule{
	Metadata: mdRuleSTE811,
	Match: func(l *linter.Linter) (ps []linter.Problem) {
		cursor := 0
		for {
			offset := strings.Index(l.Text[cursor:], ";")
			if offset < 0 {
				break
			}

			start := cursor + offset
			cursor = start + 1

			p := linter.Problem{
				Text:         ";",
				Position:     linter.Position{offset, offset + 1},
				RuleMetadata: mdRuleSTE811,
			}
			ps = append(ps, p)
		}

		return ps
	},
}
