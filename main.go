package main

import (
	"github.com/ethanmil/bolo/bullet"
	"github.com/ethanmil/bolo/maps"
	"github.com/ethanmil/bolo/tank"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sirupsen/logrus"
)

var art *ebiten.Image
var err error

func init() {
	art, _, err = ebitenutil.NewImageFromFile("images/art.png", ebiten.FilterDefault)
	if err != nil {
		logrus.Fatal(err)
	}
}

var delta float64

func main() {
	bolo := NewBolo()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBolo")

	err := ebiten.RunGame(bolo)
	if err != nil {
		logrus.Fatal(err)
	}
}

// Bolo -
type Bolo struct {
	world   maps.WorldMap
	tanks   []tank.Tank
	bullets []bullet.Bullet
}

// NewBolo -
func NewBolo() *Bolo {
	return &Bolo{
		world:   maps.NewWorldMap("./assets/test_map.txt", 1, art),
		tanks:   []tank.Tank{tank.NewTank(art)},
		bullets: make([]bullet.Bullet, 50),
	}
}

// Update -
func (b *Bolo) Update(screen *ebiten.Image) error {
	b.world.Draw(screen)
	return nil

	// b.tanks[0].Update(1)
	// b.tanks[0].Draw(screen)
}

// Draw -
func (b *Bolo) Draw(screen *ebiten.Image) error {
	return nil
}

// Layout -
func (b *Bolo) Layout(width, height int) (int, int) {
	return 800, 600
}
