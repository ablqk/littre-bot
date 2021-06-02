package dictionary

import (
	"testing"
)

func TestDictionary_NewRandomWord(t *testing.T) {
	tests := map[string]struct {
		fields []Entry
		want   string
	}{
		"one word": {
			fields: []Entry{
				{Term: "MONOGRAMME", Body: EntryBody{Def: "La monogramme, la plus petite fougère connue, et en même temps la plus simple d'organisation."}},
			},
			want: "MONOGRAMME",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := New(tt.fields)
			if got := d.NewRandomWord(); got.Term != tt.want {
				t.Errorf("NewRandomWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
