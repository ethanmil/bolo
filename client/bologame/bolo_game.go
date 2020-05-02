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
	ID          int32
	Art         *ebiten.Image
	World       *maps.WorldMap
	Tanks       []tank.Tank
	Bullets     []bullet.Bullet
	Client      guide.BoloClient
	InputStream guide.Bolo_ClientInputStreamClient
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
		b.Tanks[i].Draw(screen)
	}

	for i := range b.Bullets {
		b.Bullets[i].Draw(screen)
	}

	b.SendInputToServer()

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
		X:    200,
		Y:    200,
	})
	if err != nil {
		log.Fatalf("Failed to register player: %v", err)
	}
	println(fmt.Sprintf("Tank: %+v", t))
	b.ID = t.Id
}

// ServerGameStateStream -
func (b *Bolo) ServerGameStateStream(ctx context.Context) {
	for {
		stream, err := b.Client.ServerGameStateStream(ctx, &guide.WorldInput{Id: 1})
		if err != nil {
			log.Fatal(err)
		}
		defer stream.CloseSend()
		for {
			state, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive game stream: %v", err)
			}

			b.setGameFromState(state)
		}
	}
}

// SendInputToServer -
func (b *Bolo) SendInputToServer() {
	input := &guide.UserInput{Id: b.ID}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		input.Left = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		input.Up = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		input.Right = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		input.Down = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		input.Shoot = true
	}

	if err := b.InputStream.Send(input); err != nil {
		log.Fatalf("Failed to send input data to the server: %v", err)
	}
}

func (b *Bolo) setGameFromState(state *guide.GameState) {
	for i := range state.Tanks {
		t := tank.NewTank(
			state.Tanks[i].Id,
			physics.Vector{
				X: state.Tanks[i].X,
				Y: state.Tanks[i].Y,
			},
			physics.NewAngle(state.Tanks[i].Angle),
			b.Art,
		)

		found := false
		for j := range b.Tanks {
			if b.Tanks[j].ID == state.Tanks[i].Id {
				found = true
				b.Tanks[j] = t
				break
			}
		}

		if !found {
			b.Tanks = append(b.Tanks, t)
		}
	}

	for i := range state.Bullets {
		bul := bullet.NewBullet(
			state.Bullets[i].Id,
			physics.NewVector(state.Bullets[i].X, state.Bullets[i].Y),
			physics.NewAngle(state.Bullets[i].Angle),
			b.Art,
		)

		found := false
		for j := range b.Bullets {
			if b.Bullets[j].ID == state.Bullets[i].Id {
				found = true
				b.Bullets[j] = bul
				break
			}
		}

		if !found {
			b.Bullets = append(b.Bullets, bul)
		}
	}

	b.World = maps.NewWorldMap(state.WorldMap, b.Art)
}
