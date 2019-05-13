package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"github.com/sfomuseum/go-sfomuseum-airlines/flysfo"	
	"github.com/sfomuseum/go-sfomuseum-airlines/sfomuseum"
	"github.com/sfomuseum/go-sfomuseum-airlines/wikipedia"
	"log"
)

func main() {

	source := flag.String("source", "wikipedia", "Valid sources are: flysfo,sfomuseum,wikipedia.")
	flag.Parse()

	var lookup airlines.Lookup
	var err error

	switch *source {
	case "flysfo":
		lookup, err = flysfo.NewLookup()
	case "sfomuseum":
		lookup, err = sfomuseum.NewLookup()
	case "wikipedia":
		lookup, err = wikipedia.NewLookup()
	default:
		err = errors.New("Unknown source")
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, code := range flag.Args() {

		results, err := lookup.Find(code)

		if err != nil {
			log.Fatal(err)
		}

		for _, a := range results {
			fmt.Println(a)
		}
	}

}
