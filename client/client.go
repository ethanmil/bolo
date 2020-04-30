package main

import (
	"context"
	"io"
	"log"

	"github.com/ethanmil/bolo/client/bullet"
	"github.com/ethanmil/bolo/client/maps"
	"github.com/ethanmil/bolo/client/tank"
	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const name = "ethan"

var art *ebiten.Image
var delta float64
var err error

var ctx context.Context

func init() {
	art, _, err = ebitenutil.NewImageFromFile("client/images/art.png", ebiten.FilterDefault)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx = context.Background()
}

func main() {
	bolo := NewBolo()

	// create client
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":9876", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	bolo.client = guide.NewBoloClient(conn)
	defer conn.Close()

	// build the world map using the tiles downloaded from the server
	serverWM, err := bolo.client.GetWorldMap(ctx, &guide.WorldInput{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get world map: %v", err)
	}
	bolo.world = maps.NewWorldMap(serverWM, art)

	// create a tank
	bolo.tanks = []tank.Tank{tank.NewTank(physics.Vector{X: 200, Y: 200}, art, bolo.world, bolo.bulletManager)}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBolo")

	err = ebiten.RunGame(bolo)
	if err != nil {
		logrus.Fatal(err)
	}
}

// Bolo -
type Bolo struct {
	id            int32
	world         *maps.WorldMap
	tanks         []tank.Tank
	bulletManager *bullet.Manager
	client        guide.BoloClient
}

// NewBolo -
func NewBolo() *Bolo {
	bulletManager := bullet.NewManager(art)

	return &Bolo{
		id:            0,
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

	// sync w/ server
	tankStreamOut, err := b.client.SendTankData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = tankStreamOut.Send(&guide.Tank{
		Id: b.id,
		X:  float32(b.tanks[b.id].Element.Position.X),
		Y:  float32(b.tanks[b.id].Element.Position.Y),
	})
	if err != nil && err != io.EOF {
		log.Fatalf("Send: %v", err)
	}

	ta, err := tankStreamOut.CloseAndRecv()
	if err != nil && err != io.EOF {
		log.Fatalf("Close and Recv: %v | %v", err, ta)
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
