package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ablqk/littre-bot/parsers"
	"github.com/ablqk/littre-bot/src/dictionary"
	colour "github.com/fatih/color"
)

const (
	gobFile = "bin/dict.gob"
)

func main() {
	parsedEntries, err := parsers.ParseAlphabet("parsers/xmlittre-data")
	if err != nil {
		fmt.Println(err)
		return
	}

	// save it, for future reference
	if err := parsers.SaveGob(parsedEntries, gobFile); err != nil {
		fmt.Printf("unable to save gob: %s\n", err.Error())
	}
}

func out(w dictionary.Entry, at io.Writer) {
	pTerm(at, w.Term)
	pDef(at, w.Body.Def)
	pQuotes(at, strings.Join(w.Body.Quotes, "\n"))
	pTagLine(at, "Provided by Littré")
}

var (
	pTerm    = colour.New(colour.FgBlue).Add(colour.Bold).FprintlnFunc()
	pDef     = colour.New(colour.FgGreen).FprintlnFunc()
	pQuotes  = colour.New(colour.FgGreen).Add(colour.Italic).FprintlnFunc()
	pTagLine = colour.New(colour.FgYellow).FprintlnFunc()
)
