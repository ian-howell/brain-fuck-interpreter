package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func run(steps []byte) {
	ticker := make([]byte, 30000)
	sp := 0
	for i := 0; i < len(steps); i++ {
		switch steps[i] {
		case '+':
			ticker[sp]++
		case '-':
			ticker[sp]--
		case '>':
			if sp == 29999 {
				log.Fatal("ERROR: walked off right of array\n")
			}
			sp++
		case '<':
			if sp == 0 {
				log.Fatal("ERROR: walked off left of array\n")
			}
			sp--
		case '.':
			fmt.Printf("%c", ticker[sp])
		case ',':
			fmt.Scanf("%c", &ticker[sp])
		case '[':
			if ticker[sp] == 0 {
				bal := 1
				for bal != 0 {
					i++
					if steps[i] == '[' {
						bal++
					} else if steps[i] == ']' {
						bal--
					}
				}
			}
		case ']':
			if ticker[sp] != 0 {
				bal := 1
				for bal != 0 {
					i--
					if steps[i] == ']' {
						bal++
					} else if steps[i] == '[' {
						bal--
					}
				}
			}
		}
	}
}

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
		log.Fatalf("%s does not appear to be a brainfuck program\n", program)
	}

	instructions, err := ioutil.ReadFile(program)
	if err != nil {
		log.Fatalf("There was a problem when opening %s\n", program)
	}

	run(instructions)
}
