package adapter

import "github.com/gdamore/tcell"

type Sound struct {
	screen *tcell.Screen
}

func NewSound(screen *tcell.Screen) *Sound {
	return &Sound{screen: screen}
}

func (s *Sound) Beep() {
	(*s.screen).Beep()
	return
}
