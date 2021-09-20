package sfomuseum

import (
	"context"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"testing"
)

func TestSFOMuseumLookup(t *testing.T) {

	wofid_tests := map[string]int64{
		"AC":         1159283597,
		"ACA":        1159283597,
		"AIR CANADA": 1159283597,
		"MOV":        1360700753,
		"NN":         1360700753,
		"77":         1159283643,
		"AHC":        1159283643,
	}

	ctx := context.Background()

	lu, err := airlines.NewLookup(ctx, "sfomuseum://")

	if err != nil {
		t.Fatalf("Failed to create lookup, %v", err)
	}

	for code, wofid := range wofid_tests {

		results, err := lu.Find(ctx, code)

		if err != nil {
			t.Fatalf("Unable to find '%s', %v", code, err)
		}

		if len(results) != 1 {
			t.Fatalf("Invalid results for '%s'", code)
		}

		a := results[0].(*Airline)

		if a.WOFID != wofid {
			t.Fatalf("Invalid match for '%s', expected %d but got %d", code, wofid, a.WOFID)
		}
	}
}
