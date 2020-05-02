package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/server/bullet"
	"github.com/ethanmil/bolo/server/maps"
	"github.com/ethanmil/bolo/server/tank"
	grpc "google.golang.org/grpc"
)

var delta float32

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	boloServer := NewBoloServer()

	log.Println("Serving...")

	srv := grpc.NewServer()
	guide.RegisterBoloServer(srv, boloServer)
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}

var _ guide.BoloServer = &BoloServer{}

// BoloServer -
type BoloServer struct {
	tanks    []tank.Tank
	bullets  []bullet.Bullet
	worldMap *maps.WorldMap
}

// NewBoloServer -
func NewBoloServer() *BoloServer {
	return &BoloServer{
		worldMap: maps.BuildMapFromFile(),
	}
}

// RegisterTank -
func (s *BoloServer) RegisterTank(ctx context.Context, t *guide.Tank) (*guide.Tank, error) {
	t.Id = int32(len(s.tanks) + 1)
	s.tanks = append(s.tanks, tank.NewTank(t.Id, s.worldMap))
	return t, nil
}

// GetWorldMap -
func (s *BoloServer) GetWorldMap(ctx context.Context, world *guide.WorldInput) (*guide.WorldMap, error) {
	log.Println("world map requested")
	return s.worldMap.GetStateMap(), nil
}

// ServerGameStateStream -
func (s *BoloServer) ServerGameStateStream(world *guide.WorldInput, stream guide.Bolo_ServerGameStateStreamServer) error {
	err := stream.Send(s.getStateFromGame())
	if err != nil {
		log.Fatalf("Failed to send game state data: %v", err)
		return err
	}
	return nil
}

// ClientInputStream -
func (s *BoloServer) ClientInputStream(stream guide.Bolo_ClientInputStreamServer) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&guide.Ok{})
		}
		if err != nil {
			return err
		}

		// update tanks & bullets based on client's info
		for i := range s.tanks {
			if s.tanks[i].ID == input.Id {
				s.tanks[i].HandleMovement(input)
			}
		}
	}
}

// Chat -
func (s *BoloServer) Chat(stream guide.Bolo_ChatServer) error {
	return nil
}

func (s *BoloServer) getStateFromGame() *guide.GameState {
	gs := &guide.GameState{}

	for i := range s.tanks {
		gs.Tanks = append(gs.Tanks, s.tanks[i].GetStateTank())
	}

	for i := range s.bullets {
		gs.Bullets = append(gs.Bullets, s.bullets[i].GetStateBullet())
	}

	gs.WorldMap = s.worldMap.GetStateMap()

	return gs
}
