package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Set usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
	}

	// Parse the command line arguments
	flag.Parse()

	argc := flag.NArg()

	if argc != 1 {
		flag.Usage()
		os.Exit(1)
	}

	program := flag.Args()[0]

	if !strings.HasSuffix(program, ".bf") {
		log.Fatalf("%s does not to be a brainfuck program\n", program)
	}

	instructions, err := ioutil.ReadFile(program)
	if err != nil {
		log.Fatalf("There was a problem when opening %s\n", program)
	}
}
