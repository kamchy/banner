package banner

import (
	"flag"
	"fmt"
	"os"
)

const DEF_WIDTH = 800
const DEF_HEIGHT = 600
const DEF_WITH_TEXT = false
const DEF_TITLE = "My blogpost"
const DEF_SUB = "this time about really important things"
const DEF_TILE = 30
const DEF_ALG = 5
const DEF_OUT = "out.png"
const DEF_PAL = Warm

type DescProvider interface {
	Desc() string
}

func (a Alg) Desc() string {
	return a.desc
}

// Generates description of available pattern-drawing algorithms
// used in help message
func descriptions(vals map[int]DescProvider) string {
	var s = "\n"
	for key, val := range vals {
		s += fmt.Sprintf("%d -> %s\n", key, val.Desc())
	}
	return s
}

type Input struct {
	W        *int
	H        *int
	Texts    []*string
	AlgIdx   *int
	TileSize *float64
	Pt       *int
	OutName  *string
}

func InputFlagSet() (*flag.FlagSet, Input) {
	fs := flag.NewFlagSet(
		"input",
		flag.ExitOnError)

	widthP := fs.Int("w", DEF_WIDTH, "width of the resulting image")
	heightP := fs.Int("h", DEF_HEIGHT, "height of the resulting image")
	textP := fs.String("t", "", "text to display in the image")
	subtextP := fs.String("st", "", "explanatory text to display in the image below the text")
	painterP := fs.Int("alg", DEF_ALG, fmt.Sprintf("Background painter algorithm; valid values are: %v", descriptionsPA(PainterAlgs)))
	tileSizeP := fs.Float64("ts", DEF_TILE, "size of tile")
	outPalP := fs.Int("p", DEF_PAL, fmt.Sprintf("palette type; valid values are: %v", descriptionsPI(PaletteInfos)))
	outNameP := fs.String("o", DEF_OUT, "name of output file where banner in .png format will be saved")
	textPs := []*string{textP, subtextP}
	inp := Input{widthP, heightP, textPs, painterP, tileSizeP, outPalP, outNameP}
	return fs, inp
}

// Parses commandline parameters and gives
// Input object with fields initalized with provided values (or defaults)
// TODO - sanitize input
func GetInput() Input {
	fs, inp := InputFlagSet()
	fs.Parse(os.Args[1:])
	return inp
}
