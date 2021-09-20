package main

import (
	_ "github.com/sfomuseum/go-sfomuseum-airlines/flysfo"
	_ "github.com/sfomuseum/go-sfomuseum-airlines/sfomuseum"
)

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"log"
)

func main() {

	lookup_uri := flag.String("lookup-uri", "sfomuseum://", "...")

	flag.Parse()

	ctx := context.Background()
	lookup, err := airlines.NewLookup(ctx, *lookup_uri)

	if err != nil {
		log.Fatal(err)
	}

	for _, code := range flag.Args() {

		results, err := lookup.Find(ctx, code)

		if err != nil {
			fmt.Printf("%s *** %s\n", code, err)
			continue
		}

		for _, a := range results {
			fmt.Println(a)
		}
	}

}
