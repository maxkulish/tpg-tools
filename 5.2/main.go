package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.NewFlagSet("counter", flag.ContinueOnError)
	countWords := f.Bool("w", false, "Count words instead of lines")
	f.Parse(os.Args[1:])
	if *countWords {
		fmt.Println("We're counting words!")
	}
}