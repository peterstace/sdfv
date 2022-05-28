package main

import (
	"fmt"
	"os"

	"github.com/peterstace/sdfv/engine"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func mainE() error {
	args := engine.RunArgs{}
	buf, err := engine.Run(args)
	if err != nil {
		return err
	}
	if err := os.WriteFile("out.gif", buf, 0644); err != nil {
		return err
	}
	return nil
}
