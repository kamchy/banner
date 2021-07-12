package main

import (
	"flag"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

const FONT_FILE = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

type Size = struct {
	wi float64
	hi float64
}

func DrawRect(size Size, dc *gg.Context, col color.Color) {
	dc.SetColor(col)
	dc.DrawRectangle(0, 0, size.wi, size.hi)
	dc.Fill()
}

type BgFn = func(Size, *gg.Context, []colorful.Color)

func DrawBgLines(size Size, dc *gg.Context, palette []colorful.Color) {
	DrawRect(size, dc, colorful.FastLinearRgb(1, 1, 1))
	base := math.Min(size.wi, size.hi)
	linew := base / 50
	space := linew * 2
	lineLenMin, lineLenMax := base/40, base/2
	diff := lineLenMax - lineLenMin

	dc.SetLineCapRound()
	dc.SetLineWidth(linew)
	x := 0.0
	y := 0.0
	for y < size.hi {
		for x < size.wi {
			dc.SetColor(randFrom(palette[1:]))
			linewi := lineLenMin + diff*rand.Float64()
			dc.DrawLine(x, y, x+linewi, y)
			dc.Stroke()
			x += linewi + space
		}
		x = 0
		y += linew * 2.0
	}
}

// Single tile painter: draws concentric circles
func DrawBgHexagons(size Size, dc *gg.Context, palette []colorful.Color) {
	DrawRect(size, dc, colorful.Hsl(33.0, 0.2, 0.7))
	s := math.Min(size.wi, size.hi)
	dc.SetColor(randFrom(palette[1:]))
	dc.DrawRegularPolygon(6, s/2, s/2, s/2, 30.0*2*math.Pi/360.0)
	dc.Fill()
}

// Single tile painter: draws concentric circles
func DrawBgCircles(size Size, dc *gg.Context, palette []colorful.Color) {
	DrawRect(size, dc, colorful.FastLinearRgb(1, 1, 1))
	for _, rmul := range []float64{0.8, 0.6, 0.4, 0.2} {
		dc.SetColor(randFrom(palette[1:]))
		dc.DrawPoint(size.wi/2, size.hi/2, math.Min(size.wi, size.hi)*rmul/2)
		dc.Fill()
	}
}

// Draws rect of size with first color of the palette
func DrawBgPlain(size Size, dc *gg.Context, palette []colorful.Color) {
	DrawRect(size, dc, palette[0])
}

// Draws rect of size with randomly picked color from the palette
func DrawBgRandomRectColor(size Size, dc *gg.Context, palette []colorful.Color) {
	DrawRect(size, dc, randFrom(palette))
}

// Generates random colorful.Color from given array
func randFrom(p []colorful.Color) colorful.Color {
	return p[rand.Intn(len(p))]
}

// Draws all tiles using tilePainter for single tile
func DrawBgRectWithPainter(size Size, dc *gg.Context, palette []colorful.Color, tilePainter Painter) {
	wi, hi := size.wi, size.hi
	tile := tilePainter.tilesize
	x := 0.0
	y := 0.0
	for x < wi {
		for y < hi {
			dc.Push()
			dc.Translate(x, y)
			tilePainter.fn(Size{tile, tile}, dc, palette)
			dc.Pop()
			y += tile
		}
		y = 0
		x += tile
	}
}

type Texts = [2]string

func DrawT(dc *gg.Context, text string, fontSize float64, textMaxWidth float64) {
	dc.LoadFontFace(FONT_FILE, fontSize)
	dc.DrawStringWrapped(text, 0, 0, 0, 0, textMaxWidth, 1.8, gg.AlignLeft)
}

func textDraw(size Size, dc *gg.Context, pal []colorful.Color, texts Texts) {
	w, h := float64(size.wi), float64(size.hi)
	bannerHeight := h / 3
	dc.Push()
	dc.Translate(0, 0.5*(h-bannerHeight))
	white, _ := colorful.MakeColor(color.White)
	blended := randFrom(pal).BlendHcl(white, 0.2)
	r, g, b := blended.FastLinearRgb()
	dc.SetRGBA(r, g, b, 0.9)
	dc.DrawRectangle(0, 0, size.wi, bannerHeight)
	dc.Fill()

	fs := sizeToFontSize(size)
	dc.SetColor(white)
	dc.Translate(w/10, fs[0]/2)
	DrawT(dc, texts[0], fs[0], 0.8*w)
	dc.Translate(0, 1.5*fs[0])
	DrawT(dc, texts[1], fs[1], 0.8*w)
	dc.Pop()

}

// Draws bacground using bgFn function and both texts on canvas of size [w x h]
func Draw(w int, h int, texts Texts, bgFn BgFn) *gg.Context {
	rand.Seed(time.Now().UnixNano())
	dc := gg.NewContext(w, h)

	cols, err := colorful.WarmPalette(100)
	if err != nil {
		cols = []colorful.Color{colorful.WarmColor(), colorful.WarmColor()}
	}

	var canvasSize = Size{float64(w), float64(h)}
	bgFn(canvasSize, dc, cols)
	textDraw(canvasSize, dc, cols, texts)
	return dc
}

// Represents drawing function and the size of square tile
type Painter struct {
	fn       BgFn
	tilesize float64
}

// struct representing drawing algorithm
// fn - function BgFn
// desc - description displayed when -h option was given
// tileOnly - if true fn is applied to single tile,
//            else fn is applied to whole canvas
type Alg struct {
	fn       BgFn
	desc     string
	tileOnly bool
}

// Generates background drawing function that uses given  painter, i.e. tile-drawing function and tile size
func tiledBgFn(p Painter) BgFn {
	return func(canvasSize Size, dc *gg.Context, cols []colorful.Color) {
		DrawBgRectWithPainter(canvasSize, dc, cols, p)
	}
}

// array of identifiers presented to the user to Alg struct
var painterAlgs = []Alg{
	{DrawBgRandomRectColor, "random rectangles", true},
	{DrawBgPlain, "plain color", false},
	{DrawBgCircles, "random concentric circles", true},
	{DrawBgLines, "random horizontal lines", false},
	{DrawBgHexagons, "random hexagons", true},
}

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
// function of type BgFn that can draw on given size, using given context and palette
// name of .png output file
func GetInput(painterAlgs []Alg) (int, int, Texts, BgFn, string) {

	const DEF_WIDTH = 800
	const DEF_HEIGHT = 600
	const DEF_TITLE = "My blogpost"
	const DEF_SUB = "this time about really important things"
	const DEF_TILE = 30
	const DEF_ALG = 4
	const DEF_OUT = "out.png"

	var makeBgFn = func(pindex *int, tilesize *float64, widthP *int) BgFn {
		if *pindex < 0 || *pindex > len(painterAlgs) {
			*pindex = 0
		}
		if *tilesize < 0 || *tilesize > float64(*widthP) {
			*tilesize = DEF_TILE
		}
		alg := painterAlgs[*pindex]
		if alg.tileOnly {
			return tiledBgFn(Painter{alg.fn, *tilesize})
		} else {
			return alg.fn
		}
	}

	widthP := flag.Int("width", DEF_WIDTH, "width of the resulting image")
	heightP := flag.Int("height", DEF_HEIGHT, "height of the resulting image")
	textP := flag.String("text", DEF_TITLE, "text to display in the image")
	subtextP := flag.String("subtext", DEF_SUB, "explanatory text to display in the image below the text")
	painterP := flag.Int("alg", DEF_ALG, fmt.Sprintf("Background painter algorithm; valid values are: %v", descriptions(painterAlgs)))
	tileSizeP := flag.Float64("ts", DEF_TILE, "size of tile")
	outNameP := flag.String("outName", DEF_OUT, "name of output file where banner in .png format will be saved")
	flag.Parse()
	return *widthP, *heightP, Texts{*textP, *subtextP}, makeBgFn(painterP, tileSizeP, widthP), *outNameP
}

// Takes size of whole canvas and determines size of font
// for both primary and secondary text
func sizeToFontSize(size Size) []float64 {
	fontsizePrimary := float64(size.wi) / 15
	fontsizeSecondary := fontsizePrimary * 0.6
	return []float64{fontsizePrimary, fontsizeSecondary}
}

func main() {
	w, h, texts, bgFn, outName := GetInput(painterAlgs)
	Draw(w, h, texts, bgFn).SavePNG(outName)
}
