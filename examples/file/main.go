package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rafalmp/link"
)

func main() {
	fileName := flag.String("file", "", "A HTML file to parse")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", links)
}
