package rules

import (
	"github.com/stilist/text_linter/internal/dictionary"
	"github.com/stilist/text_linter/internal/linter"
)

var mdRuleSTE12 = linter.RuleMetadata{
	Name:        "word usage is not approved",
	Description: "Use approved words from the Dictionary only as the part of speech given.",
	ID:          "STE-1.2",
	Severity:    linter.SevError,
}
var RuleSTE12 = linter.Rule{
	Metadata: mdRuleSTE12,
	Match: func(l *linter.Linter) (ps []linter.Problem) {
		for _, tok := range l.Tokens {
			entry := dictionary.Default.Find(tok.Text)
			if entry.Text != "" && !entry.CheckTag(tok.Tag) {
				p := linter.Problem{
					Text:         l.Text,
					Position:     tok.Position,
					RuleMetadata: mdRuleSTE12,
				}
				ps = append(ps, p)
			}
		}

		return ps
	},
}
