package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ablqk/littre-bot/parsers"
	"github.com/ablqk/littre-bot/src/dictionary"
	"github.com/fatih/color"
)

func main() {
	var parsedEntries []dictionary.Entry
	var err error

	if _, err = os.Stat("bin/dict.gob"); os.IsNotExist(err) {
		parsedEntries, err = parsers.ParseAlphabet("parsers/xmlittre-data")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// assume file exists
		parsedEntries, err = parsers.ParseGob()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	d := dictionary.New(parsedEntries)
	w := d.NewRandomWord()

	out(w, os.Stdout)
}

func out(w dictionary.Entry, at io.Writer) {
	pTerm(at, w.Term)
	pDef(at, w.Body.Def)
	pDef(at, w.Body.Quotes)
	pTagLine(at, "Provided by Littré")
}

var (
	pNL = func(at io.Writer) {
		_, _ = fmt.Fprintln(at)
	}
	pTerm    = color.New(color.FgBlue).Add(color.Bold).FprintlnFunc()
	pDef     = color.New(color.FgGreen).FprintlnFunc()
	pTagLine = color.New(color.FgYellow).FprintlnFunc()
)


