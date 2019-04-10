package airlines

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"io"
	_ "log"
	"strings"
	"sync"
)

var lookup_table *sync.Map
var lookup_init sync.Once

type Lookup struct {
	// we used to store the lookup_table here but then switched to sync.Once-ing
}

func NewLookup() (*Lookup, error) {

	var lookup_err error

	lookup_func := func() {

		fh := bytes.NewReader([]byte(lookupTable))

		r, err := csv.NewDictReader(fh)

		if err != nil {
			lookup_err = err
			return
		}

		table := new(sync.Map)
		idx := 0

		for {

			row, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				lookup_err = err
				return
			}

			a := Airline{
				IATA:      row["iata_code"],
				ICAO:      row["icao_code"],
				TELEPHONY: row["callsign"],
				Name:      row["name"],
			}

			if strings.HasSuffix(row["iata_code"], "*") {
				a.Duplicate = true
				a.IATA = strings.Replace(a.IATA, "*", "", 1)
			}

			pointer := fmt.Sprintf("pointer:%d", idx)

			table.Store(pointer, &a)

			possible_codes := []string{
				a.IATA,
				a.ICAO,
				a.TELEPHONY,
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

	lookup_init.Do(lookup_func)

	if lookup_err != nil {
		return nil, lookup_err
	}

	l := Lookup{}
	return &l, nil
}

func (l *Lookup) Find(code string) ([]*Airline, error) {

	pointers, ok := lookup_table.Load(code)

	if !ok {
		return nil, errors.New("Not found")
	}

	airlines := make([]*Airline, 0)

	for _, p := range pointers.([]string) {

		if !strings.HasPrefix(p, "pointer:") {
			return nil, errors.New("Invalid pointer")
		}

		row, ok := lookup_table.Load(p)

		if !ok {
			return nil, errors.New("Invalid pointer")
		}

		airlines = append(airlines, row.(*Airline))
	}

	return airlines, nil
}
