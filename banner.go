package banner

import (
	"fmt"
	"math/rand"

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

// struct representing drawing algorithm
// fn - function BgFn
// desc - description displayed when -h option was given
// pg - function that generates Rects for drawing tile
type Alg struct {
	fn   BgFn
	desc string
	pg   RectGenerator
}

// Generates random colorful.Color from given array
func randFrom(p []colorful.Color) colorful.Color {
	return p[rand.Intn(len(p))]
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

type AlgType = int

const (
	RandomRect AlgType = iota
	RandomRectOffset
	PlainColor
	ConcentricCircles
	ConcentricCirclesOffset
	HorizontalLines
	RandomHexagons
	RandomHexagonsOffset
)

func descriptionsPA(vals map[AlgType]Alg) string {
	var s = "\n"
	for key, val := range vals {
		s += fmt.Sprintf("%d -> %s\n", key, val.Desc())
	}
	return s
}

// array of identifiers presented to the user to Alg struct
var PainterAlgs = map[AlgType]Alg{
	RandomRect:              {DrawRectRand, "random rectangles", gridGenerator},
	RandomRectOffset:        {DrawRectRand, "random rectangles with offset", gridDeltaGenerator},
	PlainColor:              {DrawRect, "plain color", plainGenerator},
	ConcentricCircles:       {DrawBgCircles, "concentric circles", gridGenerator},
	ConcentricCirclesOffset: {DrawBgCircles, "concentric circles offset", gridDeltaGenerator},
	HorizontalLines:         {DrawBgLines, "random horizontal lines", linesRandomGenerator},
	RandomHexagons:          {DrawHexagon, "random hexagons", gridGenerator},
	RandomHexagonsOffset:    {DrawHexagon, "random hexagons with offset", gridDeltaGenerator},
}

// Draws with pc as PatternContext, filling background with patternDraw and using Textx.
func Draw(pc PatternContext, texts Texts, patternDraw BgFn) {
	patternDraw(pc)
	textDraw(pc, texts)
}

func GenerateBanner(i Input) {
	wi, hi, paletteType, algType, tileSize, outName, texts :=
		*i.W, *i.W, *i.Pt, *i.AlgIdx, *i.TileSize, *i.OutName, Texts{*i.Texts[0], *i.Texts[1]}

	drawContext := gg.NewContext(wi, hi)
	var canvasSize = Size{float64(wi), float64(hi)}
	cc := PatternContext{canvasSize, drawContext, GenPaletteOf(paletteType, 10)}
	Draw(cc, texts, generatingFn(PainterAlgs[algType], tileSize))
	cc.dc.SavePNG(outName)

}
