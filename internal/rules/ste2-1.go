package rules

import (
	"github.com/stilist/text_linter/internal/linter"
)

var mdRuleSTE21 = linter.RuleMetadata{
	Name: "noun cluster is too large",
	Description: "Do not make noun clusters of more than three nouns.",
	ID:          "STE-2.1",
	Severity:    linter.SevError,
}
var RuleSTE21 = linter.Rule{
	Metadata: mdRuleSTE21,
	Match: func(l *linter.Linter) (ps []linter.Problem) {
		clusLen := 0
		for _, tok := range l.Tokens {
			if tok.Tag == "NN" {
				clusLen++
				if clusLen > 3 {
					p := linter.Problem{
						Text:         l.Text,
						Position:     tok.Position,
						RuleMetadata: mdRuleSTE21,
					}
					ps = append(ps, p)
				}
			} else {
				clusLen = 0
			}
		}

		return ps
	},
}
