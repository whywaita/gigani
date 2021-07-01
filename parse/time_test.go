package parse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNormalizeTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   bool
	}{
		{
			input: `10/7(月) 23:59〜`,
			want:  "10/7(月) 23:59",
			err:   false,
		},
		{
			input: `10/7(月) 24:00〜`,
			want:  "10/8(火) 0:00",
			err:   false,
		},
		{
			input: `10/7(土) 23:59〜`,
			want:  "10/7(土) 23:59",
			err:   false,
		},
		{
			input: `10/7(土) 24:30〜`,
			want:  "10/8(日) 0:30",
			err:   false,
		},
		{
			input: `10/31(火) 24:00〜`,
			want:  "11/1(水) 0:00",
			err:   false,
		},
	}

	for _, test := range tests {
		got, err := NormalizeTime(test.input, "2019")
		if err != nil && !test.err {
			t.Fatalf("NormalizeTime got error: %+v", err)
		}

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("NormalizeTime() mismatch (-want +got):\n%s", diff)
		}
	}
}
