package input

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/inpututil"
)

// Manager -
type Manager interface {
	Update()
	IsPressed(id int) bool
}

// StandardManager -
type StandardManager struct {
	controls map[int]int // {key: ID, value: millisecond duration}
}

// NewManager -
func NewManager() Manager {
	return &StandardManager{
		controls: make(map[int]int),
	}
}

var _ Manager = &StandardManager{}

// Update -
func (m *StandardManager) Update() {
	for id := range inpututil.() {
		m.controls[id] = 0
		println("anything")
	}

	for k := range m.controls {
		if inpututil.IsTouchJustReleased(k) {
			delete(m.controls, k)
		}
	}
}

// IsPressed -
func (m *StandardManager) IsPressed(id int) bool {
	println(fmt.Sprintf("controls: %+v", m.controls))
	_, ok := m.controls[id]
	if ok {
		println("found something")

	}
	return ok
}
