package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ablqk/littre-bot/parsers"
	"github.com/ablqk/littre-bot/src/dictionary"
	colour "github.com/fatih/color"
)

const (
	gobFile = "bin/dict.gob"
)

func main() {
	var parsedEntries []dictionary.Entry
	var err error

	if _, err = os.Stat(gobFile); os.IsNotExist(err) {
		parsedEntries, err = parsers.ParseAlphabet("parsers/xmlittre-data")
		if err != nil {
			fmt.Println(err)
			return
		}
		// save it, for future reference
		if err := parsers.SaveGob(parsedEntries, gobFile); err != nil {
			// this is not blocking
			fmt.Fprintf(os.Stderr, "unable to save gob: %s", err.Error())
		}
	} else {
		// assume file exists
		parsedEntries, err = parsers.ParseGob(gobFile)
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
	pTerm    = colour.New(colour.FgBlue).Add(colour.Bold).FprintlnFunc()
	pDef     = colour.New(colour.FgGreen).FprintlnFunc()
	pTagLine = colour.New(colour.FgYellow).FprintlnFunc()
)


