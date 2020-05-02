package bullet

import (
	"image"
	"io"
	"log"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Manager -
type Manager struct {
	bullets      []*Bullet
	art          *ebiten.Image
	bulletStream guide.Bolo_ShootBulletClient
}

// NewManager -
func NewManager(bulletStream guide.Bolo_ShootBulletClient, art *ebiten.Image) *Manager {
	return &Manager{
		art:          art,
		bulletStream: bulletStream,
	}
}

// AddBullet -
func (m *Manager) AddBullet(position physics.Vector, angle physics.Angle) {
	err := m.bulletStream.Send(&guide.Bullet{
		Id:    int32(len(m.bullets) + 1),
		X:     float32(position.X),
		Y:     float32(position.Y),
		Angle: float32(angle),
	})
	if err != nil && err != io.EOF {
		log.Fatalf("Send: %v", err)
	}
}

// SyncBulletsFromServer -
func (m *Manager) SyncBulletsFromServer(id int32, position physics.Vector, angle physics.Angle) {
	found := false
	for i := range m.bullets {
		if m.bullets[i].ID == id {
			found = true
			break
		}
	}
	if !found {
		m.bullets = append(m.bullets, &Bullet{
			ID: id,
			Element: &animation.Element{
				Sprite:   m.art.SubImage(image.Rect(16, 144, 22, 152)).(*ebiten.Image),
				Position: position,
				Angle:    angle,
			},
		})
	}
}

// Update -
func (m *Manager) Update(delta float64) {
	for i := range m.bullets {
		if m.bullets[i] != nil {
			m.bullets[i].Update(delta)
		}
	}
}

// Draw -
func (m *Manager) Draw(screen *ebiten.Image) {
	for i := range m.bullets {
		if m.bullets[i] != nil {
			m.bullets[i].Draw(screen)
		}
	}
}
