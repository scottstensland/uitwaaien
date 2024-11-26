// read in yaml config file

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/scottstensland/uitwaaien/internal/config"
	"github.com/scottstensland/uitwaaien/internal/input"
	"github.com/scottstensland/uitwaaien/internal/processor"
)

func run() {

	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading config: %v", err))
	}
	fmt.Println(config)

	input_image_path := filepath.Join(config.MediaDir, config.DefaultImage)
	// fmt.Println("input image path: ", input_image_path)

	image_input := input.Image{
		FilePath: input_image_path,
	}
	processor := processor.NewProcessor(image_input)
	data, format, err := processor.Process()
	if err != nil {
		panic(err)
	}
	println("format", format)
	// println(len(data))
	fmt.Printf("data: %T\n", data)
}

func main() {
	run()
}
