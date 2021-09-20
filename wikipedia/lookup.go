package wikipedia

// Do I remember where/how the CSV file that drives this was generated? Of course I don't...

import (
	"context"
	"fmt"
	"github.com/sfomuseum/go-csvdict"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"github.com/sfomuseum/go-sfomuseum-airlines/data"
	"io"
	_ "log"
	"strings"
	"sync"
)

var lookup_table *sync.Map
var lookup_init sync.Once
var lookup_init_err error

type WikipediaLookup struct {
	airlines.Lookup
}

func init() {
	ctx := context.Background()
	airlines.RegisterLookup(ctx, "wikipedia", NewLookup)
}

func NewLookup(ctx context.Context, uri string) (airlines.Lookup, error) {

	lookup_func := func() {

		fs := data.FS
		fh, err := fs.Open("wikipedia.csv")

		if err != nil {
			lookup_init_err = fmt.Errorf("Failed to load data, %v", err)
		}

		defer fh.Close()

		r, err := csvdict.NewReader(fh)

		if err != nil {
			lookup_init_err = err
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
				lookup_init_err = err
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

	if lookup_init_err != nil {
		return nil, lookup_init_err
	}

	l := WikipediaLookup{}
	return &l, nil
}

func (l *WikipediaLookup) Find(ctx context.Context, code string) ([]interface{}, error) {

	pointers, ok := lookup_table.Load(code)

	if !ok {
		return nil, fmt.Errorf("Code '%s' not found", code)
	}

	airlines := make([]interface{}, 0)

	for _, p := range pointers.([]string) {

		if !strings.HasPrefix(p, "pointer:") {
			return nil, fmt.Errorf("Invalid pointer '%s'", p)
		}

		row, ok := lookup_table.Load(p)

		if !ok {
			return nil, fmt.Errorf("Invalid pointer '%s'", p)
		}

		airlines = append(airlines, row.(*Airline))
	}

	return airlines, nil
}

func (l *WikipediaLookup) Append(ctx context.Context, data interface{}) error {
	return fmt.Errorf("Not implemented")
}
