package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/pflag"
	"github.com/vanng822/go-premailer/premailer"
	"gopkg.in/mail.v2"
)

// TODO: add readme
// TODO: add support for FROM env variable

type config struct {
	file        string
	attachments []string
	subject     string
	from        string
	to          []string
	cc          []string
	bcc         []string
	dryRun      bool // don't actually send email, just simulate it
	pre         bool // encapsulate output in <pre> tag
	inline      bool // run through premailer to inline style info
}

func (c config) sendEmail(input io.Reader) {
	d := mail.NewDialer("localhost", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	rawInput, _ := ioutil.ReadAll(input)
	msgBody := string(rawInput)

	m := mail.NewMessage()
	if len(c.from) > 0 {
		m.SetHeader("From", c.from)
	}
	m.SetHeader("To", c.to...)
	if len(c.cc) > 0 {
		m.SetHeader("CC", c.cc...)
	}
	if len(c.bcc) > 0 {
		m.SetHeader("BCC", c.bcc...)
	}
	m.SetHeader("Subject", c.subject)

	// process attachments
	for _, a := range c.attachments {
		m.Attach(a)
	}

	// inline style info
	if c.inline {
		pm, _ := premailer.NewPremailerFromString(msgBody, premailer.NewOptions())
		inlinedOutput, pmErr := pm.Transform()
		if pmErr != nil {
			log.Fatal(pmErr)
		}
		msgBody = inlinedOutput
	}

	// wrap with pre elements
	if c.pre {
		msgBody = "<pre>\n" + msgBody + "\n</pre>"
	}

	// wrap with basic html elements
	if !strings.HasPrefix(msgBody, "<html><body>") && !strings.HasPrefix(msgBody, "<!DOCTYPE html") {
		msgBody = "<html><body>\n" + msgBody + "\n</body></html>"
	}

	m.SetBody("text/html", msgBody)

	if c.dryRun {
		spew.Dump(m)
		fmt.Println("Message body:")
		fmt.Println(msgBody)
		os.Exit(0)
	}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	os.Exit(0)
}

func main() {
	var c config
	pflag.StringVar(&c.file, "file", "", "read input from a file")
	pflag.StringSliceVarP(&c.attachments, "attach", "a", []string{}, "files to attach to the email")
	pflag.StringVarP(&c.subject, "subject", "s", "", "email subject")
	pflag.StringVarP(&c.from, "from", "f", "", "email address to send from (defaults to system default)")
	pflag.StringSliceVarP(&c.to, "to", "t", []string{}, "addresses to send to <To>")
	pflag.StringSliceVar(&c.cc, "cc", []string{}, "addresses to send cc <CC>")
	pflag.StringSliceVar(&c.bcc, "bcc", []string{}, "addresses to send bcc <BCC>")
	pflag.BoolVarP(&c.dryRun, "dry-run", "n", false, "dry run, don't actually send email")
	pflag.BoolVar(&c.pre, "pre", true, "wrap body with a <pre> element")
	pflag.BoolVar(&c.inline, "inline", false, "inline style information via premailer")
	pflag.Parse()

	// ensure required parameters are provided by the user
	ensureProvided(&c)

	if len(c.file) > 0 {
		f, err := os.Open(c.file)
		if err != nil {
			log.Fatal(err)
		}
		c.sendEmail(f)
	}

	c.sendEmail(os.Stdin)

}

func ensureProvided(c *config) {

	errOut := func(msg string) {
		fmt.Fprintf(os.Stderr, msg)
		os.Exit(1)
	}

	switch {
	case len(c.subject) == 0:
		errOut("subject not provided\n")
	case len(c.to) == 0:
		errOut("to not provided\n")
	case len(c.from) == 0:
		errOut("from not provided\n")
	}
}
