package linter

import (
	"gopkg.in/jdkato/prose.v2"
	"log"
	"strings"
)

type Sentence struct {
	Text string
	Position
}

type Token struct {
	Text  string
	Tag   string
	Label string
	Position
}

type Linter struct {
	Text       string
	document   *prose.Document
	paragraphs []string
	Rules      RuleSet
	sentences  []Sentence
	tokens     []Token
}

func NewLinter(text string, rs RuleSet) *Linter {
	l := Linter{
		Text:  text,
		Rules: rs,
	}
	l.Initialize()
	return &l
}

func (l *Linter) Initialize() {
	doc, err := prose.NewDocument(l.Text)
	if err != nil {
		log.Fatal(err)
	}
	l.document = doc
	l.paragraphs = strings.Split(l.Text, "\n\n")

	l.sentences = []Sentence{}
	cursor := 0
	for _, ps := range doc.Sentences() {
		offset := strings.Index(l.Text[cursor:], ps.Text)
		start := cursor + offset
		cursor = start + len(ps.Text)

		s := Sentence{
			Text:     ps.Text,
			Position: Position{start, cursor},
		}
		l.sentences = append(l.sentences, s)
	}

	l.tokens = []Token{}
	cursor = 0
	for _, pt := range doc.Tokens() {
		offset := strings.Index(l.Text[cursor:], pt.Text)
		start := cursor + offset
		cursor = start + len(pt.Text)

		t := Token{
			Text:     pt.Text,
			Tag:      pt.Tag,
			Label:    pt.Label,
			Position: Position{start, cursor},
		}
		l.tokens = append(l.tokens, t)
	}
}

func (l *Linter) Lint() []Problem {
	failures := []Problem{}
	for _, r := range l.Rules {
		passed, pos := r.Test(l)

		if !passed {
			for _, p := range pos {
				problem := Problem{
					Text:     l.Text,
					Position: &p,
					Rule:     &r,
				}
				failures = append(failures, problem)
				problem.Describe()
			}
		}
	}
	return failures
}
