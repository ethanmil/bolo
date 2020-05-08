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
	// create game object
	bolo := bologame.New(art)

	// connect client
	conn := bolo.ConnectToServer()
	defer conn.Close()

	bolo.RegisterTank(ctx)

	// send our tank's data
	bolo.TankStreamOut, err = bolo.Client.SendTankData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer bolo.TankStreamOut.CloseAndRecv()

	// send our bullet data
	bolo.BulletStreamShoot, err = bolo.Client.ShootBullet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer bolo.BulletStreamShoot.CloseAndRecv()

	// send our bullet removal data
	bolo.BulletStreamRemove, err = bolo.Client.RemoveBullet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer bolo.BulletStreamRemove.CloseAndRecv()

	bolo.BulletManager = bullet.NewManager(bolo.BulletStreamShoot, bolo.BulletStreamRemove, art)

	// build the world map using the tiles downloaded from the server
	serverWM, err := bolo.Client.GetWorldMap(ctx, &guide.WorldInput{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get world map: %v", err)
	}
	bolo.World = maps.NewWorldMap(serverWM, art)

	// create our player's tank
	bolo.Tanks = append(bolo.Tanks, tank.NewTank(bolo.ID, physics.Vector{X: 200, Y: 200}, art, bolo.World, bolo.BulletManager))

	// always check for more tank data
	go func() {
		bolo.SyncTankData(ctx)
	}()

	// always check for more bullet data
	go func() {
		bolo.SyncBulletData(ctx)
	}()

	// run our ebiten game
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBolo")
	ebiten.SetRunnableOnUnfocused(true)

	err = ebiten.RunGame(bolo)
	if err != nil {
		logrus.Fatal(err)
	}
}
