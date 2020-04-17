package rules

import (
	"testing"
)

func TestRuleSTE67(t *testing.T) {
	tcs := testcases{
		{"test", 0, false},
		{"Test.", 0, false},
		{"Test. Test. Test.", 0, false},
		{"Test. Test. Test. Test. Test. Test.", 0, false},
		{"Test. Test. Test. Test. Test. Test. Test.", 1, false},
		{"Test! Test. Test. Test. Test. Test. Test.", 1, false},
	}
	runTestcases(t, RuleSTE67, tcs)
}
