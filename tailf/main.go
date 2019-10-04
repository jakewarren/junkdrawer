package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/AndriiChuzhynov/tail"
	"github.com/arsham/rainbow/rainbow"
	ninetail "github.com/gongo/9t"
	"github.com/jakewarren/hr"
	"github.com/snaggen/keyboard"
	"github.com/spf13/pflag"
	"github.com/traviscampbell/sinebow"
)

func tailFile(file string) {
	t, err := tail.TailFile(file, tail.Config{Follow: true, Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}, Logger: tail.DiscardingLogger})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func tailStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func printMarker() {
	if ct := os.Getenv("COLORTERM"); ct == "truecolor" || ct == "24bit" {
		fmt.Print(sinebow.RainbroBG(hr.HorizontalRule("-")))
	} else {
		hr := hr.HorizontalRule("-")
		in := strings.NewReader(hr)
		l := rainbow.Light{
			Reader: in,
			Writer: os.Stdout,
		}
		l.Paint()
	}
}

func main() {
	pflag.Parse()

	// HACK: running into issues where the keyboard events don't respond, so setting a SIGINT handler as a backup
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for _ = range signals {
			fmt.Println("\nReceived an interrupt, stopping...")
			os.Exit(0)
		}
	}()

	// NOTE: this whole section is pretty gross.
	// couldn't find one library to do everything so using 3 different libraries to handle the various use cases

	if pflag.NArg() < 1 {
		// check if stdin is attached
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeNamedPipe != 0 {
			go tailStdin()
		} else {
			// exit out
			fmt.Fprintln(os.Stderr, "No file or stdin provided")
			os.Exit(1)
		}
	}

	runner, err := ninetail.Runner(pflag.Args(), ninetail.Config{Colorize: true})
	if err != nil {
		log.Fatal(err)
	}

	if pflag.NArg() > 1 {
		go runner.Run()
	} else if pflag.NArg() > 0 {
		go tailFile(pflag.Arg(0))
	}

	err = keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, keyErr := keyboard.GetKey()
		if keyErr != nil {
			panic(keyErr)
		} else if key == keyboard.KeyEsc || char == 'q' || key == keyboard.KeyCtrlC {
			if pflag.NArg() > 1 {
				runner.Stop()
			}
			break
		} else if char == 'h' {
			printMarker()
		}

		// clear the screen
		if char == 'c' || (key == keyboard.KeyCtrlL) {
			fmt.Print("\033[H\033[2J")
		}
	}
}
