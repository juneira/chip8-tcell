package main

import (
	"flag"
	"log"
	"os"

	"github.com/MarceloMPJR/chip8-tcell/adapter"
	chip8 "github.com/MarceloMPJR/go-chip-8"
	"github.com/gdamore/tcell"
)

func drawText(s tcell.Screen, x1, y1 int, style tcell.Style, text string) {
	for i, r := range []rune(text) {
		s.SetContent(x1+i, y1, r, nil, style)
	}
}

func main() {
	filepath := flag.String("file", "", "file path of CHIP-8 program")
	flag.Parse()

	// Initialize screen
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(style)
	screen.Clear()

	keyboardInput, keyboard := adapter.NewKeyboard()
	go keyboardEventLoop(keyboardInput, &screen)

	display := adapter.NewDisplay(&screen)

	memory := adapter.NewMemory(filepath)

	sound := adapter.NewSound(&screen)

	cpu(&screen, display, keyboard, memory, sound)
}

func cpu(screen *tcell.Screen, display chip8.Display, keyboard chip8.Keyboard, memory chip8.Memory, sound chip8.Sound) {
	cpu := chip8.NewCpu(&chip8.ConfigCpu{
		Display:  display,
		Keyboard: keyboard,
		Sound:    sound,
		Memory:   memory,
		PC:       0x200,
	})

	cpu.Start()
	for {
		pc := cpu.NextInstruction()
		instr := memory.LoadInstruction(pc)

		k := [0x1FF]byte{}
		memory.Load(k[:], 0)

		cpu.Process(instr)
	}
}

func keyboardEventLoop(ki *adapter.KeyboardInput, s *tcell.Screen) {
	for {
		ev := (*s).PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				(*s).Fini()
				os.Exit(0)
				return
			}
			ki.SetKey(ev.Rune())
		}
	}
}
