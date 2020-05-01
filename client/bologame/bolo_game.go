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
		Tanks:         []tank.Tank{},
		BulletManager: bulletManager,
	}
}

// Update -
func (b *Bolo) Update(screen *ebiten.Image) error {
	b.World.Draw(screen)

	for i := range b.Tanks {
		if b.Tanks[i].ID == b.ID {
			b.Tanks[i].Update(2)
		}
		b.Tanks[i].Draw(screen)
		// log.Printf("TANKS: %v", b.Tanks)
	}

	b.BulletManager.Update(2)
	b.BulletManager.Draw(screen)

	// sync tank to server
	err = b.TankStreamOut.Send(&guide.Tank{
		Id: b.ID,
		X:  float32(b.Tanks[0].Element.Position.X),
		Y:  float32(b.Tanks[0].Element.Position.Y),
	})
	if err != nil && err != io.EOF {
		log.Fatalf("Send: %v", err)
	}

	// testing

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
