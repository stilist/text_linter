package linter

import (
	"github.com/stilist/text_linter/internal/dictionary"
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
	Sentences  []Sentence
	Tokens     []Token
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
	if err := dictionary.LoadDefault(); err != nil {
		log.Fatal(err)
	}

	doc, err := prose.NewDocument(l.Text)
	if err != nil {
		log.Fatal(err)
	}
	l.document = doc
	l.paragraphs = strings.Split(l.Text, "\n\n")

	l.Sentences = []Sentence{}
	cursor := 0
	for _, ps := range doc.Sentences() {
		offset := strings.Index(l.Text[cursor:], ps.Text)
		start := cursor + offset
		cursor = start + len(ps.Text)

		s := Sentence{
			Text:     ps.Text,
			Position: Position{start, cursor},
		}
		l.Sentences = append(l.Sentences, s)
	}

	l.Tokens = []Token{}
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
		l.Tokens = append(l.Tokens, t)
	}
}

func (l *Linter) Lint() []Problem {
	failures := []Problem{}
	for _, r := range l.Rules {
		problems := r.Match(l)
		failures = append(failures, problems...)
	}
	return failures
}
