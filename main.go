package main

import (
	"flag"
	"fmt"
    "os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [secret]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("secret is missing.")
		os.Exit(1)
	}
	code, err := Generate2FACode(args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("2FA Code: %s\n", code)
	}
}
