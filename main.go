// main
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")
	importFile := flag.String("filename", "trainer.txt", "Place filename between quotes \"\"")
	url := flag.String("scrape", "", "scrape 4chan API for cancer \"\"")
	flag.Parse()

	_, err := os.Stat("trainer.txt")
	if err != nil {
		_, err := os.Create("trainer.txt")
		if err != nil {
			fmt.Println("trainer.txt could not be created! Maybe make your own?")
		} else {
			fmt.Println("trainer.txt created in current directory")
		}
		os.Exit(1)
	}

	file, err := os.OpenFile(*importFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("No file was input!")
		os.Exit(1)
	}

	if *url != "" {
		Scrape(url, file)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	c := NewChain(*prefixLen)
	c.Build(file)
	text := c.Generate(*numWords)
	fmt.Println(text)
}
