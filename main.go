package main

import (
	"flag"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

const FONT_FILE = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

type Size = struct {
	wi int
	hi int
}
type BgFn = func(Size, *gg.Context, []colorful.Color)

func drawBg(size Size, dc *gg.Context, palette []colorful.Color) {
	dc.SetColor(palette[0])
	dc.DrawRectangle(0, 0, float64(size.wi), float64(size.hi))
	dc.Fill()
}
func randFrom(p []colorful.Color) colorful.Color {
	return p[rand.Intn(len(p))]
}
func drawBgRect(size Size, dc *gg.Context, palette []colorful.Color) {
	wi, hi := float64(size.wi), float64(size.hi)
	tile := 15.0
	x := 0.0
	y := 0.0
	for x < wi {
		for y < hi {
			dc.SetColor(randFrom(palette))
			dc.DrawRectangle(float64(x), float64(y), tile, tile)
			dc.Fill()
			y += tile
		}
		y = 0
		x += tile
	}
}

type Texts = [2]string

func drawT(dc *gg.Context, text string, fontSize float64, textMaxWidth float64) {
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
	dc.SetColor(blended)
	dc.DrawRectangle(0, 0, w, bannerHeight)
	dc.Fill()

	fs := sizeToFontSize(size)
	dc.SetColor(white)
	dc.Translate(w/10, fs[0]/2)
	drawT(dc, texts[0], fs[0], 0.8*w)
	dc.Translate(0, 1.5*fs[0])
	drawT(dc, texts[1], fs[1], 0.8*w)
	dc.Pop()

}
func draw(size Size, bgFn BgFn, texts Texts) *gg.Context {
	w, h := size.wi, size.hi
	dc := gg.NewContext(w, h)

	cols, err := colorful.WarmPalette(100)
	if err != nil {
		cols = []colorful.Color{colorful.WarmColor(), colorful.WarmColor()}
	}

	bgFn(Size{w, h}, dc, cols)
	textDraw(Size{w, h}, dc, cols, texts)
	return dc
}

func getInput() (Size, Texts) {
	widthP := flag.Int("width", 800, "width of the resulting image")
	heightP := flag.Int("height", 600, "height of the resulting image")
	textP := flag.String("text", "My blogpost", "text to display in the image")
	subtextP := flag.String("subtext", "this time about really important things", "explanatory text to display in the image below the text")
	flag.Parse()
	return Size{*widthP, *heightP}, Texts{*textP, *subtextP}

}

func sizeToFontSize(size Size) []float64 {
	fontsizePrimary := float64(size.wi) / 15
	fontsizeSecondary := fontsizePrimary * 0.5
	return []float64{fontsizePrimary, fontsizeSecondary}
}

func main() {
	rand.Seed(2)
	size, texts := getInput()
	draw(size, drawBgRect, texts).SavePNG("out.png")
}
