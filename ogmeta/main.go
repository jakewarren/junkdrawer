package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/jakewarren/metascraper"
)

//TODO: add json output a la https://github.com/johnreutersward/opengraph

func main() {
	flag.Parse()
	p, err := metascraper.Scrape(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	print(p)

}

func print(p *metascraper.Page) {
	for _, m := range p.MetaData() {
		switch {
		case len(m.Property) > 0:
			fmt.Printf("%s = %s\n", color.YellowString(m.Property), m.Content)
		case len(m.Name) > 0:
			fmt.Printf("%s = %s\n", color.YellowString(m.Name), m.Content)
		}
	}
}
