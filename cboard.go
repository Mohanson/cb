package main

import (
	"image"
	"image/color"
)

var (
	ColorDraw       = color.Black
	ColorBackGround = color.White
	ColorGrid       = color.RGBA{R: 238, G: 238, B: 238, A: 255}
	GridW           = 20
	GridH           = 20
	GridD           = 4
)

func NewCB(w, h int) *CB {
	Pix := make([][]color.Color, h)
	for i := 0; i < h; i++ {
		Pix[i] = make([]color.Color, w)
	}
	return &CB{
		W:               w,
		H:               h,
		GridW:           GridW,
		GridH:           GridH,
		GridD:           GridD,
		ColorDraw:       ColorDraw,
		ColorBackGround: ColorBackGround,
		ColorGrid:       ColorGrid,
		Pix:             Pix,
	}
}

type CB struct {
	W               int
	H               int
	GridW           int
	GridH           int
	GridD           int
	ColorDraw       color.Color
	ColorBackGround color.Color
	ColorGrid       color.Color
	Pix             [][]color.Color
}

func (cb *CB) Get(x, y int) color.Color {
	c := cb.Pix[y][x]
	if c == nil {
		return cb.ColorGrid
	}
	return c
}

func (cb *CB) Set(x, y int, c color.Color) {
	cb.Pix[y][x] = c
}

func (cb *CB) Gen() image.Image {
	w := cb.W*(cb.GridW+cb.GridD) + cb.GridD
	h := cb.H*(cb.GridH+cb.GridD) + cb.GridD
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for px := 0; px < w; px++ {
		for py := 0; py < h; py++ {
			m.Set(px, py, cb.ColorBackGround)
		}
	}
	for x := 0; x < cb.W; x++ {
		for y := 0; y < cb.H; y++ {
			color := cb.Get(x, y)
			for px := (cb.GridW+cb.GridD)*x + cb.GridD; px < (cb.GridW+cb.GridD)*(x+1); px++ {
				for py := (cb.GridH+cb.GridD)*y + cb.GridD; py < (cb.GridH+cb.GridD)*(y+1); py++ {
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
