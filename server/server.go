package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ethanmil/bolo/guide"
	grpc "google.golang.org/grpc"
)

var _ guide.BoloServer = &BoloServer{}

// BoloServer -
type BoloServer struct {
	players  []*guide.Player
	tanks    []*guide.Tank
	worldMap *guide.WorldMap
}

// GetPlayersOnline -
func (s *BoloServer) GetPlayersOnline(world *guide.World, stream guide.Bolo_GetPlayersOnlineServer) error {
	for _, player := range s.players {
		if err := stream.Send(player); err != nil {
			return err
		}
	}

	return nil
}

// GetWorldMap -
func (s *BoloServer) GetWorldMap(ctx context.Context, world *guide.World) (*guide.WorldMap, error) {
	return s.worldMap, nil
}

// GetWorldModifications -
func (s *BoloServer) GetWorldModifications(world *guide.World, stream guide.Bolo_GetWorldModificationsServer) error {
	return nil
}

// GetTanks -
func (s *BoloServer) GetTanks(world *guide.World, stream guide.Bolo_GetTanksServer) error {
	for _, tank := range s.tanks {
		if err := stream.Send(tank); err != nil {
			return err
		}
	}
	return nil
}

// SendTankData -
func (s *BoloServer) SendTankData(stream guide.Bolo_SendTankDataServer) error {
	return nil
}

// Chat -
func (s *BoloServer) Chat(stream guide.Bolo_ChatServer) error {
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving...")

	grpcServer := grpc.NewServer()
	guide.RegisterBoloServer(grpcServer, &BoloServer{})
	grpcServer.Serve(lis)
}
