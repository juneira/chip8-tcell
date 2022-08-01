package adapter

import (
	"bufio"
	"os"

	chip8 "github.com/MarceloMPJR/go-chip-8"
)

func NewMemory(filepath *string) *chip8.StandardMemory {
	f, _ := os.Open(*filepath)
	b := bufio.NewReader(f)

	return chip8.NewStandardMemory(&chip8.ConfigMemory{Rom: b})
}
