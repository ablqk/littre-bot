package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ablqk/littre-bot/src/dictionary"
)

func Test_out(t *testing.T) {
	tests := map[string]struct {
		args   dictionary.Entry
		wantAt string
	}{
		"THORITE": {
			args: dictionary.Entry{
				Term: "THORITE",
				Body: dictionary.EntryBody{
					Def: "Minéral trouvé en Norwége, près Brevig, sur l'île de Loeven",
				},
			},
			wantAt: `THORITE
Minéral trouvé en Norwége, près Brevig, sur l'île de Loeven`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			at := &bytes.Buffer{}
			out(tt.args, at)
			gotAt := at.String()
			if !strings.Contains(gotAt, tt.wantAt) {
				t.Errorf("out() = %v, want %v", gotAt, tt.wantAt)
			}
		})
	}
}
