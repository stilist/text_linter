package rules

import (
	"github.com/stilist/text_linter/internal/dictionary"
	"log"
	"testing"
)

// @note This test is fragile and depends on the contents of
//   `dictionary.Default`.
func TestRuleSTE12(t *testing.T) {
	err := dictionary.LoadDefault()
	if err != nil {
		log.Fatal(err)
	}

	tcs := testcases{
		{"For data about the location of circuit breakers, refer to the wiring list.", 0, false},
		{"Drain approximately 2 liters of fuel from the tank.", 0, false},
		{"Turn the shaft around its axis.", 0, false},
		{"Drain about 2 liters of fuel from the tank.", 1, true},
		{"Rotate the shaft about its axis.", 1, true},
	}
	runTestcases(t, RuleSTE12, tcs)
}
