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

	schemes := []string{
		"sfomuseum://",
		"sfomuseum://github",
		// Leaving as an example because it depends on github.com/whosonfirst/go-whosonfirst-iterate-git
		// which we don't need to make a requirement for this package
		// "sfomuseum://iterator?uri=git%3A%2F%2F%3Finclude%3Dproperties.sfomuseum%3Aplacetype%3Dairline&source=https://github.com/sfomuseum-data/sfomuseum-data-enterprise.git,		
	}
	
	ctx := context.Background()

	for _, s := range schemes {
		
		lu, err := airlines.NewLookup(ctx, s)
		
		if err != nil {
			t.Fatalf("Failed to create lookup for '%s', %v", s, err)
		}
		
		for code, wofid := range wofid_tests {
			
			results, err := lu.Find(ctx, code)
			
			if err != nil {
				t.Fatalf("Unable to find '%s' using scheme '%s', %v", code, s, err)
			}
			
			if len(results) != 1 {
				t.Fatalf("Invalid results for '%s' using scheme '%s'", code, s)
			}
			
			a := results[0].(*Airline)
			
			if a.WOFID != wofid {
				t.Fatalf("Invalid match for '%s', expected %d but got %d using scheme '%s'", code, wofid, a.WOFID, s)
			}
		}
	}
	
}
