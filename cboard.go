package cboard

import (
	"image"
	"image/color"
)

var (
	ColorDraw       = color.Black
	ColorBackGround = color.White
	ColorGrid       = color.RGBA{R: 238, G: 238, B: 238, A: 255}
)

func NewCB(w, h int) *CB {
	matrix := make([][]color.Color, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]color.Color, w)
	}
	return &CB{
		Rect:            image.Rect(0, 0, w, h),
		Matrix:          matrix,
		ColorDraw:       ColorDraw,
		ColorBackGround: ColorBackGround,
		ColorGrid:       ColorGrid,
	}
}

type CB struct {
	Rect            image.Rectangle
	Matrix          [][]color.Color
	ColorDraw       color.Color
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

func (cb *CB) DrawLineX(y, x0, x1 int) {
	if x1 >= x0 {
		for x := x0; x <= x1; x++ {
			cb.Set(x, y, cb.ColorDraw)
		}
	} else {
		for x := x0; x >= x1; x-- {
			cb.Set(x, y, cb.ColorDraw)
		}
	}
}

func (cb *CB) DrawLineY(x, y0, y1 int) {
	if y1 >= y0 {
		for y := y0; y <= y1; y++ {
			cb.Set(x, y, cb.ColorDraw)
		}
	} else {
		for y := y0; y >= y1; y-- {
			cb.Set(x, y, cb.ColorDraw)
		}
	}
}

func (cb *CB) DrawRect(x0, y0, x1, y1 int) {
	r := image.Rect(x0, y0, x1, y1)
	for x := r.Min.X; x <= r.Max.X; x++ {
		for y := r.Min.Y; y <= r.Max.Y; y++ {
			cb.Set(x, y, cb.ColorDraw)
		}
	}
}

func (cb *CB) DrawRectBorder(x0, y0, x1, y1 int) {
	r := image.Rect(x0, y0, x1, y1)
	cb.DrawLineY(r.Min.X, r.Max.Y, r.Min.Y)
	cb.DrawLineX(r.Min.Y, r.Min.X, r.Max.X)
	cb.DrawLineY(r.Max.X, r.Min.Y, r.Max.Y)
	cb.DrawLineX(r.Max.Y, r.Max.X, r.Min.X)
}
