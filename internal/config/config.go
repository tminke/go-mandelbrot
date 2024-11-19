package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	PixelWidth8K  = 7680
	PixelHeight8K = 4320

	PixelWidth4K  = 3840
	PixelHeight4K = 2160

	PixelWidthQHD  = 2560
	PixelHeightQHD = 1440

	PixelWidthFHD  = 1920
	PixelHeightFHD = 1080
)

type Configuration struct {
	Mandelbrot MandelbrotConfiguration `yaml:"mandelbrot"`
	Image      ImageConfiguration      `yaml:"image"`
}

type MandelbrotConfiguration struct {
	EscapeThreshold float64 `yaml:"escapeThreshold"`
	MaxIterations   int     `yaml:"maxIterations"`
	Parallelism     int     `yaml:"parallelism"`
}

type ImageConfiguration struct {
	XCoordinateMin float64 `yaml:"xCoordinateMin"`
	XCoordinateMax float64 `yaml:"xCoordinateMax"`
	YCoordinateMin float64 `yaml:"yCoordinateMin"`
	YCoordinateMax float64 `yaml:"yCoordinateMax"`
	CanvasWidth    int     `yaml:"canvasWidth"`
	CanvasHeight   int     `yaml:"canvasHeight"`
}

// DefaultConfiguration is a Configuration with default values for a full-view QHD Mandelbrot image
var DefaultConfiguration = Configuration{
	Mandelbrot: MandelbrotConfiguration{
		EscapeThreshold: 4.0,
		MaxIterations:   100,
		Parallelism:     16,
	},
	Image: ImageConfiguration{
		XCoordinateMin: -2.25,
		XCoordinateMax: 0.75,
		YCoordinateMin: -1.1,
		YCoordinateMax: 1.1,
		CanvasWidth:    PixelWidthQHD,
		CanvasHeight:   PixelHeightQHD,
	},
}

// ParseConfiguration returns a Configuration based on the contents of the
// specified YAML config file, or nil and an error if any occurs.
func ParseConfiguration(configFile string) (*Configuration, error) {

	// Open The Specified ConfigFile (ReadOnly)
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode The ConfigFile Into A Configuration Struct
	newConfig := &Configuration{}
	yamlDecoder := yaml.NewDecoder(file)
	err = yamlDecoder.Decode(newConfig)
	if err != nil {
		return nil, err
	}

	// Succes - Return The Decoded Configuration
	return newConfig, nil
}
