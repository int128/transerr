package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
)

// check returns nil if s is a positive number.
func check(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return errors.Wrapf(err, "invalid number")
	}
	if n < 0 {
		// comment should be kept
		return errors.Errorf("number is negative: %d", n)
	}
	if n == 0 {
		return errors.New("number is zero")
	}
	return nil
}

// main is an entry point.
func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s NUMBER", os.Args[0])
	}
	if err := check(os.Args[1]); err != nil {
		log.Printf("error: %+v", err)
	}
}
