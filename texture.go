package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newTexture(renderer *sdl.Renderer, path string) *sdl.Texture {
	img, err := sdl.LoadBMP(path)
	if err != nil {
		println("errr")
	}
	defer img.Free()

	t, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		println("errr")
	}

	return t
}
