package rules

import (
	"testing"
)

func TestRuleSTE21(t *testing.T) {
	tcs := testcases{
		{"test", true},
		{"test test test", true},
		{"test test test test", false},
	}
	runTestcases(t, RuleSTE21, tcs)
}
