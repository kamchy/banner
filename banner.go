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
	tl   Point
	size Size
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
			alg.fn(c.withSize(r.size))
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
	RandomHexagonsOffset:    {DrawHexagon2, "random hexagons with offset", gridHexDeltaGenerator},
}

// Draws with pc as PatternContext, filling background with patternDraw and using Textx.
func Draw(pc PatternContext, texts []*string, patternDraw BgFn) {
	patternDraw(pc)
	if texts[0] != nil || texts[1] != nil {
		textDraw(pc, texts)
	}
}

func GenerateBanner(i Input) {
	var v InpData = new(InpData).From(i)
	drawContext := gg.NewContext(v.W, v.H)
	var canvasSize = Size{float64(v.W), float64(v.H)}
	cc := PatternContext{canvasSize, drawContext, GenPaletteOf(v.P, 10)}
	Draw(cc, i.Texts, generatingFn(PainterAlgs[v.Alg], v.Ts))
	cc.dc.SavePNG(v.O)

}
