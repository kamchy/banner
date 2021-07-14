package main

import (
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

// Draws bacground using patternDraw function and both texts on canvas of size [w x h]
func Draw(pc PatternContext, texts Texts, patternDraw BgFn) {
	patternDraw(pc)
	textDraw(pc, texts)
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
