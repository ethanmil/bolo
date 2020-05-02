package main

import (
	"context"
	"log"

	"github.com/ethanmil/bolo/client/bologame"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sirupsen/logrus"
)

const name = "ethan"

var art *ebiten.Image
var delta float32
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

	bolo.InputStream, err = bolo.Client.ClientInputStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to start client input stream %v", err)
	}
	defer bolo.InputStream.CloseAndRecv()

	// always update game state
	go func() {
		bolo.ServerGameStateStream(ctx)
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
