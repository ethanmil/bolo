package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var renderer *sdl.Renderer
var art *sdl.Texture
var delta float64

func main() {
	// set up the window, renderer, & main texture
	sdlSetup()
	// defer the destruction of window, renderer
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()

	// set up the world map
	world := newWorldMap(vector{x: 50, y: 50}, 1)

	// set up players
	tank := newTank()

	// game loop
	running := true
	for running {
		beginningOfFrame := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				art.Destroy()
				println("Quit")
				running = false
				break
			}
		}

		// draw the world
		world.draw()

		// draw players
		tank.update()
		tank.element.draw()

		// draw bullets
		for _, bullet := range bullets { // bullets comes from the bullet class
			if bullet != nil {
				bullet.update()
				bullet.element.draw()
			}
		}

		// present everything
		renderer.Present()
		delta = time.Since(beginningOfFrame).Seconds() * 1000
	}
}

func sdlSetup() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err = sdl.CreateWindow("bolo", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1200, 800, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	art = newTexture(renderer, "images/art.bmp")
}
