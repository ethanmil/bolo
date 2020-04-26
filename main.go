package main

import (
	"github.com/ethanmil/bolo/bullet"
	"github.com/ethanmil/bolo/lib/physics"
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
	world         *maps.WorldMap
	tanks         []tank.Tank
	bulletManager *bullet.Manager
}

// NewBolo -
func NewBolo() *Bolo {
	world := maps.NewWorldMap("./assets/test_map.txt", 1, art)
	bulletManager := bullet.NewManager(art)
	return &Bolo{
		world:         world,
		tanks:         []tank.Tank{tank.NewTank(physics.Vector{X: 100, Y: 100}, art, world, bulletManager)},
		bulletManager: bulletManager,
	}
}

// Update -
func (b *Bolo) Update(screen *ebiten.Image) error {
	// draw & update
	b.world.Draw(screen)

	b.tanks[0].Update(2)
	b.tanks[0].Draw(screen)

	b.bulletManager.Update(2)
	b.bulletManager.Draw(screen)
	return nil
}

// Draw -
func (b *Bolo) Draw(screen *ebiten.Image) error {
	return nil
}

// Layout -
func (b *Bolo) Layout(width, height int) (int, int) {
	return 800, 600
}
