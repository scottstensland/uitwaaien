package input

import (
	"image"
)

type Input interface {
	Read() (image.Image, string, error)
}
