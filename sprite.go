package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type sprite struct {
	size  vector
	chunk sdl.Rect

	// log helper
	lastLogged time.Time
}
