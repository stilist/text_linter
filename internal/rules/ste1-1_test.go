package rules

import (
	"github.com/stilist/text_linter/internal/dictionary"
	"github.com/stilist/text_linter/internal/linter"
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

	type result struct {
		in string
		v  bool
	}
	expected := []result{
		{"test", true},
		{"TeSt", true},
		{"abaft", false},
		{"AbAfT", false},
	}

	r := RuleSTE11
	rs := []linter.Rule{r}

	for _, e := range expected {
		l := linter.NewLinter(e.in, rs)
		ps := l.Lint()
		if len(ps) > 0 && e.v {
			t.Errorf("Rule %s incorrectly rejected '%s'", r.Metadata.ID, e.in)
		} else if len(ps) == 0 && !e.v {
			t.Errorf("Rule %s incorrectly accepted '%s'", r.Metadata.ID, e.in)
		}
	}
}
