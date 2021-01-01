package encoder

import (
	log "github.com/sirupsen/logrus"
	"image"
	"image/png"
	"os"
)

// SaveImageAsPNG saves an image.Image to the disk as filename in the PNG format.
func SaveImageAsPNG(filename string, img image.Image) error {

	outputFile, err := os.Create(filename)
	if err != nil {
		log.Println("SaveImageAsPNG: unable create file on disk ", err)
		return err
	}

	err = png.Encode(outputFile, img)
	if err != nil {
		log.Println("SaveImageAsPNG: unable to encode image ", err)
		return err
	}

	err = outputFile.Close()
	if err != nil {
		log.Println("SaveImageAsPNG: unable to close image file ", err)
	}

	return err

}
