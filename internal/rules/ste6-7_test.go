package rules

import (
	"testing"
)

func TestRuleSTE67(t *testing.T) {
	tcs := testcases{
		{"test", 0},
		{"Test.", 0},
		{"Test. Test. Test.", 0},
		{"Test. Test. Test. Test. Test. Test.", 0},
		{"Test. Test. Test. Test. Test. Test. Test.", 1},
		{"Test! Test. Test. Test. Test. Test. Test.", 1},
	}
	runTestcases(t, RuleSTE67, tcs)
}
