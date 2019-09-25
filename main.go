package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var art *sdl.Texture

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("bolo", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1200, 800, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	art = newTexture(renderer, "images/art.bmp")

	tank := newTank()

	// err = window.UpdateSurface()
	// if err != nil {
	// 	panic(err)
	// }

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				art.Destroy()
				println("Quit")
				running = false
				break
			}
		}

		err = renderer.SetDrawColor(255, 255, 255, 255)
		if err != nil {
			panic(err)
		}
		err = renderer.Clear()
		if err != nil {
			panic(err)
		}

		tank.update()
		tank.element.draw(renderer)

		renderer.Present()
	}
}
