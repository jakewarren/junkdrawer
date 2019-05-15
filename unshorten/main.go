package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	traceFlag := flag.Bool("trace", false, "display verbose trace information")

	flag.Parse()

	switch *traceFlag {
	case false:
		resp, _ := http.Get(flag.Args()[0])

		fmt.Println(resp.Request.URL.String())

	case true:
		trace(flag.Args()[0])
	}
}

// based on https://jonathanmh.com/tracing-preventing-http-redirects-golang/
func trace(url string) {
	nextURL := url
	var i int
	for i < 100 {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		resp, err := client.Get(nextURL)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("[%d] %s\n", resp.StatusCode, resp.Request.URL)

		if resp.StatusCode == 200 {
			break
		} else {
			nextURL = resp.Header.Get("Location")
			i += 1
		}
	}

}
