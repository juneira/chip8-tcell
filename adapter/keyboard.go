package adapter

import chip8 "github.com/MarceloMPJR/go-chip-8"

type KeyboardInput struct {
	key rune
}

// NewKeyboard returns a *adapter.KeyboardInput and *chip8.StandardKeyboard
func NewKeyboard() (*KeyboardInput, *chip8.StandardKeyboard) {
	ki := &KeyboardInput{}
	return ki, chip8.NewStandardKeyboard(&chip8.ConfigKeyboard{Input: ki})
}

func (ki *KeyboardInput) SetKey(key rune) {
	ki.key = key
}

func (ki *KeyboardInput) Read(p []byte) (n int, err error) {
	if ki.key == 0x0 {
		return 0, nil
	}

	p[0] = byte(ki.key)
	ki.key = 0x0
	return 1, nil
}
