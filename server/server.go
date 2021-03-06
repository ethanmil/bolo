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
	log.Printf("TANKS: %v", s.tanks)
	return tank, nil
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
	if s.tanks != nil {
		for _, tank := range s.tanks {
			if tank != nil {
				if err := stream.Send(tank); err != nil {
					log.Fatalf("Failed to send tank data: %v", err)
					return err
				}
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

// GetBullets -
func (s *BoloServer) GetBullets(world *guide.WorldInput, stream guide.Bolo_GetBulletsServer) error {
	if s.bullets != nil {
		for _, bullet := range s.bullets {
			if bullet != nil {
				if err := stream.Send(bullet); err != nil {
					log.Fatalf("Failed to send bullet data: %v", err)
					return err
				}
			}
		}
	}
	return nil
}

// ShootBullet -
func (s *BoloServer) ShootBullet(stream guide.Bolo_ShootBulletServer) error {
	startTime := time.Now()
	for {
		bullet, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			println(int(endTime.Sub(startTime).Seconds()))
			return stream.SendAndClose(bullet)
		}
		if err != nil {
			log.Printf("error receiving: %v | %v", err, stream.Context())
			return err
		}

		if bullet != nil {
			found := false
			for i := range s.bullets {
				if s.bullets[i].Id == bullet.Id {
					found = true
					s.bullets[i] = bullet
					break
				}
			}
			if !found {
				s.bullets = append(s.bullets, bullet)
			}
		}
	}
}

// // RemoveBullet -
// func (s *BoloServer) RemoveBullet(stream guide.Bolo_RemoveBulletClient) error {
// 	startTime := time.Now()
// 	for {
// 		bullet, err := stream.Recv()
// 		if err == io.EOF {
// 			endTime := time.Now()
// 			println(int(endTime.Sub(startTime).Seconds()))
// 			return stream.SendAndClose(bullet)
// 		}
// 		if err != nil {
// 			log.Printf("error receiving: %v | %v", err, stream.Context())
// 			return err
// 		}

// 		if bullet != nil {
// 			for i := range s.bullets {
// 				if s.bullets[i].Id == bullet.Id {
// 					s.bullets[i] = bullet // remove bullet actually, lol
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// Chat -
func (s *BoloServer) Chat(stream guide.Bolo_ChatServer) error {
	return nil
}
