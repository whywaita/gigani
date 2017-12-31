package parse

import (
	"strings"
	"testing"
)

func TestNormalizeTime(t *testing.T) {
	sample1 := `10/7(月) 23:59〜`
	sample2 := `10/7(月) 24:00〜`
	sample3 := `10/7(土) 23:59〜`
	sample4 := `10/7(土) 24:30〜`

	p1, err := NormalizeTime(sample1)
	p2, err := NormalizeTime(sample2)
	p3, err := NormalizeTime(sample3)
	p4, err := NormalizeTime(sample4)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.EqualFold(p1, "10/7(月) 23:59") {
		t.Fatalf("%s is fail", sample1)
	}

	if !strings.EqualFold(p2, "10/8(火) 0:00") {
		t.Fatalf("%s is fail", sample2)
	}

	if !strings.EqualFold(p3, "10/7(土) 23:59") {
		t.Fatalf("%s is fail", sample3)
	}

	if !strings.EqualFold(p4, "10/8(日) 0:30") {
		t.Fatalf("%s is fail", sample4)
	}

}
