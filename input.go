package main

import (
	"flag"
	"fmt"
)

// Generates description of available pattern-drawing algorithms
// used in help message
func descriptions(algs []Alg) string {
	var s = "\n"
	for idx, alg := range algs {
		s += fmt.Sprintf("%d -> %s\n", idx, alg.desc)
	}
	return s
}

type Input struct {
	w        int
	h        int
	texts    Texts
	algIdx   int
	tileSize float64
	outName  string
}

// Parses commandline parameters and gives
// width - width of resulting image
// height - height of resulting image
// Texts - primary and secondary
// name of .png output file
// TODO - sanitize input
func GetInput(painterAlgs []Alg) Input {

	const DEF_WIDTH = 800
	const DEF_HEIGHT = 600
	const DEF_TITLE = "My blogpost"
	const DEF_SUB = "this time about really important things"
	const DEF_TILE = 30
	const DEF_ALG = 5
	const DEF_OUT = "out.png"

	widthP := flag.Int("width", DEF_WIDTH, "width of the resulting image")
	heightP := flag.Int("height", DEF_HEIGHT, "height of the resulting image")
	textP := flag.String("text", DEF_TITLE, "text to display in the image")
	subtextP := flag.String("subtext", DEF_SUB, "explanatory text to display in the image below the text")
	painterP := flag.Int("alg", DEF_ALG, fmt.Sprintf("Background painter algorithm; valid values are: %v", descriptions(painterAlgs)))
	tileSizeP := flag.Float64("ts", DEF_TILE, "size of tile")
	outNameP := flag.String("outName", DEF_OUT, "name of output file where banner in .png format will be saved")
	flag.Parse()

	return Input{*widthP, *heightP, Texts{*textP, *subtextP}, *painterP, *tileSizeP, *outNameP}
}
