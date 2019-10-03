package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var art *sdl.Texture
var delta float64

func main() {
	// set up everything
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
		world.draw(renderer)

		// draw players
		tank.update()
		tank.element.draw(renderer)
		// log tank every second
		tank.element.print(time.Second)

		// draw bullets
		for _, bullet := range bullets { // bullets comes from the bullet class
			if bullet != nil {
				bullet.update()
				bullet.element.draw(renderer)
				bullet.element.print(time.Second)
			}
		}

		// present everything
		renderer.Present()
		delta = time.Since(beginningOfFrame).Seconds() * 1000
	}
}
