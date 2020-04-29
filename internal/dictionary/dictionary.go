package dictionary

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"sync"
)

type Entry struct {
	Approved         bool
	Text             string
	Tag              string
	Meaning          string
	Note             string
	Alternatives     []string
	Examples         []string
	NegativeExamples []string
}

func (e *Entry) CheckTag(t string) bool {
	return t == e.Tag || t != "!"+e.Tag
}

type Dictionary map[string]Entry

// @see https://www.ling.upenn.edu/courses/Fall_2003/ling001/penn_treebank_pos.html
var posLookup = map[string]string{
	"adj": "JJ",
	"adv": "RB",
	"art": "DT",
	// CC for coordinating conjunction
	// IN subordinating conjunction
	"con": "CC",
	"n":   "NN",
	// PRP for personal pronoun
	// PRP$ for possessive pronoun
	// WP for wh-pronoun
	// WP$ for possessive wh-pronoun
	"pn":  "PRP",
	"pre": "IN",
	"v":   "VB",
}

var (
	once             sync.Once
	pos_re           *regexp.Regexp
	note_re          = regexp.MustCompile(`\s+NOTE:\s+(.+)`)
	approved_ks_re   = regexp.MustCompile(`^[A-Z,\s]+`)
	unapproved_ks_re = regexp.MustCompile(`^[a-z,\s]+`)
)

func prepare_pos_re() {
	var pos_tags []string
	for pos := range posLookup {
		pos_tags = append(pos_tags, pos)
	}
	pos_re = regexp.MustCompile(`\s+\((` + strings.Join(pos_tags, "|") + `)\)`)
}

func entriesFromRow(row []string) ([]Entry, error) {
	es := []Entry{}

	keyword := row[0]

	var (
		pos          string
		note         string
		meaning      string
		alternatives []string
		ks_re        *regexp.Regexp
	)

	once.Do(prepare_pos_re)
	if matches := pos_re.FindStringSubmatch(keyword); matches != nil {
		pos = matches[1]
		keyword = strings.Replace(keyword, matches[0], "", 1)
	} else {
		fmt.Printf("Didn't find a PoS tag in '%s'.", keyword)
	}

	if matches := note_re.FindStringSubmatch(keyword); matches != nil {
		note = matches[1]
		keyword = strings.Replace(keyword, matches[0], "", 1)
	}

	l1 := string(keyword[0])
	approved := l1 == strings.ToUpper(l1)

	// This variable has a generic name because the column's data serves two
	// incompatible purposes: for approved words it lists the approved
	// definition; for unapproved words it lists approved alternative words.
	col2 := row[1]
	// Either of the first two columns may have a NOTE, but there won't be a
	// NOTE in *both* columns.
	if note == "" {
		if matches := note_re.FindStringSubmatch(col2); matches != nil {
			note = matches[1]
			col2 = strings.Replace(col2, matches[0], "", 1)
		}
	}
	if approved {
		meaning = col2
	} else {
		alternatives = strings.Split(strings.ToLower(col2), ",")
	}

	examples := strings.Split(row[2], "\n\n")
	negative_examples := strings.Split(row[3], "\n\n")

	// This regex strips out cross-referenced words.
	//
	// @example
	//   "BE, IS, WAS (also ARE, WERE)" will match "BE, IS, WAS"
	// @example
	//   "BAD (adj) (WORSE, WORST)" will match "BAD" (WORSE/WORST are listed
	//   separately)
	if approved {
		ks_re = approved_ks_re
	} else {
		ks_re = unapproved_ks_re
	}
	all_ks := ks_re.FindString(keyword)
	ks := strings.Split(all_ks, ", ")
	tag := posLookup[pos]
	for i, k := range ks {
		// Some verbs will have a list of three or four approved tenses. This
		// tags the additional tenses.
		//
		// @example
		//   BECOME, BECOMES, BECAME
		//   BLOW, BLOWS, BLEW, BLOWN
		if pos == "v" && i > 0 {
			if i == 1 {
				// BECOMES, BLOWS: Verb, gerund or present participle
				tag = "VBG"
			} else if i == len(ks)-1 {
				// BECAME, BLOWN: Verb, past tense
				tag = "VBD"
			} else {
				// BLEW: Verb, past participle
				tag = "VBN"
			}
		}

		text := strings.Trim(strings.ToLower(k), " ")
		e := Entry{
			Approved:         approved,
			Text:             text,
			Note:             note,
			Tag:              tag,
			Alternatives:     alternatives,
			Meaning:          meaning,
			Examples:         examples,
			NegativeExamples: negative_examples,
		}
		es = append(es, e)
	}

	return es, nil
}

var CSVColumns = []string{
	"Keyword",
	"Approved meaning/ALTERNATIVES",
	"APPROVED EXAMPLE",
	"Not approved",
}

func FromCSV(r io.Reader) (Dictionary, error) {
	d := Dictionary{}
	// skip header row
	br := bufio.NewReader(r)
	_, err := br.ReadString('\n')
	if err != nil {
		return d, err
	}

	cr := csv.NewReader(br)
	cr.FieldsPerRecord = len(CSVColumns)
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		es, err := entriesFromRow(row)
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range es {
			d[e.Text] = e
		}
	}

	return d, nil
}

func (d Dictionary) Find(w string) (e Entry) {
	lc := strings.ToLower(w)
	return d[lc]
}
