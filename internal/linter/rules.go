package linter

type Severity int

const (
	SevHint Severity = iota + 1
	SevWarn
	SevError
)

type testFunc func(l *Linter) (bool, []Position)
type Rule struct {
	Description string
	ID          string
	Severity    Severity
	Test        testFunc
}
type RuleSet []Rule
