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

// // ImageConfig defines the configuration used to render a single Mandelbrot image.
// type ImageConfig struct {
// 	EscapeThreshold float64 // The threshold value to determine if a point is in the Mandelbrot set
// 	MaxIterations   int     // Maximum number of iterations to calculate the Mandelbrot value before assuming it is in the set
// 	XCoordinateMin  float64 // The minimum x-coordinate value to calculate the Mandelbrot set
// 	XCoordinateMax  float64 // The maximum x-coordinate value to calculate the Mandelbrot set
// 	YCoordinateMin  float64 // The minimum y-coordinate value to calculate the Mandelbrot set
// 	YCoordinateMax  float64 // The maximum y-coordinate value to calculate the Mandelbrot set
// 	CanvasWidth     int     // Image width in pixels
// 	CanvasHeight    int     // Image height in pixels
// }

// // DefaultImageConfig is an ImageConfig with default values for a full-view QHD Mandelbrot image
// var DefaultImageConfig = ImageConfig{
// 	EscapeThreshold: 4.0,
// 	MaxIterations:   100,
// 	XCoordinateMin:  -2.25,
// 	XCoordinateMax:  0.75,
// 	YCoordinateMin:  -1.1,
// 	YCoordinateMax:  1.1,
// 	CanvasWidth:     PixelWidthQHD,
// 	CanvasHeight:    PixelHeightQHD,
// }

// // NewImageConfig creates a new ImageConfig with the specified values
// func NewImageConfig(escapeThreshold float64,
// 	maxIterations int,
// 	xCoordinateMin float64,
// 	xCoordinateMax float64,
// 	yCoordinateMin float64,
// 	yCoordinateMax float64,
// 	imageCanvasWidth int,
// 	imageCanvasHeight int) ImageConfig {

// 	return ImageConfig{
// 		EscapeThreshold: escapeThreshold,
// 		MaxIterations:   maxIterations,
// 		XCoordinateMin:  xCoordinateMin,
// 		XCoordinateMax:  xCoordinateMax,
// 		YCoordinateMin:  yCoordinateMin,
// 		YCoordinateMax:  yCoordinateMax,
// 		CanvasWidth:     imageCanvasWidth,
// 		CanvasHeight:    imageCanvasHeight,
// 	}
// }

// TODO - NEW !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

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
