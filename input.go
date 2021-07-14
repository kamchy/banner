package banner

import (
	"flag"
	"fmt"
	"os"
)

const DEF_WIDTH = 800
const DEF_HEIGHT = 600
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
func descriptions(vals []DescProvider) string {
	var s = "\n"
	for idx, val := range vals {
		s += fmt.Sprintf("%d -> %s\n", idx, val.Desc())
	}
	return s
}

// https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces/12754757#12754757
func makeInterfaceAlg(aa []Alg) []DescProvider {
	b := make([]DescProvider, len(aa), len(aa))
	for i := range aa {
		b[i] = aa[i]
	}
	return b
}
func makeInterfacePaletteType(aa []PaletteType) []DescProvider {
	b := make([]DescProvider, len(aa), len(aa))
	for i := range aa {
		b[i] = aa[i]
	}
	return b
}

type Input struct {
	w        *int
	h        *int
	texts    []*string
	algIdx   *int
	tileSize *float64
	pt       *int
	outName  *string
}

func InputFlagSet() (*flag.FlagSet, Input) {
	fs := flag.NewFlagSet("input", flag.ContinueOnError)

	widthP := fs.Int("width", DEF_WIDTH, "width of the resulting image")
	heightP := fs.Int("height", DEF_HEIGHT, "height of the resulting image")
	textP := fs.String("text", DEF_TITLE, "text to display in the image")
	subtextP := fs.String("subtext", DEF_SUB, "explanatory text to display in the image below the text")
	painterP := fs.Int("alg", DEF_ALG, fmt.Sprintf("Background painter algorithm; valid values are: %v", descriptions(makeInterfaceAlg(painterAlgs))))
	tileSizeP := fs.Float64("ts", DEF_TILE, "size of tile")
	outPalP := fs.Int("palette", int(DEF_PAL), fmt.Sprintf("palette type; valid values are: %v", descriptions(makeInterfacePaletteType(paletteGenerators))))
	outNameP := fs.String("outName", DEF_OUT, "name of output file where banner in .png format will be saved")
	inp := Input{widthP, heightP, []*string{textP, subtextP}, painterP, tileSizeP, outPalP, outNameP}
	return fs, inp
}

// Parses commandline parameters and gives
// width - width of resulting image
// height - height of resulting image
// Texts - primary and secondary
// name of .png output file
// TODO - sanitize input
func GetInput() Input {
	fs, inp := InputFlagSet()
	fs.Parse(os.Args)
	return inp
}
