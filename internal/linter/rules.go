package linter

type Severity struct {
	Label string
	Score int
}

var (
	SevHint = Severity{
		Label: "hint",
		Score: 1,
	}
	SevWarn = Severity{
		Label: "warning",
		Score: 2,
	}
	SevError = Severity{
		Label: "error",
		Score: 3,
	}
)

type RuleMetadata struct {
	Description string
	ID          string
	Severity
}
type matchFunc func(l *Linter) []Problem
type Rule struct {
	Metadata RuleMetadata
	Match    matchFunc
}
type RuleSet []Rule
