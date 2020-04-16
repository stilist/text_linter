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
		{"test", 0},
		{"TeSt", 0},
		{"abaft", 1},
		{"AbAfT", 1},
	}
	runTestcases(t, RuleSTE11, tcs)
}
