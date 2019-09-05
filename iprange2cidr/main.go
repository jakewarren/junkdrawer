package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/EvilSuperstars/go-cidrman"
)

func main() {
	flag.Parse()

	ips := regexp.MustCompile("[ -]+").Split(strings.Join(flag.Args(), " "), -1)

	cidrs, err := cidrman.IPRangeToCIDRs(ips[0], ips[1])
	if err != nil {
		log.Fatalln(err)
	}

	for _, c := range cidrs {
		fmt.Println(c)
	}
}
