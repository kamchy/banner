package banner

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

type PatternContext struct {
	size Size
	dc   *gg.Context
	p    []colorful.Color
}

func (c PatternContext) withPalette(p []colorful.Color) PatternContext {
	return PatternContext{c.size, c.dc, p}
}

func (c PatternContext) withSize(s Size) PatternContext {
	return PatternContext{s, c.dc, c.p}
}

func rectWith(c PatternContext, colFn func() color.Color) {
	c.dc.SetColor(colFn())
	c.dc.DrawRectangle(0, 0, c.size.wi, c.size.hi)
	c.dc.Fill()
}

// Draws rect of size with first color of the palette
func DrawRect(c PatternContext) {
	rectWith(c, func() color.Color { return c.p[0] })
}

// Draws rect of size with random color of the palette
func DrawRectRand(c PatternContext) {
	rectWith(c, func() color.Color { return randFrom(c.p) })
}

// DrawHexagon draws hexagon
func DrawHexagon(c PatternContext) {
	size, dc, palette := c.size, c.dc, c.p
	DrawRect(c.withPalette([]colorful.Color{colorful.Hsl(330.0, 0.5, 0.7)}))
	s := math.Min(size.wi, size.hi)
	dc.SetColor(randFrom(palette[1:]))
	dc.DrawRegularPolygon(6, s/2, s/2, s/2, 30.0*2*math.Pi/360.0)
	dc.Fill()
	dc.DrawRegularPolygon(6, s/2, s/2, s/2, 30.0*2*math.Pi/360.0)
	dc.Stroke()
}

// here c.size is the size of the rectangle, not the tile
func DrawHexagon2(c PatternContext) {
	size, dc, palette := c.size, c.dc, c.p
	DrawRect(c.withPalette([]colorful.Color{colorful.Hsl(330.0, 0.5, 0.7)}))
	r := size.wi / 3.0
	dc.SetColor(randFrom(palette[1:]))
	dc.DrawRegularPolygon(6, 0, 0, r, 0)
	dc.Fill()
}

// Single tile painter: draws concentric circles
func DrawBgCircles(c PatternContext) {
	DrawRect(c.withPalette([]colorful.Color{colorful.FastLinearRgb(1, 1, 1)}))
	size, dc, palette := c.size, c.dc, c.p
	for _, rmul := range []float64{0.8, 0.6, 0.4, 0.2} {
		dc.SetColor(randFrom(palette[1:]))
		dc.DrawPoint(size.wi/2, size.hi/2, math.Min(size.wi, size.hi)*rmul/2)
		dc.Fill()
	}
}

func DrawBgLines(c PatternContext) {
	//DrawRect(c.withPalette([]colorful.Color{colorful.FastLinearRgb(1, 1, 1)}))
	size, dc, palette := c.size, c.dc, c.p
	dc.SetLineCapRound()
	dc.SetLineWidth(0.8 * size.hi)
	dc.SetColor(randFrom(palette[1:]))
	dc.DrawLine(0, 0, size.wi-size.hi*1.2, 0)
	dc.Stroke()
}
