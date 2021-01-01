package decoder

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/webp"
	"golang.org/x/xerrors"
)

// ImageDecoder allows for the decoding of various types of images.
type ImageDecoder struct{}

// Decode decodes images into image.Image.
func (*ImageDecoder) Decode(extension string, reader io.Reader) (img image.Image, err error) {

	switch extension {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(reader)
	case "png":
		img, err = png.Decode(reader)
	case "webp":
		img, err = webp.Decode(reader)
	default:
		err = xerrors.Errorf("ImageDecoder: format %s is not supported", extension)
	}

	return

}
