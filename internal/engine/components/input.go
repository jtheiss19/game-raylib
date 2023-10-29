package components

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Keybindings int32

const (
	MOVE_FORWARD   Keybindings = rl.KeyW
	MOVE_LEFT      Keybindings = rl.KeyA
	MOVE_BACKWARDS Keybindings = rl.KeyS
	MOVE_RIGHT     Keybindings = rl.KeyD
	MOVE_UP        Keybindings = rl.KeySpace
	MOVE_DOWN      Keybindings = rl.KeyLeftControl
	UNLOCK_CURSOR  Keybindings = rl.KeyTab
	LOCK_CURSOR    Keybindings = rl.KeyLeftShift
)

type KeyState bool

const (
	KEY_UP   KeyState = false
	KEY_DOWN KeyState = true
)

type InputComponent struct {
	*ecs.BaseComponent
	Keys        map[Keybindings]KeyState
	MouseLocked bool
}

func NewInputComponent() *InputComponent {
	newKeyMap := map[Keybindings]KeyState{
		MOVE_FORWARD:   KEY_UP,
		MOVE_LEFT:      KEY_UP,
		MOVE_BACKWARDS: KEY_UP,
		MOVE_RIGHT:     KEY_UP,
		MOVE_UP:        KEY_UP,
		MOVE_DOWN:      KEY_UP,
		UNLOCK_CURSOR:  KEY_UP,
		LOCK_CURSOR:    KEY_UP,
	}

	return &InputComponent{
		BaseComponent: &ecs.BaseComponent{},
		Keys:          newKeyMap,
		MouseLocked:   true,
	}
}

func (ic *InputComponent) CalculateInput() {
	for KeyBinding := range ic.Keys {
		ic.Keys[KeyBinding] = KeyState(rl.IsKeyDown(int32(KeyBinding)))
	}
}
