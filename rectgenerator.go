package main

import (
	"math"
	"math/rand"
)

//
// Function type: calculation of dx and dy
//
// For given column and row number in a grid returns dx and dy
// used to calculate top-left position of a next tile
type DeltaCalculation = func(col int, row int) (float64, float64)

// Returns slice of Rect representing rectangular locations
// for tile drawing operation
func iteration(p Point, s Size, ts Size, fn DeltaCalculation) []Rect {
	ox := p.x
	oy := p.y
	dx, dy := 0.0, 0.0
	res := make([]Rect, 0, 1)
	wi, hi := math.Max(s.wi, 0.0), math.Max(s.hi, 0.0)
	col, row := 0, 0
	for oy < hi {
		for ox < wi {
			dx, dy = fn(col, row)
			res = append(res, Rect{Point{ox, oy}, Size{dx, dy}})
			ox += dx
			col++
		}
		ox = p.x
		oy += dy
		col = 0
		row++
	}
	return res
}

// Generates slice of Rects for regular grid.
func gridGenerator(p Point, s Size, ts Size) []Rect {
	lg := func(int, int) (float64, float64) { return ts.wi, ts.hi }
	return iteration(p, s, ts, lg)
}

// Generates slice of Rects on a grid where every second row is translated by half width of a tile
func gridDeltaGenerator(p Point, s Size, ts Size) []Rect {
	lg := func(col int, row int) (float64, float64) {
		if col == 0 && row%2 == 0 {
			return -ts.wi / 2, ts.hi
		} else {
			return ts.wi, ts.hi
		}
	}
	return iteration(p, s, ts, lg)
}

// Generates one-element slice with a Rect of size s and top left at p
func plainGenerator(p Point, s Size, _ Size) []Rect {
	return []Rect{{p, s}}
}

func lineLenGenerator(cs Size, widthHPart float64, minPart float64, maxPart float64) DeltaCalculation {
	base := math.Min(cs.wi, cs.hi)
	linew := base * widthHPart
	space := linew * 2
	lineLenMin, lineLenMax := base*minPart, base*maxPart
	diff := lineLenMax - lineLenMin
	return func(col int, row int) (float64, float64) {
		return lineLenMin + diff*rand.Float64(), space
	}

}
func linesRandomGenerator(p Point, s Size, ts Size) []Rect {
	lg := lineLenGenerator(s, 0.015, 0.03, 0.3)
	return iteration(p, s, ts, lg)
}
