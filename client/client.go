package main

import (
	"context"
	"log"

	"github.com/ethanmil/bolo/client/bologame"
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
	bolo := bologame.New(0, art, bullet.NewManager(art))

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

	// sync w/ server
	bolo.tankStreamOut, err = bolo.client.SendTankData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer bolo.tankStreamOut.CloseAndRecv()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBolo")

	err = ebiten.RunGame(bolo)
	if err != nil {
		logrus.Fatal(err)
	}
}
