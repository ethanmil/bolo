package bologame

import (
	"io"
	"log"

	"github.com/ethanmil/bolo/client/bullet"
	"github.com/ethanmil/bolo/client/maps"
	"github.com/ethanmil/bolo/client/tank"
	"github.com/ethanmil/bolo/guide"
	"github.com/hajimehoshi/ebiten"
)

var art *ebiten.Image
var delta float64
var err error

// Bolo -
type Bolo struct {
	id            int32
	art           *ebiten.Image
	world         *maps.WorldMap
	tanks         []tank.Tank
	bulletManager *bullet.Manager
	client        guide.BoloClient
	tankStreamOut guide.Bolo_SendTankDataClient
}

// New -
func New(id int32, art *ebiten.Image, bulletManager *bullet.Manager) *Bolo {
	return &Bolo{
		id:            id,
		art:           art,
		bulletManager: bulletManager,
	}
}

// Update -
func (b *Bolo) Update(screen *ebiten.Image) error {
	b.world.Draw(screen)

	for i := range b.tanks {
		b.tanks[i].Update(2)
		b.tanks[i].Draw(screen)
	}

	b.bulletManager.Update(2)
	b.bulletManager.Draw(screen)

	err = b.tankStreamOut.Send(&guide.Tank{
		Id: b.id,
		X:  float32(b.tanks[b.id].Element.Position.X),
		Y:  float32(b.tanks[b.id].Element.Position.Y),
	})
	if err != nil && err != io.EOF {
		log.Fatalf("Send: %v", err)
	}

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
