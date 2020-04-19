package main

import (
	"time"

	"github.com/ethanmil/bolo/bullet"
	"github.com/ethanmil/bolo/maps"
	"github.com/ethanmil/bolo/tank"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten"
	"github.com/veandco/go-sdl2/sdl"
)

var delta float64

func main() {
	bolo := NewBolo()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Your game's title")

	err := ebiten.RunGame(bolo)
	if err != nil {
		log.Fatal(err)
	}

	// set up the world map
	world := maps.NewWorldMap("./assets/test_map.txt", 1)

	// set up players
	tank := tank.NewTank()

	// TODO - User bullet.Manager here
	// set up bullets
	bullets := make([]*bullet.Bullet, 50)

	bolo.HandleUpdate(screen)

		// draw the world
		world.Draw(art, renderer)

		// draw players
		tank.Update(delta)
		tank.Draw(art, renderer)

		// present everything
		renderer.Present()
		delta = time.Since(beginningOfFrame).Seconds() * 1000
	}
}

// Bolo -
type Bolo struct {
	world maps.WorldMap
	tanks []tank.Tank
	bullets []bullet.Bullet
}

// NewBolo -
func NewBolo() (bolo *Bolo) {
bolo.world = maps.NewWorldMap("./assets/test_map.txt", 1)

bolo.tanks = []tank.Tank{tank.NewTank()}

// TODO - User bullet.Manager here
bolo.bullets := make([]*bullet.Bullet, 50)

	return bolo
}

// Update -
func(b *Bolo) Update(screen *ebiten.Image) error {

}

// Draw -
func(b *Bolo) Draw(screen *ebiten.Image) error {

}

// Layout -
func(b *Bolo) Layout(width, height, int) (width, height int) {
	return 200, 100
}