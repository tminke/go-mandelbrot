package mandelbrot

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tminke/go-mandelbrot/internal/config"
)

func TestNewMandelbrotRGBA(t *testing.T) {
	img := NewMandelbrotRGBA(config.DefaultConfiguration)
	assert.NotNil(t, img)
	assert.Equal(t, image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: config.DefaultConfiguration.Image.CanvasWidth, Y: config.DefaultConfiguration.Image.CanvasHeight}}, mandelbrotImage.Bounds())
	assert.Equal(t, BlackColor, img.At(1920, 720))                  // Cartesian 0,0 = Black
	assert.Equal(t, color.RGBA{71, 71, 71, 255}, img.At(1270, 650)) // Cartesian Point In Seahorse Valley = Gray
}

func TestCalculateMandelbrotPixelColor(t *testing.T) {
	assert.Equal(t, BlackColor, calculateMandelbrotPixelColor(1920, 720, config.DefaultConfiguration))                  // Cartesian 0,0 = Black
	assert.Equal(t, color.RGBA{71, 71, 71, 255}, calculateMandelbrotPixelColor(1270, 650, config.DefaultConfiguration)) // Cartesian Point In Seahorse Valley = Gray
}

func TestCalculateGrayscalePixelColor(t *testing.T) {
	assert.Equal(t, BlackColor, calculateGrayscalePixelColor(4, 2))
	assert.Equal(t, color.RGBA{54, 54, 54, 0xff}, calculateGrayscalePixelColor(20, 100))
}

func TestMapIntRangeToIntRange(t *testing.T) {
	assert.Equal(t, 20, mapIntRangeToIntRange(2, 0, 10, 0, 100))
	assert.Equal(t, 50, mapIntRangeToIntRange(5, 0, 10, 0, 100))
	assert.Equal(t, 160, mapIntRangeToIntRange(8, 0, 10, 0, 200))
}

func TestMapIntRangeToFloatRange(t *testing.T) {
	assert.Equal(t, 20.0, mapIntRangeToFloatRange(2, 0, 10, 0.0, 100.0))
	assert.Equal(t, 50.0, mapIntRangeToFloatRange(5, 0, 10, 0.0, 100.0))
	assert.Equal(t, 160.0, mapIntRangeToFloatRange(8, 0, 10, 0.0, 200.0))
}
