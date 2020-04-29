package main

import (
	"context"
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

func init() {
	art, _, err = ebitenutil.NewImageFromFile("client/images/art.png", ebiten.FilterDefault)
	if err != nil {
		logrus.Fatal(err)
	}
}

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
	id            int32
	world         *maps.WorldMap
	tanks         []tank.Tank
	bulletManager *bullet.Manager
	client        guide.BoloClient
}

// NewBolo -
func NewBolo() *Bolo {
	bulletManager := bullet.NewManager(art)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":9876", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := guide.NewBoloClient(conn)

	// build the world map using the tiles downloaded from the server
	serverWM, err := client.GetWorldMap(context.Background(), &guide.WorldInput{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get world map: %v", err)
	}
	world := maps.NewWorldMap(serverWM, art)

	// create a tank
	t := tank.NewTank(physics.Vector{X: 200, Y: 200}, art, world, bulletManager)

	return &Bolo{
		world:         world,
		tanks:         []tank.Tank{t},
		bulletManager: bulletManager,
		client:        client,
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
	stream, err := b.client.SendTankData(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	stream.Send(&guide.Tank{
		Id: b.id,
		X:  float32(b.tanks[b.id].Element.Position.X),
		Y:  float32(b.tanks[b.id].Element.Position.Y),
	})

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
