package decoder

import (
	"image"
	"io"
)

// MediaDecoder is an interface for media decoding.
type MediaDecoder interface {
	Decode(extension string, reader io.Reader) (image.Image, error)
}
