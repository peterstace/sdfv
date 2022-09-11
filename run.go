package main

import (
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
)

func run(pxWide, pxHigh int, filename string, samples int, debug bool, scn scene) error {
	acc := newAccumulator(pxWide, pxHigh)
	rng := rand.New(rand.NewSource(0))
	eng := &engine{acc, scn.cam, scn.fn, scn.sk, rng, samples, false}
	eng.renderFrame()
	if err := writeAccAsImage(acc, filename, false); err != nil {
		return err
	}

	if debug {
		acc := newAccumulator(pxWide, pxHigh)
		rng := rand.New(rand.NewSource(0))
		eng := &engine{acc, scn.cam, scn.fn, scn.sk, rng, 1, true}
		eng.renderFrame()
		if err := writeAccAsImage(acc, debugNormsFilename(filename), true); err != nil {
			return err
		}
	}

	return nil
}

func debugNormsFilename(filename string) string {
	ext := filepath.Ext(filename)
	return filename[:len(ext)-1] + "_debug_norms" + ext
}

func writeAccAsImage(acc *accumulator, filename string, raw bool) error {
	img := acc.image(raw)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return f.Close()
}
