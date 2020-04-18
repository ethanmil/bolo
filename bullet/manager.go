package bullet

// Manager -
type Manager struct {
	bullets []*Bullet
}

// AddBullet -
func (m *Manager) AddBullet(bullet Bullet) {
	m.bullets = append(m.bullets, &bullet)
}
