package main

import (
	"fmt"
	"github.com/faiface/pixel"
	// "github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"
)

const (
	winWidth, winHeight float64 = 600, 900
	numPlats            int     = 20
	platDistance        float64 = 80
)

var plats []*platform

func initPlats() {
	plats = make([]*platform, 0)

	for i := 0; i < numPlats; i++ {
		plats = append(plats, randomPlatform(float64(i)*platDistance))
	}
}

func run() {
	initPlats()
	cfg := pixelgl.WindowConfig{
		Title:  "Jump!",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Println("Somethting went wrong")
		panic(err)
	}

	for !win.Closed() {
		win.Update()
		// win.Clear(colornames.Black)
		for i := 0; i < numPlats; i++ {
			plat := plats[i]
			plat.render(win)
		}
	}
}

func main() {
	pixelgl.Run(run)
}
