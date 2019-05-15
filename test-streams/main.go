package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/spf13/pflag"
)

func main() {

	count := pflag.IntP("count", "c", -1, "number of items to print (-1 to print forever)")
	pflag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := *count; i != 0; i-- {
			fmt.Fprint(os.Stdout, ">>> stdout\n")
			time.Sleep(1 * time.Second)
			fmt.Fprint(os.Stderr, ">>> stderr\n")
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Printf("<<< %q\n", scanner.Text())
		}
	}()

	wg.Wait()

}
