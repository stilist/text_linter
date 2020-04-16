package linter

import (
	"github.com/stilist/text_linter/internal/dictionary"
)

var DefaultRules = RuleSet{
	{
		Description: "unapproved word",
		ID:          "STE-1.1",
		Severity:    SevWarn,
		Test: func(l *Linter) (bool, []Position) {
			valid := true
			pos := []Position{}

			for _, tok := range l.tokens {
				entry := dictionary.Default.Find(tok.Text)
				if entry.Alternatives != nil {
					valid = false
					if !valid {
						pos = append(pos, tok.Position)
					}
				}
			}

			return valid, pos
		},
	},
	{
		Description: "noun cluster too large",
		ID:          "STE-2.1",
		Severity:    SevError,
		Test: func(l *Linter) (bool, []Position) {
			valid := true
			clusLen := 0
			pos := []Position{}

			for _, tok := range l.tokens {
				if tok.Tag == "NN" {
					clusLen += 1
					if clusLen > 3 {
						valid = false
						pos = append(pos, tok.Position)
					}
				} else {
					clusLen = 0
				}
			}

			return valid, pos
		},
	},
	{
		Description: "too many sentences in paragraph",
		ID:          "STE-6.7",
		Severity:    SevError,
		Test: func(l *Linter) (bool, []Position) {
			pos := []Position{}

			valid := len(l.sentences) <= 6
			if !valid {
				pos = append(pos, l.sentences[6].Position)
			}
			return valid, pos
		},
	},
}
