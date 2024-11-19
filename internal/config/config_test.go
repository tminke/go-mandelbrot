package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseConfiguration(t *testing.T) {

	// Test Data
	const configFile = "test-config.yaml"

	// Define The Test Cases
	testCases := map[string]struct {
		fileContents        []byte
		expectConfiguration *Configuration
		expectErr           bool
	}{
		"Missing File Error": {
			fileContents:        nil,
			expectConfiguration: nil,
			expectErr:           true,
		},
		"Non YAML Content Error": {
			fileContents:        []byte("this is not yaml"),
			expectConfiguration: nil,
			expectErr:           true,
		},
		"Success": {
			fileContents: []byte(`
mandelbrot:
  escapeThreshold: 4.0
  maxIterations: 20
  parallelism: 4
image:
  xCoordinateMin: -2.25
  xCoordinateMax: 0.75
  yCoordinateMin: -1.2
  yCoordinateMax: 1.2
  canvasWidth: 1920
  canvasHeight: 1080`),
			expectConfiguration: &Configuration{
				Mandelbrot: MandelbrotConfiguration{
					EscapeThreshold: 4.0,
					MaxIterations:   20,
					Parallelism:     4,
				},
				Image: ImageConfiguration{
					XCoordinateMin: -2.25,
					XCoordinateMax: 0.75,
					YCoordinateMin: -1.2,
					YCoordinateMax: 1.2,
					CanvasWidth:    1920,
					CanvasHeight:   1080,
				},
			},
			expectErr: false,
		},
	}

	// Execute The Test Cases
	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, func(t *testing.T) {

			// Create The Config YAML File To Parse
			if len(testCase.fileContents) > 0 {
				err := os.WriteFile(configFile, testCase.fileContents, 0644)
				require.NoError(t, err)
				defer func() { require.Nil(t, os.Remove(configFile)) }()
			}

			// Perform The TestCase
			configuration, err := ParseConfiguration(configFile)

			// Verify The Results
			assert.Equal(t, testCase.expectConfiguration, configuration)
			assert.Equal(t, testCase.expectErr, err != nil)
		})
	}
}
