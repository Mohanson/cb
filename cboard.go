package cboard

import (
	"image"
	"image/color"
)

var (
	ColorBackGround = color.White
	ColorGrid       = color.RGBA{R: 207, G: 111, B: 193, A: 255}
)

func New(w, h int) *CB {
	matrix := make([][]color.Color, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]color.Color, w)
	}
	return &CB{
		Rect:            image.Rect(0, 0, w, h),
		Matrix:          matrix,
		ColorBackGround: ColorBackGround,
		ColorGrid:       ColorGrid,
	}
}

type CB struct {
	Rect            image.Rectangle
	Matrix          [][]color.Color
	ColorBackGround color.Color
	ColorGrid       color.Color
}

func (cb *CB) Get(x, y int) color.Color {
	c := cb.Matrix[y][x]
	if c == nil {
		return cb.ColorGrid
	}
	return c
}

func (cb *CB) Set(x, y int, c color.Color) {
	cb.Matrix[y][x] = c
}

func (cb *CB) Gen() image.Image {
	w := cb.Rect.Dx()*12 + 2
	h := cb.Rect.Dy()*12 + 2
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for px := 0; px < w; px++ {
		for py := 0; py < h; py++ {
			m.Set(px, py, cb.ColorBackGround)
		}
	}
	for x := 0; x < cb.Rect.Dx(); x++ {
		for y := 0; y < cb.Rect.Dy(); y++ {
			color := cb.Get(x, y)
			for px := 12*x + 2; px < 12*(x+1); px++ {
				for py := 12*y + 2; py < 12*(y+1); py++ {
					m.Set(px, py, color)
				}
			}
		}
	}
	return m
}
