package main

import (
	"flag"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/tminke/go-mandelbrot/internal/config"
	"github.com/tminke/go-mandelbrot/internal/mandelbrot"
)

func main() {

	// Parse Flags
	configFile := flag.String("config", "", "Name of the optional configuration file.")
	imageFile := flag.String("image", "mandelbrot.png", "Name to use for output image file.")
	flag.Parse()

	// Load The Optional Configuration Or Use Default
	configuration := config.DefaultConfiguration
	if configFile != nil && *configFile != "" {
		paresedConfiguration, err := config.ParseConfiguration(*configFile)
		if err != nil {
			log.Fatalf("Failed to parse configuration file '%s': err=%s", *configFile, err.Error())
		}
		configuration = *paresedConfiguration
	}

	// Calculate The Mandelbrot Image Tracking Calculation Time
	startTime := time.Now()
	img := mandelbrot.NewMandelbrotRGBA(configuration)
	calculationDuration := time.Since(startTime)

	// Create A File To Hold The Image
	file, err := os.Create(*imageFile)
	if err != nil {
		log.Fatalf("Failed to create PNG file: err=%s", err.Error())
	}

	// PNG Encode The Image To The File
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("Failed to encode PNG image: err=%s", err.Error())
	}

	// Log Success
	log.Printf("Created Mandelbrot image file '%s' with resolution '%d x %d' and parallelism of %d in %s",
		*imageFile, configuration.Image.CanvasWidth, configuration.Image.CanvasHeight, configuration.Mandelbrot.Parallelism, calculationDuration.String())
}
