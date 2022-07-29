package adapter

import (
	chip8 "github.com/MarceloMPJR/go-chip-8"
	"github.com/gdamore/tcell"
)

type DisplayOutput struct {
	screen *tcell.Screen
}

func NewDisplay(screen *tcell.Screen) *chip8.StandardDisplay {
	return chip8.NewStandardDisplay(
		&chip8.ConfigDisplay{Output: &DisplayOutput{screen: screen}},
	)
}

func (do *DisplayOutput) Write(p []byte) (n int, err error) {
	i, j := 0, 0

	style := tcell.StyleDefault.Foreground(tcell.ColorViolet)
	for _, r := range string(p) {
		if r == '\n' {
			i++
			j = 0
			continue
		}

		if string(r) == chip8.White {
			style = tcell.StyleDefault.Foreground(tcell.ColorViolet)
		} else {
			style = tcell.StyleDefault.Foreground(tcell.ColorRed)
		}

		(*do.screen).SetContent(j, i, '▓', nil, style)
		j++
	}

	(*do.screen).Show()
	return len(p), nil
}
