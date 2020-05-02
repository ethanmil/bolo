package bullet

// Manager -
type Manager struct {
	Bullets []Bullet
}

// NewManager -
func NewManager() *Manager {
	return &Manager{
		Bullets: []Bullet{},
	}
}

// Update -
func (b *Manager) Update() {
	for i := range b.Bullets {
		b.Bullets[i].Update()
	}
}
