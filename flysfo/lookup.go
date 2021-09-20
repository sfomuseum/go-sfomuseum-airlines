package flysfo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"github.com/sfomuseum/go-sfomuseum-airlines/data"
	"io"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var lookup_table *sync.Map
var lookup_idx int64

var lookup_init sync.Once
var lookup_init_err error

type FlysfoLookupFunc func(context.Context)

type FlysfoLookup struct {
	airlines.Lookup
}

func init() {
	ctx := context.Background()
	airlines.RegisterLookup(ctx, "flysfo", NewLookup)

	lookup_idx = int64(0)
}

// NewLookup will return an `airlines.Lookup` instance derived from precompiled (embedded) data in `data/flysfo.json`
func NewLookup(ctx context.Context, uri string) (airlines.Lookup, error) {

	fs := data.FS
	fh, err := fs.Open("flysfo.json")

	if err != nil {
		return nil, fmt.Errorf("Failed to load data, %v", err)
	}

	lookup_func := NewLookupFuncWithReader(ctx, fh)
	return NewLookupWithLookupFunc(ctx, lookup_func)
}

// NewLookup will return an `FlysfoLookupFunc` function instance that, when invoked, will populate an `airlines.Lookup` instance with data stored in `r`.
// `r` will be closed when the `FlysfoLookupFunc` function instance is invoked.
// It is assumed that the data in `r` will be formatted in the same way as the procompiled (embedded) data stored in `data/flysfo.json`.
func NewLookupFuncWithReader(ctx context.Context, r io.ReadCloser) FlysfoLookupFunc {

	lookup_func := func(ctx context.Context) {

		defer r.Close()

		var airline []*Airline

		dec := json.NewDecoder(r)
		err := dec.Decode(&airline)

		if err != nil {
			lookup_init_err = err
			return
		}

		table := new(sync.Map)

		for _, data := range airline {

			select {
			case <-ctx.Done():
				return
			default:
				// pass
			}

			appendData(ctx, table, data)
		}

		lookup_table = table
	}

	return lookup_func
}

// NewLookupWithLookupFunc will return an `airlines.Lookup` instance derived by data compiled using `lookup_func`.
func NewLookupWithLookupFunc(ctx context.Context, lookup_func FlysfoLookupFunc) (airlines.Lookup, error) {

	fn := func() {
		lookup_func(ctx)
	}

	lookup_init.Do(fn)

	if lookup_init_err != nil {
		return nil, lookup_init_err
	}

	l := FlysfoLookup{}
	return &l, nil
}

func (l *FlysfoLookup) Find(ctx context.Context, code string) ([]interface{}, error) {

	pointers, ok := lookup_table.Load(code)

	if !ok {
		msg := fmt.Sprintf("code '%s' not found", code)
		return nil, errors.New(msg)
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

func (l *FlysfoLookup) Append(ctx context.Context, data interface{}) error {
	return appendData(ctx, lookup_table, data.(*Airline))
}

func appendData(ctx context.Context, table *sync.Map, data *Airline) error {

	idx := atomic.AddInt64(&lookup_idx, 1)

	pointer := fmt.Sprintf("pointer:%d", idx)
	table.Store(pointer, data)

	str_wofid := strconv.FormatInt(data.WOFID, 10)

	possible_codes := []string{
		data.IATACode,
		data.ICAOCode,
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

	return nil
}
