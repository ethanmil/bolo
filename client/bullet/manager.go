package bullet

import (
	"image"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Manager -
type Manager struct {
	bullets []*Bullet
	art     *ebiten.Image
}

// NewManager -
func NewManager(art *ebiten.Image) *Manager {
	return &Manager{
		bullets: make([]*Bullet, 50),
		art:     art,
	}
}

// AddBullet -
func (m *Manager) AddBullet(tankID int32, position physics.Vector, angle physics.Angle) {
	m.bullets = append(m.bullets, &Bullet{
		TankID: tankID,
		Element: &animation.Element{
			Sprite:   m.art.SubImage(image.Rect(16, 144, 22, 152)).(*ebiten.Image),
			Position: position,
			Angle:    angle,
		},
	})
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
