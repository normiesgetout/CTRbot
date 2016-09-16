// main
package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")
	url := flag.String("scrape", "", "scrape 4chan API for cancer \"\"")
	flag.Parse()

	buf := new(bytes.Buffer)

	if *url != "" {
		Scrape(url, buf)
	} else {

		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	c := NewChain(*prefixLen)
	c.Build(buf)
	text := c.Generate(*numWords)
	fmt.Println(text)
}
