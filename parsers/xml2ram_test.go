package parsers

import (
	"testing"

	"github.com/ablqk/littre-bot/src/dictionary"
)

func Test_parseLetter(t *testing.T) {
	tests := map[string]struct {
		file string
		want []dictionary.Entry
	}{
		"x": {
			file: "testdata/x.xml",
			want: []dictionary.Entry{
				{Term: "X"},
				{Term: "XANTHE"},
				{Term: "XANTHOPROTÉIQUE"},
				{Term: "XYLOFER"},
				{Term: "XYLOGÉNE"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseLetter(tt.file)
			if err != nil {
				t.Errorf("error while parsing %s", err.Error())
			}
			if len(got) != len(tt.want) {
				t.Errorf("different sizes : %v, want %v", len(got), len(tt.want))
			}
		})
	}
}
