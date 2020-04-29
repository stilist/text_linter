package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/stilist/text_linter/internal/linter"
	"github.com/stilist/text_linter/internal/rules"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	defaultFile = "-"
	usage       = "name of file to lint (or - to read from STDIN)"
)

var filePtr = flag.String("file", defaultFile, usage)

func init() {
	flag.StringVar(filePtr, "f", defaultFile, usage+" (shorthand)")
}

func readStdin() (string, error) {
	var str strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return str.String(), err
	}
	return str.String(), nil
}

func readFile(path string) (string, error) {
	var str strings.Builder

	// #nosec G304 -- specifically *want* to read arbitrary file paths
	f, err := os.Open(path)
	if err != nil {
		return str.String(), err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			switch err {
			default:
				return str.String(), err
			case io.EOF:
				return str.String(), nil
			}
		}
		str.WriteString(string(line))
	}
}

func readFiles(path string) (texts []string, err error) {
	p, err := os.Stat(path)
	if err != nil {
		return texts, err
	}

	switch mode := p.Mode(); {
	case mode.IsDir():
		fis, err := ioutil.ReadDir(path)
		if err != nil {
			return texts, err
		}
		for _, fi := range fis {
			text, err := readFile(fi.Name())
			if err != nil {
				return texts, err
			}
			texts = append(texts, text)
		}
	case mode.IsRegular():
		text, err := readFile(path)
		if err != nil {
			return texts, err
		}
		texts = append(texts, text)
	}

	return texts, nil
}

func lint(text string) {
	l := linter.NewLinter(text, rules.Default)
	probs := l.Lint()
	if len(probs) > 0 {
		fmt.Printf("%d error(s)\n", len(probs))
		for _, prob := range probs {
			prob.Describe()
		}
	} else {
		lg := log.New(os.Stderr, "", 0)
		lg.Println("No errors!")
	}
}

func main() {
	flag.Parse()

	if *filePtr == "-" {
		text, err := readStdin()
		if err != nil {
			log.Fatal(err)
		}
		lint(text)
	} else {
		texts, err := readFiles(*filePtr)
		if err != nil {
			log.Fatal(err)
		}
		for _, text := range texts {
			lint(text)
		}
	}
}
