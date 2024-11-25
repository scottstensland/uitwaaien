package input

type Audio struct{}

func (a *Audio) Read() ([]byte, error) {
	return []byte("audio data ... stens TODO"), nil
}
