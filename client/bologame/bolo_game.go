package bologame

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ethanmil/bolo/client/bullet"
	"github.com/ethanmil/bolo/client/maps"
	"github.com/ethanmil/bolo/client/tank"
	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
	"google.golang.org/grpc"
)

var err error

// Bolo -
type Bolo struct {
	ID                 int32
	Art                *ebiten.Image
	World              *maps.WorldMap
	Tanks              []tank.Tank
	BulletManager      *bullet.Manager
	Client             guide.BoloClient
	TankStreamIn       guide.Bolo_GetTanksClient
	TankStreamOut      guide.Bolo_SendTankDataClient
	BulletStreamIn     guide.Bolo_GetBulletsClient
	BulletStreamShoot  guide.Bolo_ShootBulletClient
	BulletStreamRemove guide.Bolo_RemoveBulletClient
}

// New -
func New(art *ebiten.Image) *Bolo {
	return &Bolo{
		Art:   art,
		Tanks: []tank.Tank{},
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
	}

	b.BulletManager.Update(2)
	b.BulletManager.Draw(screen)

	// sync tank to server
	err = b.TankStreamOut.Send(&guide.Tank{
		Id:    b.ID,
		X:     float32(b.Tanks[0].Element.Position.X),
		Y:     float32(b.Tanks[0].Element.Position.Y),
		Angle: float32(b.Tanks[0].Element.Angle),
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

// ConnectToServer -
func (b *Bolo) ConnectToServer() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":9876", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	b.Client = guide.NewBoloClient(conn)
	return conn
}

// RegisterTank -
func (b *Bolo) RegisterTank(ctx context.Context) {
	t, err := b.Client.RegisterTank(ctx, &guide.Tank{
		Name: "ethan",
	})
	if err != nil {
		log.Fatalf("Failed to register player: %v", err)
	}
	println(fmt.Sprintf("Tank: %+v", t))
	b.ID = t.Id
}

// SyncTankData -
func (b *Bolo) SyncTankData(ctx context.Context) {
	for {
		b.TankStreamIn, err = b.Client.GetTanks(ctx, &guide.WorldInput{Id: 1})
		if err != nil {
			log.Fatal(err)
		}
		defer b.TankStreamIn.CloseSend()
		for {
			t, err := b.TankStreamIn.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive tank stream in: %v", err)
			}

			if t != nil && t.Id != b.ID {
				found := false
				for i := range b.Tanks {
					if b.Tanks[i].ID == t.Id {
						found = true
						b.Tanks[i].Element.Position = physics.NewVector(float64(t.X), float64(t.Y))
						b.Tanks[i].Element.Angle = physics.NewAngle(float64(t.Angle))
						break
					}
				}
				if !found {
					b.Tanks = append(b.Tanks, tank.NewOtherTank(t.Id, physics.NewVector(float64(t.X), float64(t.Y)), b.Art))
				}
			}
		}
	}
}

// SyncBulletData -
func (b *Bolo) SyncBulletData(ctx context.Context) {
	for {
		b.BulletManager.Clear()
		b.BulletStreamIn, err = b.Client.GetBullets(ctx, &guide.WorldInput{Id: 1})
		if err != nil {
			log.Fatal(err)
		}
		defer b.BulletStreamIn.CloseSend()
		for {
			bul, err := b.BulletStreamIn.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive bullet stream in: %v", err)
			}

			if bul != nil {
				b.BulletManager.SyncBulletFromServer(
					bul.Id,
					physics.NewVector(float64(bul.X), float64(bul.Y)),
					physics.NewAngle(float64(bul.Angle)),
				)
			}
		}
	}
}
