package dictionary

import (
	"math/rand"
	"time"
)

func New(entries []Entry) *Dictionary {
	return &Dictionary{
		entries: entries,
	}
}

type Dictionary struct {
	entries []Entry
}

func (d *Dictionary) NewRandomWord() Entry {
	length := len(d.entries)
	rand.Seed(time.Now().Unix())
	r := rand.Intn(length)
	return d.entries[r]
}

type Entry struct {
	Term      string
	Acception int
	Header    string
	Body      EntryBody
	Link      string
}

type EntryBody struct {
	Def    string
	Quotes []string
}
