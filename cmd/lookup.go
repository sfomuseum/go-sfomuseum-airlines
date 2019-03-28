package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-airlines"
	"log"
)

func main() {

	flag.Parse()

	a, err := airlines.NewLookup()

	if err != nil {
		log.Fatal(err)
	}

	for _, code := range flag.Args() {

		airlines, err := a.Find(code)

		if err != nil {
			log.Fatal(err)
		}

		for _, a := range airlines {
			fmt.Println(a, a.Name)
		}
	}

}
