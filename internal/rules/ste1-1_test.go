package rules

import (
	"testing"
)

// @note This test is fragile and depends on the contents of
//   `dictionary.Default`.
func TestRuleSTE11(t *testing.T) {
	tcs := testcases{
		{"test", 0, false},
		{"TeSt", 0, false},
		{"abaft", 1, false},
		{"AbAfT", 1, false},
	}
	runTestcases(t, RuleSTE11, tcs)
}
