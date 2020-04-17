package rules

import (
	"github.com/stilist/text_linter/internal/linter"
	"testing"
)

type testcase struct {
	in string
	c  int
	fn bool
}
type testcases []testcase

func runTestcases(t *testing.T, r linter.Rule, tcs testcases) {
	rs := linter.RuleSet{r}

	for _, tc := range tcs {
		l := linter.NewLinter(tc.in, rs)
		ps := l.Lint()
		if len(ps) != tc.c {
			if tc.fn {
				t.Logf("Warning: Rule %s found %d problem(s); should have been %d. Likely a false negative. Input: '%s'.",
					r.Metadata.ID,
					len(ps),
					tc.c,
					tc.in)
			} else {
				t.Errorf("Rule %s found %d problem(s); should have been %d. Input: '%s'.",
					r.Metadata.ID,
					len(ps),
					tc.c,
					tc.in)
			}
		}
	}
}
