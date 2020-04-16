package rules

import (
	"github.com/stilist/text_linter/internal/linter"
)

var mdRuleSTE67 = linter.RuleMetadata{
	Description: "too many sentences in paragraph",
	ID:          "STE-6.7",
	Severity:    linter.SevError,
}
var RuleSTE67 = linter.Rule{
	Metadata: mdRuleSTE67,
	Match: func(l *linter.Linter) (ps []linter.Problem) {
		if len(l.Sentences) > 6 {
			p := linter.Problem{
				Text:         l.Text,
				Position:     l.Sentences[6].Position,
				RuleMetadata: mdRuleSTE67,
			}
			ps = append(ps, p)
		}

		return ps
	},
}
