package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var (
		pxWide   = flag.Int("px-wide", 640, "width of the output image in pixels")
		pxHigh   = flag.Int("px-high", 360, "height of the output image in pixels")
		filename = flag.String("output-filename", "out.png", "output filename")
	)
	flag.Parse()

	if err := run(*pxWide, *pxHigh, *filename); err != nil {
		log.Printf("error running: %v", err)
		os.Exit(1)
	}
}
