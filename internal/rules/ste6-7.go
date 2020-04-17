package rules

import (
	"github.com/stilist/text_linter/internal/linter"
)

var mdRuleSTE67 = linter.RuleMetadata{
	Name:        "paragraph lengths",
	Description: "The maximum length of a paragraph is 6 sentences. Do not use one-sentence paragraphs more than once in every 10 paragraphs.",
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
