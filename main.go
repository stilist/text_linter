package main

import (
	"fmt"
	"github.com/stilist/text_linter/internal/dictionary"
	"github.com/stilist/text_linter/internal/linter"
	"github.com/stilist/text_linter/internal/rules"
	"log"
)

func main() {
	err := dictionary.LoadDefault()
	if err != nil {
		log.Fatal(err)
	}

	inputs := []string{}
	for _, e := range dictionary.Default {
		if len(e.Example) > 0 {
			inputs = append(inputs, e.Example)
		}
	}
	fmt.Printf("%d examples\n", len(inputs))
	for _, s := range inputs {
		l := linter.NewLinter(s, rules.Default)
		ps := l.Lint()
		fmt.Printf("%d error(s)\n", len(ps))
		for _, p := range ps {
			p.Describe()
		}
	}
}
