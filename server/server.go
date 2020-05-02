package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/server/util"
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
	tanks    []*guide.Tank
	bullets  []*guide.Bullet
	worldMap *guide.WorldMap
}

// NewBoloServer -
func NewBoloServer() *BoloServer {
	return &BoloServer{
		worldMap: util.BuildMapFromFile(),
	}
}

// RegisterTank -
func (s *BoloServer) RegisterTank(ctx context.Context, tank *guide.Tank) (*guide.Tank, error) {
	tank.Id = int32(len(s.tanks) + 1)
	s.tanks = append(s.tanks, tank)
	return tank, nil
}

// GetWorldMap -
func (s *BoloServer) GetWorldMap(ctx context.Context, world *guide.WorldInput) (*guide.WorldMap, error) {
	log.Println("world map requested")
	return s.worldMap, nil
}

// ServerGameStateStream -
func (s *BoloServer) ServerGameStateStream(world *guide.WorldInput, stream guide.Bolo_ServerGameStateStreamServer) error {
	err := stream.Send(&guide.GameState{
		WorldMap: s.worldMap,
		Tanks:    s.tanks,
		Bullets:  s.bullets,
	})
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
			if s.tanks[i].Id == input.Id {
				s.tanks[i].HandleMovement(input)
			}
		}
	}
}

// Chat -
func (s *BoloServer) Chat(stream guide.Bolo_ChatServer) error {
	return nil
}
