package main

import (
	"bufio"
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

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

// NewBoloServer -
func NewBoloServer() *BoloServer {
	return &BoloServer{}
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
	return s.worldMap, nil
}

// GetWorldModifications -
func (s *BoloServer) GetWorldModifications(world *guide.WorldInput, stream guide.Bolo_GetWorldModificationsServer) error {
	return nil
}

// GetTanks -
func (s *BoloServer) GetTanks(world *guide.WorldInput, stream guide.Bolo_GetTanksServer) error {
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

func buildMapFromFile() *guide.WorldMap {
	wm := &guide.WorldMap{}
	file, err := os.Open("./assets/test_map.txt")
	if err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}

	wm.SizeW = int32(len(lines[0])/2 + 1)
	wm.SizeH = int32(len(lines))

	wm.Tiles = make([]*guide.WorldMap_Tile, wm.SizeH*wm.SizeW)
	for y := 0; y < int(wm.SizeH); y++ {
		for x, tileType := range strings.Split(lines[y], ",") {
			wm.Tiles[x*y] = &guide.WorldMap_Tile{
				Type: tileType,
				X:    int32(x * 32),
				Y:    int32(y * 32),
			}
		}
	}

	return wm
}
