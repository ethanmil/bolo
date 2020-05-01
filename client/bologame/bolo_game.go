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
	ID            int32
	Art           *ebiten.Image
	World         *maps.WorldMap
	Tanks         []tank.Tank
	BulletManager *bullet.Manager
	Client        guide.BoloClient
	TankStreamIn  guide.Bolo_GetTanksClient
	TankStreamOut guide.Bolo_SendTankDataClient
}

// New -
func New(art *ebiten.Image, bulletManager *bullet.Manager) *Bolo {
	return &Bolo{
		Art:           art,
		Tanks:         make([]tank.Tank, 8),
		BulletManager: bulletManager,
	}
}

// Update -
func (b *Bolo) Update(screen *ebiten.Image) error {
	b.World.Draw(screen)

	for i := range b.Tanks {
		if b.Tanks[i] != (tank.Tank{}) {
			if i == int(b.ID) {
				b.Tanks[i].Update(2)
			}
			b.Tanks[i].Draw(screen)
		}
	}

	b.BulletManager.Update(2)
	b.BulletManager.Draw(screen)

	// sync tank to server
	err = b.TankStreamOut.Send(&guide.Tank{
		Id: b.ID,
		X:  float32(b.Tanks[b.ID].Element.Position.X),
		Y:  float32(b.Tanks[b.ID].Element.Position.Y),
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
