package adapter_test

import (
	"testing"

	"github.com/MarceloMPJR/chip8-tcell/adapter"
)

func TestKeyboardInput_Read(t *testing.T) {
	var result [1]byte
	t.Run("when key is NULL", func(t *testing.T) {
		k := &adapter.KeyboardInput{}
		n, err := k.Read(result[:])
		if err != nil {
			t.Fatalf("error unexpected: %s", err.Error())
		}

		if n != 0 {
			t.Fatalf("result: %d, expected: %d", n, 0)
		}
	})

	t.Run("when key is 'a'", func(t *testing.T) {
		expected := 'a'
		k := &adapter.KeyboardInput{Key: expected}
		n, err := k.Read(result[:])
		if err != nil {
			t.Fatalf("error unexpected: %s", err.Error())
		}

		if n != 1 {
			t.Fatalf("result: %d, expected: %d", n, 1)
		}

		if rune(result[0]) != expected {
			t.Errorf("result: %c, expected: %c", rune(result[0]), expected)
		}
	})
}
