package main

import (
	"context"
	"log"
	"os"
	"time"

	cdp "github.com/knq/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var val1, val2, val3 string
	err = c.Run(ctxt, sendkeys(&val1, &val2, &val3))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("#input1 value: %s", val1)
	log.Printf("#textarea1 value: %s", val2)
	log.Printf("#input2 value: %s", val3)
}

func sendkeys(val1, val2, val3 *string) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("file:" + os.Getenv("GOPATH") + "/src/github.com/knq/chromedp/testdata/visible.html"),
		cdp.WaitVisible(`#input1`, cdp.ByID),
		cdp.WaitVisible(`#textarea1`, cdp.ByID),
		cdp.SendKeys(`#textarea1`, "\b\b\n\naoeu\n\ntest1\n\nblah2\n\n\t\b\bother box!\t\ntest4", cdp.ByID),
		cdp.Value(`#input1`, val1, cdp.ByID),
		cdp.Value(`#textarea1`, val2, cdp.ByID),
		cdp.SetValue(`#input2`, "test3", cdp.ByID),
		cdp.Value(`#input2`, val3, cdp.ByID),
		cdp.Sleep(30 * time.Second),
	}
}
