package rules

import (
	"testing"
)

func TestRuleSTE67(t *testing.T) {
	tcs := testcases{
		{"test", true},
		{"Test.", true},
		{"Test. Test. Test.", true},
		{"Test. Test. Test. Test. Test. Test.", true},
		{"Test. Test. Test. Test. Test. Test. Test.", false},
		{"Test! Test. Test. Test. Test. Test. Test.", false},
	}
	runTestcases(t, RuleSTE67, tcs)
}
