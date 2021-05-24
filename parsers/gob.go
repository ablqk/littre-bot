package parsers

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/ablqk/littre-bot/src/dictionary"
)

func ParseGob(gobFile string) ([]dictionary.Entry, error) {
	f, err := os.Open(gobFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open gob: %w", err)
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	var dict []dictionary.Entry
	err = decoder.Decode(&dict)
	if err != nil {
		return nil, fmt.Errorf("unable to parse gob: %w", err)
	}

	return dict, nil
}

func SaveGob(dict []dictionary.Entry, gobFile string) error {
	f, err := os.Create(gobFile)
	if err != nil {
		return fmt.Errorf("unable to create gob file: %w", err)
	}
	defer f.Close()
	encoder := gob.NewEncoder(f)
	err = encoder.Encode(dict)
	if err != nil {
		return fmt.Errorf("unable to encode gob: %w", err)
	}
	return nil
}
