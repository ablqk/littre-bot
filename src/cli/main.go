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
	fromXML, err := parsers.ParseAlphabet("parsers/xmlittre-data")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := dictionary.New(fromXML)
	w := d.NewRandomWord()

	out(w, os.Stdout)
}

func out(w dictionary.Entry, at io.Writer) {
	pTerm(at, w.Term)
	pNL(at)
	pDef(at, w.Body.Def)
	pNL(at)
	pDef(at, w.Body.Quotes)
	pNL(at)
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


