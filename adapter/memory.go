package adapter

import (
	"bufio"
	"os"

	chip8 "github.com/MarceloMPJR/go-chip-8"
)

func NewMemory(filepath *string) *chip8.StandardMemory {
	f, _ := os.Open(*filepath)
	buf := bufio.NewReader(f)

	rom := chip8.NewRom(buf)
	return chip8.NewStandardMemory(&chip8.ConfigMemory{Rom: rom})
}
