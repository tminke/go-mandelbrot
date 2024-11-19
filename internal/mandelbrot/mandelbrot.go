package mandelbrot

import (
	"image"
	"image/color"
	"math"
	"sync"

	"github.com/tminke/go-mandelbrot/internal/config"
)

var (
	BlackColor = color.RGBA{0, 0, 0, 0xff} // Colors are defined by Red, Green, Blue, Alpha uint8 values.
)

func NewMandelbrotRGBA(config config.Configuration) *image.RGBA {

	// Create The Empty Mandelbrot RGBA Image Based On Canvas Dimensions
	canvasUpperLeft := image.Point{X: 0, Y: 0}
	canvasLowerRight := image.Point{X: config.Image.CanvasWidth, Y: config.Image.CanvasHeight}
	img := image.NewRGBA(image.Rectangle{Min: canvasUpperLeft, Max: canvasLowerRight})

	// Create A WaitGroup
	var wg sync.WaitGroup

	// Calculate The Mandelbrot Value For Each Pixel
	columnSize := config.Image.CanvasWidth / config.Mandelbrot.Parallelism
	for column := 0; column < config.Mandelbrot.Parallelism; column++ {
		wg.Add(1)
		loopColumn := column
		go func() {
			defer wg.Done()
			startX := loopColumn * columnSize
			for x := startX; x < startX+columnSize; x++ {
				for y := 0; y < config.Image.CanvasHeight; y++ {
					pixelColor := calculateMandelbrotPixelColor(x, y, config)
					img.Set(x, y, pixelColor)
				}
			}
		}()
	}

	// Wait For Completion
	wg.Wait()

	// Return The Mandelbrot RGBA Image
	return img
}

// calculateMandelbrotPixelColor performs the Mandelbrot calculation and
// returns the resulting pixel Color of the specified coordinate.
func calculateMandelbrotPixelColor(x int, y int, config config.Configuration) color.Color {

	// Map The Canvas Coordinates Into Cartesian Coordinates
	origA := mapIntRangeToFloatRange(x, 0, config.Image.CanvasWidth-1, config.Image.XCoordinateMin, config.Image.XCoordinateMax)
	origB := mapIntRangeToFloatRange(y, 0, config.Image.CanvasHeight-1, config.Image.YCoordinateMin, config.Image.YCoordinateMax)

	// Initialize The Moving A & B Values To The Original A & B Values
	a := origA
	b := origB

	// Iterate The Mandelbrot Function To Determine Inclusivity
	iteration := 0
	for iteration = 0; iteration < config.Mandelbrot.MaxIterations; iteration++ {
		aa := a*a - b*b
		bb := 2 * a * b
		a = aa + origA
		b = bb + origB
		if math.Abs(a+b) > config.Mandelbrot.EscapeThreshold {
			break
		}
	}

	// Determine The Pixel Color (Grayscale) Based On Iteration Count
	pixelColor := calculateGrayscalePixelColor(iteration, config.Mandelbrot.MaxIterations)
	return pixelColor
}

// calculateGrayscalePixelColor returns a gray-scale color of the pixel based
// on the number of iterations it took to "escape" the boundary.
func calculateGrayscalePixelColor(iteration int, maxIterations int) color.RGBA {
	pixelColor := BlackColor
	if iteration < maxIterations {
		gradientValue := uint8(mapIntRangeToIntRange(iteration, 0, maxIterations, 0, 255))
		pixelColor = color.RGBA{R: gradientValue, G: gradientValue, B: gradientValue, A: 0xff}
	}
	return pixelColor
}

// mapIntRangeToIntRange returns an int representing the location in the
// range of newStart...newEnd equivalent to the original location of origValue
// in the range of origStart...origEnd.
func mapIntRangeToIntRange(origValue int, origStart int, origEnd int, newStart int, newEnd int) int {
	return int(math.Floor((float64(origValue)-float64(origStart))/(float64(origEnd)-float64(origStart))*(float64(newEnd)-float64(newStart)) + float64(newStart)))
}

// mapIntRangeToFloatRange returns a float representing the location in the
// range of newStart...newEnd equivalent to the original location of origValue
// in the range of origStart...origEnd.
func mapIntRangeToFloatRange(origValue int, origStart int, origEnd int, newStart float64, newEnd float64) float64 {
	return (float64(origValue)-float64(origStart))/(float64(origEnd)-float64(origStart))*(newEnd-newStart) + newStart
}
