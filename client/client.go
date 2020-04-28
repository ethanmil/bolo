package main

import (
	"context"
	"log"

	"github.com/ethanmil/bolo/bullet"
	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/ethanmil/bolo/maps"
	"github.com/ethanmil/bolo/tank"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

const name = "ethan"

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
	client        guide.BoloClient
}

// NewBolo -
func NewBolo() *Bolo {
	// world := maps.NewWorldMap("./assets/test_map.txt", 1, art)
	bulletManager := bullet.NewManager(art)
	client := connectToServer()

	world, err := client.GetWorldMap(context.Background(), &guide.World{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get world map: %v", err)
	}

	return &Bolo{
		world:         world,
		tanks:         []tank.Tank{tank.NewTank(physics.Vector{X: 250, Y: 250}, art, world, bulletManager)},
		bulletManager: bulletManager,
		client:        client,
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

func connectToServer() guide.BoloClient {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to server!")

	return guide.NewBoloClient(conn)
}
