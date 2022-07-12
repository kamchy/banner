package banner

import (
	"fmt"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

type Palette = []colorful.Color
type PaletteType = int

const (
	Warm PaletteType = iota
	Happy
	Hue
	Unknown
)

func DefaultPalette(t PaletteType, n int) Palette {
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
	Warm:  {colorful.WarmPalette, "Warm"},
	Happy: {colorful.HappyPalette, "Happy"},
	Hue:   {HuePalette, "Hue"},
}

func HuePalette(n int) (Palette, error) {
	return huePaletteGenerator(n), nil
}
func huePaletteGenerator(inhue int) Palette {
	var angle = 60
	var delta = 20
	var hue = inhue
	cols := make([]colorful.Color, 0)

	for h := hue - angle; h <= hue+angle; h += delta {
		hh := float64((h + 360) % int(360))
		c := colorful.Hsl(hh, 0.36, 0.55)
		cols = append(cols, c)
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
