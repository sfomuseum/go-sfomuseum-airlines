package sfomuseum

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"github.com/sfomuseum/go-sfomuseum-airlines/data"
	"io"
	"strconv"
	"strings"
	"sync"
)

var lookup_table *sync.Map
var lookup_init sync.Once
var lookup_init_err error

type SFOMuseumLookupFunc func()

type SFOMuseumLookup struct {
	airlines.Lookup
}

// NewLookup will return an `airlines.Lookup` instance derived from precompiled (embedded) data in `data/sfomuseum.json`
func NewLookup() (airlines.Lookup, error) {

	fs := data.FS
	fh, err := fs.Open("sfomuseum.json")

	if err != nil {
		return nil, fmt.Errorf("Failed to load data, %v", err)
	}

	lookup_func := NewLookupFuncWithReader(fh)
	return NewLookupWithLookupFunc(lookup_func)
}

// NewLookup will return an `SFOMuseumLookupFunc` function instance that, when invoked, will populate an `airlines.Lookup` instance with data stored in `r`.
// `r` will be closed when the `SFOMuseumLookupFunc` function instance is invoked.
// It is assumed that the data in `r` will be formatted in the same way as the procompiled (embedded) data stored in `data/sfomuseum.json`.
func NewLookupFuncWithReader(r io.ReadCloser) SFOMuseumLookupFunc {

	lookup_func := func() {

		defer r.Close()

		var airline []*Airline

		dec := json.NewDecoder(r)
		err := dec.Decode(&airline)

		if err != nil {
			lookup_init_err = err
			return
		}

		table := new(sync.Map)

		for idx, craft := range airline {

			pointer := fmt.Sprintf("pointer:%d", idx)
			table.Store(pointer, craft)

			str_wofid := strconv.FormatInt(craft.WOFID, 10)

			possible_codes := []string{
				craft.IATACode,
				craft.ICAOCode,
				craft.ICAOCallsign,
				str_wofid,
			}

			for _, code := range possible_codes {

				if code == "" {
					continue
				}

				pointers := make([]string, 0)
				has_pointer := false

				others, ok := table.Load(code)

				if ok {

					pointers = others.([]string)
				}

				for _, dupe := range pointers {

					if dupe == pointer {
						has_pointer = true
						break
					}
				}

				if has_pointer {
					continue
				}

				pointers = append(pointers, pointer)
				table.Store(code, pointers)
			}

			idx += 1
		}

		lookup_table = table
	}

	return lookup_func
}

// NewLookupWithLookupFunc will return an `airlines.Lookup` instance derived by data compiled using `lookup_func`.
func NewLookupWithLookupFunc(lookup_func SFOMuseumLookupFunc) (airlines.Lookup, error) {

	lookup_init.Do(lookup_func)

	if lookup_init_err != nil {
		return nil, lookup_init_err
	}

	l := SFOMuseumLookup{}
	return &l, nil
}

func (l *SFOMuseumLookup) Find(code string) ([]interface{}, error) {

	pointers, ok := lookup_table.Load(code)

	if !ok {
		return nil, errors.New("Not found")
	}

	airline := make([]interface{}, 0)

	for _, p := range pointers.([]string) {

		if !strings.HasPrefix(p, "pointer:") {
			return nil, errors.New("Invalid pointer")
		}

		row, ok := lookup_table.Load(p)

		if !ok {
			return nil, errors.New("Invalid pointer")
		}

		airline = append(airline, row.(*Airline))
	}

	return airline, nil
}
