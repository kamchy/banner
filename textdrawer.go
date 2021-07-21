package banner

import (
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

type Texts [2]*string

func (ts Texts) atLeastOne() bool {
	res := ts[0] != nil || ts[1] != nil
	return res
}

func DrawT(dc *gg.Context, t Text, textMaxWidth float64) {
	dc.LoadFontFace(FONT_FILE, t.fontSize)
	dc.DrawStringWrapped(t.t, 0, 0, 0, 0, textMaxWidth, t.spacing, gg.AlignLeft)
}

var white, _ = colorful.MakeColor(color.White)

func textStripeBlended(base colorful.Color) (float64, float64, float64) {
	blended := base.BlendHcl(white, 0.2)
	return blended.FastLinearRgb()
}

func calcSize(dc *gg.Context, fontSize float64, width float64, t string, spacing float64) Size {
	dc.LoadFontFace(FONT_FILE, fontSize)
	lines := dc.WordWrap(t, width)
	w, h := dc.MeasureMultilineString(strings.Join(lines, "\n"), spacing)
	return Size{w, h}

}

// Text with rendering info
type Text struct {
	// string
	t string
	// font size
	fontSize float64
	// spacing
	spacing float64
	// color
	col color.Color
}

// Calculates Size for Text when given maz width
func (s Text) Size(dc *gg.Context, width float64) Size {
	dc.LoadFontFace(FONT_FILE, s.fontSize)
	lines := dc.WordWrap(s.t, width)
	w, h := dc.MeasureMultilineString(strings.Join(lines, "\n"), s.spacing)
	return Size{w, h}
}

func makeText(texts Texts, fs []float64) []Text {
	tts := make([]Text, 0)
	for idx, text := range texts {
		if text != nil {
			tt := Text{t: *text, fontSize: fs[idx], spacing: 1.8, col: color.White}
			tts = append(tts, tt)
		}
	}
	return tts
}

func drawBanner(c PatternContext, bannerHeight float64) {
	if bannerHeight == 0 {
		return
	}
	dc := c.dc
	dc.Push()
	r, g, b := textStripeBlended(randFrom(c.p))
	dc.SetRGBA(r, g, b, 0.9)
	dc.DrawRectangle(0, 0, c.size.wi, bannerHeight)
	dc.Fill()
	dc.Pop()
}

func calcTextHeight(c PatternContext, tts []Text, width float64) float64 {
	var hi float64
	for idx, t := range tts {
		hi += t.Size(c.dc, width).hi
		if idx < len(tts)-1 {
			hi += (t.spacing - 1) * t.fontSize
		}
	}
	if len(tts) > 0 {
	}
	return hi
}

func textDraw(c PatternContext, texts Texts) {
	size, dc, _ := c.size, c.dc, c.p
	tSizes := sizeToFontSize(size)
	textWidth := size.wi * 0.8
	tts := makeText(texts, tSizes)
	textHeight := calcTextHeight(c, tts, textWidth)
	textOffset := Point{(size.wi - textWidth) / 2, (size.hi - textHeight) / 2}
	textSize := Size{textWidth, textHeight}

	dc.Push()
	dc.Translate(0.0, textOffset.y)
	bannerHeight := textSize.hi + 0.1*size.hi
	drawBanner(c, bannerHeight)

	dc.Translate(textOffset.x, (bannerHeight-textHeight)/2)
	for _, t := range tts {
		dc.SetColor(t.col)
		DrawT(dc, t, textWidth)
		dc.Translate(0, t.Size(dc, textWidth).hi+(t.spacing-1)*t.fontSize)
	}
	dc.Pop()

}

// Takes size of whole canvas and determines size of font
// for both primary and secondary text
func sizeToFontSize(size Size) []float64 {
	fontsizePrimary := float64(size.wi) / 15
	fontsizeSecondary := fontsizePrimary * 0.6
	return []float64{fontsizePrimary, fontsizeSecondary}
}
