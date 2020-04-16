package rules

import (
	"testing"
)

func TestRuleSTE21(t *testing.T) {
	tcs := testcases{
		{"test", 0},
		{"test test test", 0},
		{"test test test test", 1},
		{"test test test test test", 2},
	}
	runTestcases(t, RuleSTE21, tcs)
}
