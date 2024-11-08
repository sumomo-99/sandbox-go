package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s BINARY_NUMBER\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	dec, err := strconv.ParseInt(flag.Arg(0), 2, 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing binary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(dec)
}
