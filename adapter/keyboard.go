package adapter

type KeyboardInput struct {
	Key rune
}

func (ki *KeyboardInput) Read(p []byte) (n int, err error) {
	if ki.Key == 0x0 {
		return 0, nil
	}

	p[0] = byte(ki.Key)
	return 1, nil
}
