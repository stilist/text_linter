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

type testFunc func(l *Linter) (bool, []Position)
type Rule struct {
	Description string
	ID          string
	Severity
	Test testFunc
}
type RuleSet []Rule
