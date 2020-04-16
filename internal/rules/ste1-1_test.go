package rules

import (
	"github.com/stilist/text_linter/internal/dictionary"
	"log"
	"testing"
)

// @note This test is fragile and depends on the contents of
//   `dictionary.Default`.
func TestRuleSTE11(t *testing.T) {
	err := dictionary.LoadDefault()
	if err != nil {
		log.Fatal(err)
	}

	tcs := testcases{
		{"test", true},
		{"TeSt", true},
		{"abaft", false},
		{"AbAfT", false},
	}

	runTestcases(t, RuleSTE11, tcs)
}
