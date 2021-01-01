package decoder

import (
	"fmt"
	"github.com/opennota/screengen"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"image"
	"io"
	"os"
)

// VideoDecoder allows for the decoding of various types of videos.
type VideoDecoder struct{}

// Decode decodes videos into image.Image.
func (*VideoDecoder) Decode(extension string, reader io.Reader) (img image.Image, err error) {

	// FFMPEG wants a local file to open, so we need to save
	// the stream to a file.
	name := fmt.Sprintf("%s.%s", xid.New(), extension)
	path := fmt.Sprintf("/tmp/%s", name)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	defer func() {

		err := file.Close()
		if err != nil {
			log.Println("Unable to close correctly file ", name, err)
		}

		err = os.Remove(path)
		if err != nil {
			log.Println("Unable to remove correctly file ", name, err)
		}

	}()

	_, err = io.Copy(file, reader)
	if err != nil {
		return
	}

	generator, err := screengen.NewGenerator(path)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to create a screenshot generator for file %s", name))
		return
	}

	defer func() {
		err := generator.Close()
		if err != nil {
			log.Println("Unable to close generator for file ", name, err)
		}
	}()

	// Take the screenshot in the middle of the video
	frame, err := generator.Image(generator.Duration / 2)
	if err != nil {
		log.Println("Unable to extract frame from file ", name, err)
	}

	return frame, err

}
