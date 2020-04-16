package dictionary

import (
	"bufio"
	"github.com/markbates/pkger"
)

var Default = Dictionary{}

func LoadDefault() error {
	f, err := pkger.Open("/data/dictionary.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	Default, err = FromCSV(r)

	return err
}
