package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

const FONT_FILE = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

type BgFn = func(PatternContext)

// Size struct with wi (width) and hi (height) of drawing area
type Size = struct {
	wi float64
	hi float64
}

// p - top left corner of canvas
// s - size of canvas
// ts - tile size
type RectGenerator = func(p Point, s Size, tileSize Size) []Rect

// Point struct, represents location inside drawing area
type Point struct {
	x float64
	y float64
}

// Represents rectangle used to draw a tile
type Rect struct {
	tl Point
	s  Size
}

// Generates random colorful.Color from given array
func randFrom(p []colorful.Color) colorful.Color {
	return p[rand.Intn(len(p))]
}

// struct representing drawing algorithm
// fn - function BgFn
// desc - description displayed when -h option was given
// pg - function that generates Rects for drawing tile
type Alg struct {
	fn   BgFn
	desc string
	pg   RectGenerator
}

// Generates function that calculate rects, iterates over them
// and calls alg.fn on each.
func generatingFn(alg Alg, tilesize float64) BgFn {
	return func(c PatternContext) {
		DrawRect(c)
		dc := c.dc
		rects := alg.pg(Point{}, c.size, Size{tilesize, tilesize})
		for _, r := range rects {
			dc.Push()
			dc.Translate(r.tl.x, r.tl.y)
			alg.fn(c.withSize(r.s))
			dc.Pop()
		}

	}
}

// array of identifiers presented to the user to Alg struct
var painterAlgs = []Alg{
	{DrawRectRand, "random rectangles", gridGenerator},
	{DrawRectRand, "random rectangles with offset", gridDeltaGenerator},
	{DrawRect, "plain color", plainGenerator},
	{DrawBgCircles, "concentric circles", gridGenerator},
	{DrawBgCircles, "concentric circles offset", gridDeltaGenerator},
	{DrawBgLines, "random horizontal lines", linesRandomGenerator},
	{DrawHexagon, "random hexagons", gridGenerator},
	{DrawHexagon, "random hexagons with offset", gridDeltaGenerator},
}

// Generates description of available pattern-drawing algorithms
// used in help message
func descriptions(algs []Alg) string {
	var s = "\n"
	for idx, alg := range algs {
		s += fmt.Sprintf("%d -> %s\n", idx, alg.desc)
	}
	return s
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

// Takes size of whole canvas and determines size of font
// for both primary and secondary text
func sizeToFontSize(size Size) []float64 {
	fontsizePrimary := float64(size.wi) / 15
	fontsizeSecondary := fontsizePrimary * 0.6
	return []float64{fontsizePrimary, fontsizeSecondary}
}

// Draws bacground using patternDraw function and both texts on canvas of size [w x h]
func Draw(pc PatternContext, texts Texts, patternDraw BgFn) {
	patternDraw(pc)
	textDraw(pc, texts)
}

// Generates palette for pattern and text bg
func GenPalette() []colorful.Color {
	cols, err := colorful.WarmPalette(100)
	if err != nil {
		cols = []colorful.Color{colorful.WarmColor(), colorful.WarmColor()}
	}
	return cols
}

type Input struct {
	w        int
	h        int
	texts    Texts
	algIdx   int
	tileSize float64
	outName  string
}

func GenerateBanner(i Input) {
	drawContext := gg.NewContext(i.w, i.h)
	var canvasSize = Size{float64(i.w), float64(i.h)}
	cc := PatternContext{canvasSize, drawContext, GenPalette()}
	Draw(cc, i.texts, generatingFn(painterAlgs[i.algIdx], i.tileSize))
	cc.dc.SavePNG(i.outName)

}
func main() {
	rand.Seed(time.Now().UnixNano())
	i := GetInput(painterAlgs)
	GenerateBanner(i)
}
