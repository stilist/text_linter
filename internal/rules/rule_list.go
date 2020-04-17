package rules

import (
	"github.com/stilist/text_linter/internal/linter"
)

var Default linter.RuleSet = []linter.Rule{
	RuleSTE11,
	RuleSTE12,
	RuleSTE21,
	RuleSTE67,
}
