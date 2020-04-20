package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/stilist/text_linter/internal/linter"
	"github.com/stilist/text_linter/internal/rules"
	"io"
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

func readFile(filename string) (string, error) {
	var str strings.Builder

	// #nosec G304 -- specifically *want* to read arbitrary file paths
	f, err := os.Open(filename)
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

func main() {
	flag.Parse()

	var text string
	var err error
	if *filePtr == "-" {
		text, err = readStdin()
	} else {
		text, err = readFile(*filePtr)
	}
	if err != nil {
		log.Fatal(err)
	}

	l := linter.NewLinter(text, rules.Default)
	ps := l.Lint()
	if len(ps) > 0 {
		fmt.Printf("%d error(s)\n", len(ps))
		for _, p := range ps {
			p.Describe()
		}
	} else {
		lg := log.New(os.Stderr, "", 0)
		lg.Println("No errors!")
	}
}
