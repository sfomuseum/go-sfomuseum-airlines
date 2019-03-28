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

type Lookup struct {
	table *sync.Map
}

func NewLookup() (*Lookup, error) {

	fh := bytes.NewReader([]byte(lookupTable))

	r, err := csv.NewDictReader(fh)

	if err != nil {
		return nil, err
	}

	table := new(sync.Map)
	idx := 0

	for {

		row, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		a := Airline{
			IATA:      row["iata_code"],
			ICAO:      row["icao_code"],
			TELEPHONY: row["callsign"],
			Name:      row["name"],
			Duplicate: false,
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

	l := Lookup{
		table: table,
	}

	return &l, nil
}

func (l *Lookup) Find(code string) ([]*Airline, error) {

	pointers, ok := l.table.Load(code)

	if !ok {
		return nil, errors.New("Not found")
	}

	airlines := make([]*Airline, 0)

	for _, p := range pointers.([]string) {

		if !strings.HasPrefix(p, "pointer:") {
			return nil, errors.New("Invalid pointer")
		}

		row, ok := l.table.Load(p)

		if !ok {
			return nil, errors.New("Invalid pointer")
		}

		airlines = append(airlines, row.(*Airline))
	}

	return airlines, nil
}
