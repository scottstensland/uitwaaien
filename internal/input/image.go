package input

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type Image struct {
	FilePath string
}

func (i Image) Read() (image.Image, string, error) {

	fmt.Printf("Reading image from file: %s\n", i.FilePath)

	// read file
	imgfile, err := os.Open(i.FilePath)
	if err != nil {
		return nil, "", err
	}
	defer imgfile.Close()

	fileInfo, err := imgfile.Stat()
	if err != nil {
		return nil, "", err
	}

	fileSize := fileInfo.Size() //  stens TODO  impose a limit on file size
	fmt.Printf("File size: %d bytes\n", fileSize)

	// Decode will detect the format and return an image.Image
	img, format, err := image.Decode(imgfile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, "", err
	}
	return img, format, nil
}
