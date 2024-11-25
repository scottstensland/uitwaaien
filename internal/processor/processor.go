package processor

import (
	"image"

	"github.com/scottstensland/uitwaaien/internal/input"
)

type Processor struct {
	input input.Input
}

func NewProcessor(i input.Input) *Processor {
	return &Processor{input: i}
}

func (p *Processor) Process() (image.Image, string, error) {
	data, format, err := p.input.Read()
	if err != nil {
		return nil, "", err
	}
	return data, format, nil
}
