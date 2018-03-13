package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	// "golang.org/x/image/colornames"
	"image/color"
	"math/rand"
)

const (
	platHeight float64 = 10
	minWidth   int     = 50
	maxWidth   int     = 200
)

type platform struct {
	pos pixel.Vec
	dim pixel.Vec
	col color.Color
}

func newPlatform(position pixel.Vec, width float64, color color.Color) *platform {
	return &platform{
		pos: position,
		dim: pixel.V(width, platHeight),
		col: color,
	}
}

func randomPlatform(posY float64) *platform {
	rectWidth := float64(rand.Intn(maxWidth-minWidth) + minWidth)
	rectHeight := float64(platHeight)

	// Genertes x value from 0 to winWidth - rectWidth
	posX := float64(rand.Intn(int(winWidth - rectWidth)))
	randColor := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(255)}

	return &platform{
		pos: pixel.V(posX, posY),
		dim: pixel.V(rectWidth, rectHeight),
		col: randColor,
	}
}

func (p platform) getX() float64 {
	return p.pos.X
}

func (p platform) getY() float64 {
	return p.pos.Y
}

func (p platform) getWidth() float64 {
	return p.dim.X
}

func (p platform) getHeight() float64 {
	return p.dim.Y
}

func (p platform) render(t pixel.Target) {
	imd := imdraw.New(nil)
	imd.Reset()
	imd.Color = p.col
	imd.Push(pixel.V(p.pos.X, p.pos.Y))
	imd.Push(pixel.V(p.pos.X+p.dim.X, p.pos.Y))
	imd.Push(pixel.V(p.pos.X+p.dim.X, p.pos.Y+p.dim.Y))
	imd.Push(pixel.V(p.pos.X, p.pos.Y+p.dim.Y))
	imd.Polygon(0)
	imd.Draw(t)
}

// This controlls the platform's print statement
func (p platform) String() string {
	return fmt.Sprintf("===========Platform==========\nPosition:   (%v, %v)\nDimensions: %vx%v\nColor:      %v\n\n\n", p.pos.X, p.pos.Y, p.dim.X, p.dim.Y, p.col)
}
