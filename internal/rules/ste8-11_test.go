package rules

import (
	"testing"
)

func TestRuleSTE811(t *testing.T) {
	tcs := testcases{
		{"test", 0},
		{"Test.", 0},
		{"Test. Test. Test.", 0},
		{"Test; Test. Test.", 1},
		{"Test; Test; Test.", 2},
		{"Test; Test; Test;", 3},
		{";Test; Test; Test", 3},
	}
	runTestcases(t, RuleSTE811, tcs)
}
