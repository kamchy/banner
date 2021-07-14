package main

import "github.com/lucasb-eyer/go-colorful"

type Palette = []colorful.Color
type PaletteType int

func fromIntToPaletteType(v int) PaletteType {
	switch v {
	case 0:
		return Warm
	case 1:
		return Happy
	default:
		return Unknown
	}
}

const (
	Warm PaletteType = iota
	Happy
	Unknown
)

var paletteGenerators []PaletteType = []PaletteType{Warm, Happy}

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

func (t PaletteType) Desc() string {
	switch t {
	case Warm:
		return "Warm"
	case Happy:
		return "Happy"
	default:
		return "Unknown"
	}
}

func GenPaletteOf(t PaletteGenerator, n int) Palette {
	return t.Generate(n)
}
