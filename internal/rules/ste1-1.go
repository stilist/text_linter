package rules

import (
	"github.com/stilist/text_linter/internal/dictionary"
	"github.com/stilist/text_linter/internal/linter"
)

var mdRuleSTE11 = linter.RuleMetadata{
	Description: "unapproved word",
	ID:          "STE-1.1",
	Severity:    linter.SevWarn,
}
var RuleSTE11 = linter.Rule{
	Metadata: mdRuleSTE11,
	Match: func(l *linter.Linter) (ps []linter.Problem) {
		for _, tok := range l.Tokens {
			entry := dictionary.Default.Find(tok.Text)
			if !entry.Approved && entry.Text != "" {
				p := linter.Problem{
					Text:         l.Text,
					Position:     tok.Position,
					RuleMetadata: mdRuleSTE11,
				}
				ps = append(ps, p)
			}
		}

		return ps
	},
}
