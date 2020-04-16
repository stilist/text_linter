package rules

import (
	"github.com/stilist/text_linter/internal/linter"
	"testing"
)

type testcase struct {
	in string
	v  bool
}
type testcases []testcase

func runTestcases(t *testing.T, r linter.Rule, tcs testcases) {
	rs := linter.RuleSet{r}

	for _, tc := range tcs {
		l := linter.NewLinter(tc.in, rs)
		ps := l.Lint()
		if len(ps) > 0 && tc.v {
			t.Errorf("Rule %s incorrectly rejected '%s'", r.Metadata.ID, tc.in)
		} else if len(ps) == 0 && !tc.v {
			t.Errorf("Rule %s incorrectly accepted '%s'", r.Metadata.ID, tc.in)
		}
	}
}
