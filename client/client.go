package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// create game object
	bolo := bologame.New(art, bullet.NewManager(art))

	// connect client
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":9876", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	bolo.Client = guide.NewBoloClient(conn)
	defer conn.Close()

	// get players
	players, err := bolo.Client.GetPlayersOnline(ctx, &guide.WorldInput{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get players: %v", err)
	}
	println(players)

	// register ourself
	player, err := bolo.Client.RegisterPlayer(ctx, &guide.Player{
		Name: "ethan",
	})
	if err != nil {
		log.Fatalf("Failed to register player: %v", err)
	}
	println(fmt.Sprintf("PLAYER: %+v", player))
	bolo.ID = player.Id

	// build the world map using the tiles downloaded from the server
	serverWM, err := bolo.Client.GetWorldMap(ctx, &guide.WorldInput{Id: 1})
	if err != nil {
		log.Fatalf("Failed to get world map: %v", err)
	}
	bolo.World = maps.NewWorldMap(serverWM, art)

	// create our player's tank
	bolo.Tanks = append(bolo.Tanks, tank.NewTank(player.Id, physics.Vector{X: 200, Y: 200}, art, bolo.World, bolo.BulletManager))

	// set up our server streams

	// send our tank's data
	bolo.TankStreamOut, err = bolo.Client.SendTankData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer bolo.TankStreamOut.CloseAndRecv()

	// always check for more tank data
	go func() {
		for {
			bolo.TankStreamIn, err = bolo.Client.GetTanks(context.Background(), &guide.WorldInput{Id: 1})
			if err != nil {
				log.Fatal(err)
			}
			defer bolo.TankStreamIn.CloseSend()
			t, err := bolo.TankStreamIn.Recv()
			if err != nil && err != io.EOF {
				log.Fatalf("Failed to receive tank stream in: %v", err)
			}

			if t != nil && t.Id != bolo.ID {
				found := false
				for i := range bolo.Tanks {
					if bolo.Tanks[i].ID == t.Id {
						found = true
						bolo.Tanks[i].Element.Position = physics.NewVector(float64(t.X), float64(t.Y))
						bolo.Tanks[i].Element.Angle = physics.NewAngle(float64(t.Angle))
						break
					}
				}
				if !found {
					bolo.Tanks = append(bolo.Tanks, tank.NewOtherTank(t.Id, physics.NewVector(float64(t.X), float64(t.Y)), bolo.Art))
				}
			}
		}
	}()

	// run our ebiten game
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBolo")
	ebiten.SetRunnableOnUnfocused(true)

	time.Sleep(500)

	err = ebiten.RunGame(bolo)
	if err != nil {
		logrus.Fatal(err)
	}
}
