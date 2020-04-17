package dictionary

import (
	"bufio"
	"encoding/csv"
	"io"
	"strings"
)

type Entry struct {
	Approved bool
	Text     string
	// @see https://www.ling.upenn.edu/courses/Fall_2003/ling001/penn_treebank_pos.html
	Tag              string
	Meaning          string
	Alternatives     []string
	Examples         []string
	NegativeExamples []string
}

func (e *Entry) CheckTag(t string) bool {
	return t == e.Tag || t != "!"+e.Tag
}

type Dictionary map[string]Entry

var csvHeader = []string{
	"keyword",
	"tag",
	"meaning",
	"alternatives",
	"example",
	"negative_example",
}

func FromCSV(r io.Reader) (Dictionary, error) {
	d := Dictionary{}

	// skip header row
	br := bufio.NewReader(r)
	_, err := br.ReadString('\n')
	if err != nil {
		return nil, err
	}

	cr := csv.NewReader(br)
	cr.FieldsPerRecord = len(csvHeader)
	for {
		// @todo skip header
		row, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		text := row[0]
		e := Entry{
			Approved:         row[2] != "",
			Text:             text,
			Tag:              row[1],
			Alternatives:     strings.Split(row[3], ","),
			Meaning:          row[2],
			Examples:         strings.Split(row[4], ". "),
			NegativeExamples: strings.Split(row[5], ". "),
		}

		d[text] = e
	}

	return d, nil
}

func (d Dictionary) Find(w string) (e Entry) {
	lc := strings.ToLower(w)
	return d[lc]
}
