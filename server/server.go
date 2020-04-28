package server

import context "context"

var _ BoloServerServer = &BoloServer{}

// BoloServer -
type BoloServer struct {
	players []*Player
}

// GetPlayersOnline -
func (s *BoloServer) GetPlayersOnline(world *World, stream BoloServer_GetPlayersOnlineServer) error {
	for _, player := range s.players {
		if err := stream.Send(player); err != nil {
			return err
		}
	}

	return nil
}

// GetWorldMap -
func (s *BoloServer) GetWorldMap(context.Context, *World) (*WorldMap, error) {
	return nil, nil
}

// GetWorldModifications -
func (s *BoloServer) GetWorldModifications(*World, BoloServer_GetWorldModificationsServer) error {
	return nil
}

// GetTanks -
func (s *BoloServer) GetTanks(*World, BoloServer_GetTanksServer) error {
	return nil
}

// SendTankData -
func (s *BoloServer) SendTankData(BoloServer_SendTankDataServer) error {
	return nil
}

// Chat -
func (s *BoloServer) Chat(BoloServer_ChatServer) error {
	return nil
}
