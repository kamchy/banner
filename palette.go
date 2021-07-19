package banner

import (
	"fmt"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

type Palette = []colorful.Color
type PaletteType = int

const (
	Warm    PaletteType = iota
	Happy   PaletteType = iota
	Hue     PaletteType = iota
	Unknown PaletteType = iota
)

func DefaultPalette(t int, n int) Palette {
	var single = func() colorful.Color {
		return colorful.Hsl(float64(rand.Intn(360)), 0.5, 0.5)
	}
	switch t {
	case Warm:
		single = colorful.WarmColor
	case Happy:
		single = colorful.HappyColor
	}
	res := make([]colorful.Color, 3, 3)
	for idx, _ := range res {
		res[idx] = single()
	}
	return res
}

type PaletteGenerator = func(int) (Palette, error)
type PaletteInfo struct {
	Generator PaletteGenerator
	Desc      string
}

func descriptionsPI(vals map[PaletteType]PaletteInfo) string {
	var s = "\n"
	for key, val := range vals {
		s += fmt.Sprintf("%d -> %s\n", key, val.Desc)
	}
	return s
}

var PaletteInfos = map[PaletteType]PaletteInfo{
	Warm:  PaletteInfo{colorful.WarmPalette, "Warm"},
	Happy: PaletteInfo{colorful.HappyPalette, "Happy"},
	Hue:   PaletteInfo{HuePalette, "Hue"},
}

func HuePalette(n int) (Palette, error) {
	return huePaletteGenerator(50), nil
}
func huePaletteGenerator(inhue int) Palette {
	var angle byte = 30
	var delta byte = 10
	var hue = byte(inhue)
	cols := make([]colorful.Color, angle*2/delta)

	for h := hue - angle; h <= hue+angle; h += angle {
		cols = append(cols, colorful.Hsl(float64(h), 0.5, 0.5))
	}
	return cols

}
func GenPaletteOf(t PaletteType, n int) Palette {
	info, is := PaletteInfos[t]
	if !is {
		return DefaultPalette(t, n)
	}
	pal, err := info.Generator(n)
	if err != nil {
		return DefaultPalette(t, n)
	}
	return pal

}
