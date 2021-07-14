package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

type Texts = [2]string

func DrawT(dc *gg.Context, text string, fontSize float64, textMaxWidth float64) {
	dc.LoadFontFace(FONT_FILE, fontSize)
	dc.DrawStringWrapped(text, 0, 0, 0, 0, textMaxWidth, 1.8, gg.AlignLeft)
}

var white, _ = colorful.MakeColor(color.White)

func textStripeBlended(base colorful.Color) (float64, float64, float64) {
	blended := base.BlendHcl(white, 0.2)
	return blended.FastLinearRgb()
}

// FIXME - text layout
func textDraw(c PatternContext, texts Texts) {
	size, dc, pal := c.size, c.dc, c.p
	w, h := float64(size.wi), float64(size.hi)
	bannerHeight := h / 3
	dc.Push()
	dc.Translate(0, 0.5*(h-bannerHeight))
	r, g, b := textStripeBlended(randFrom(pal))
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

// Takes size of whole canvas and determines size of font
// for both primary and secondary text
func sizeToFontSize(size Size) []float64 {
	fontsizePrimary := float64(size.wi) / 15
	fontsizeSecondary := fontsizePrimary * 0.6
	return []float64{fontsizePrimary, fontsizeSecondary}
}
