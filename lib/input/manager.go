package input

import (
	"github.com/hajimehoshi/ebiten"
)

// Manager -
type Manager interface {
	Update()
	IsKeyPressed(id ebiten.Key) bool
}

// StandardManager -
type StandardManager struct {
	keyboard map[ebiten.Key]int // {key: ID, value: millisecond duration}
	gamepad  map[int]struct{}
}

// NewManager -
func NewManager() Manager {
	return &StandardManager{
		keyboard: make(map[ebiten.Key]int),
		gamepad:  make(map[int]struct{}),
	}
}

var _ Manager = &StandardManager{}

// Update -
func (m *StandardManager) Update() {
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			m.keyboard[k] = 0
		}
	}
}

// IsKeyPressed -
func (m *StandardManager) IsKeyPressed(id ebiten.Key) bool {
	// println(fmt.Sprintf("controls: %+v", m.controls))
	_, ok := m.keyboard[id]
	if ok {
		println("found something")

	}
	return ok
}

// TODO: Gamepad logic to handle later
// func (m *StandardManager) updateGamepad() {
// 	for _, id := range inpututil.JustConnectedGamepadIDs() {
// 		log.Printf("gamepad connected: id: %d", id)
// 		m.gamepad[id] = struct{}{}
// 	}
// 	for id := range m.gamepad {
// 		if inpututil.IsGamepadJustDisconnected(id) {
// 			log.Printf("gamepad disconnected: id: %d", id)
// 			delete(m.gamepad, id)
// 		}
// 	}
// }
