package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"time"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/server/util"
	grpc "google.golang.org/grpc"
)

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
	players  []*guide.Player
	tanks    []*guide.Tank
	worldMap *guide.WorldMap
}

// NewBoloServer -
func NewBoloServer() *BoloServer {
	return &BoloServer{
		worldMap: util.BuildMapFromFile(),
	}
}

// RegisterPlayer -
func (s *BoloServer) RegisterPlayer(ctx context.Context, player *guide.Player) (*guide.Player, error) {
	player.Id = int32(len(s.players) + 1)
	s.players = append(s.players, player)
	log.Printf("PLAYER: %v", s.players)
	return player, nil
}

// GetPlayersOnline -
func (s *BoloServer) GetPlayersOnline(world *guide.WorldInput, stream guide.Bolo_GetPlayersOnlineServer) error {
	for _, player := range s.players {
		if err := stream.Send(player); err != nil {
			return err
		}
	}
	return nil
}

// GetWorldMap -
func (s *BoloServer) GetWorldMap(ctx context.Context, world *guide.WorldInput) (*guide.WorldMap, error) {
	log.Println("world map requested")
	return s.worldMap, nil
}

// GetWorldModifications -
func (s *BoloServer) GetWorldModifications(world *guide.WorldInput, stream guide.Bolo_GetWorldModificationsServer) error {
	return nil
}

// GetTanks -
func (s *BoloServer) GetTanks(world *guide.WorldInput, stream guide.Bolo_GetTanksServer) error {
	log.Println("get tanks called")
	if s.tanks != nil {
		for _, tank := range s.tanks {
			if tank != nil {
				if err := stream.Send(tank); err != nil {
					return err
				}
				log.Printf("Tanks from server: %v", tank)
			}
		}
	}
	return nil
}

// SendTankData -
func (s *BoloServer) SendTankData(stream guide.Bolo_SendTankDataServer) error {
	startTime := time.Now()
	for {
		tank, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			println(int(endTime.Sub(startTime).Seconds()))
			return stream.SendAndClose(tank)
		}
		if err != nil {
			log.Printf("error receiving: %v | %v", err, stream.Context())
			return err
		}

		if tank != nil {
			found := false
			for i := range s.tanks {
				if s.tanks[i].Id == tank.Id {
					found = true
					s.tanks[i] = tank
					break
				}
			}
			if !found {
				s.tanks = append(s.tanks, tank)
			}
		}
	}
}

// Chat -
func (s *BoloServer) Chat(stream guide.Bolo_ChatServer) error {
	return nil
}
