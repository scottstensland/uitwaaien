package models

type Constants struct {
	SampleRate   int    `yaml:"sample_rate"`
	DefaultImage string `yaml:"default_image"`
	MediaDir     string `yaml:"media_dir"`
	ImageSize    struct {
		Width  int `yaml:"width"`
		Height int `yaml:"height"`
	} `yaml:"image_size"`
}
