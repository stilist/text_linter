package rules

import (
	"testing"
)

func TestRuleSTE811(t *testing.T) {
	tcs := testcases{
		{"test", 0, false},
		{"Test.", 0, false},
		{"Test. Test. Test.", 0, false},
		{"Test; Test. Test.", 1, false},
		{"Test; Test; Test.", 2, false},
		{"Test; Test; Test;", 3, false},
		{";Test; Test; Test", 3, false},
	}
	runTestcases(t, RuleSTE811, tcs)
}
