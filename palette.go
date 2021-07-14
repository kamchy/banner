package main

import "github.com/lucasb-eyer/go-colorful"

type Palette = []colorful.Color
type PaletteType int

const (
	Warm PaletteType = iota
	Happy
)

type PaletteGenerator interface {
	Generate(int) Palette
}

func (t PaletteType) Generate(n int) Palette {
	var p func(int) (Palette, error)
	switch t {
	case Warm:
		p = colorful.WarmPalette
	case Happy:
		p = colorful.HappyPalette
	}
	ret, err := p(n)
	if err != nil {
		ret = []colorful.Color{colorful.WarmColor(), colorful.WarmColor()}
	}
	return ret
}
func GenPaletteOf(t PaletteType, n int) Palette {
	return t.Generate(n)
}

// Generates palette for pattern and text bg
func GenPalette() []colorful.Color {
	return GenPaletteOf(Warm, 10)
}
