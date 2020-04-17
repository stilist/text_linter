package rules

import (
	"testing"
)

func TestRuleSTE21(t *testing.T) {
	tcs := testcases{
		{"test", 0, false},
		{"test test test", 0, false},
		{"test test test test", 1, false},
		{"test test test test test", 2, false},
	}
	runTestcases(t, RuleSTE21, tcs)
}
