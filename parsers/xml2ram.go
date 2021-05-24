package parsers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ablqk/littre-bot/src/dictionary"
)

// ParseAlphabet reads a file for each letter, called a.xml, b.xml, etc, in the given folder.
func ParseAlphabet(folderPath string) ([]dictionary.Entry, error) {
	var dico []dictionary.Entry
	// we want to protect the dictionary from concurrent writes
	var mut sync.Mutex
	// all letters will be parsed in parallel
	wg := sync.WaitGroup{}
	var wgErr error
	for r := 'a'; r <= 'z'; r++ {
		wg.Add(1)
		go func(r rune) {
			defer wg.Done()
			e, err := parseLetter(fmt.Sprintf("%s/%s.xml", folderPath, string(r)))
			if err != nil {
				wgErr = err
			}
			mut.Lock()
			dico = append(dico, e...)
			mut.Unlock()
		}(r)
	}
	wg.Wait()
	if wgErr != nil {
		return nil, wgErr
	}

	return dico, nil
}

// parseLetter parses one file and returns its entries
func parseLetter(file string) ([]dictionary.Entry, error) {
	// Open our xmlFile
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	letter, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var entries xmlittre
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(letter, &entries)
	if err != nil {
		return nil, err
	}
	return format(entries), nil
}

func format(x xmlittre) []dictionary.Entry {
	ret := make([]dictionary.Entry, 0)
	for _, e := range x.Entree {
		for _, v := range e.Corps.Variante {
			ee := dictionary.Entry{
				Term:   e.Terme,
				Header: e.Entete.Text,
				Body: dictionary.EntryBody{
					Def: func() string {
						if len(v.Num) > 0 {
							return v.Num + ". " + v.Text
						} else {
							return v.Text
						}
					}(),
				},
			}
			ee.Body.Quotes = func() []string {
				quotes := make([]string, 0)
				for _, q := range v.Cit {
					quotes = append(quotes, q.Text)
				}
				return quotes
			}()
			ret = append(ret, ee)
		}
	}
	return ret
}

// xmlittre is a marshalling-only struct
// generated via https://www.onlinetool.io/xmltogo/
type xmlittre struct {
	XMLName xml.Name `xml:"xmlittre"`
	Lettre  string   `xml:"lettre,attr"`
	Entree  []struct {
		Text       string `xml:",chardata"`
		Terme      string `xml:"terme,attr"`
		Supplement string `xml:"supplement,attr"`
		Sens       int    `xml:"sens,attr"`
		Entete     struct {
			Text          string `xml:",chardata"`
			Prononciation string `xml:"prononciation"`
			Nature        string `xml:"nature"`
		} `xml:"entete"`
		Corps struct {
			Text     string `xml:",chardata"`
			Variante []struct {
				Text   string `xml:",chardata"`
				Num    string `xml:"num,attr"`
				Indent []struct {
					Text       string `xml:",chardata"`
					Semantique struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"semantique"`
					I   []string `xml:"i"`
					Cit []struct {
						Text string `xml:",chardata"`
						Aut  string `xml:"aut,attr"`
						Ref  string `xml:"ref,attr"`
						I    string `xml:"i"`
					} `xml:"cit"`
					A struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"a"`
					Nature string `xml:"nature"`
				} `xml:"indent"`
				I          string `xml:"i"`
				Semantique struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"semantique"`
				A struct {
					Text string `xml:",chardata"`
					Ref  string `xml:"ref,attr"`
				} `xml:"a"`
				Cit []struct {
					Text string `xml:",chardata"`
					Aut  string `xml:"aut,attr"`
					Ref  string `xml:"ref,attr"`
					A    struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"a"`
				} `xml:"cit"`
			} `xml:"variante"`
		} `xml:"corps"`
		Rubrique []struct {
			Text   string `xml:",chardata"`
			Nom    string `xml:"nom,attr"`
			Indent []struct {
				Text string `xml:",chardata"`
				I    []struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"i"`
				Cit struct {
					Text string `xml:",chardata"`
					Aut  string `xml:"aut,attr"`
					Ref  string `xml:"ref,attr"`
				} `xml:"cit"`
				A struct {
					Text string `xml:",chardata"`
					Ref  string `xml:"ref,attr"`
				} `xml:"a"`
			} `xml:"indent"`
		} `xml:"rubrique"`
	} `xml:"entree"`
}
