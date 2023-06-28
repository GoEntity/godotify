package godotify_test

import (
	"os"
	"testing"

	godotify "github.com/GoEntity/godotify/pkg"
)

func TestGoDotify(t *testing.T) {
	testCases := []struct {
		name        string
		inputFile   string
		outputFile  string
		intensity   float64
		expectedErr bool
	}{
		{
			name:       "Test valid PNG",
			inputFile:  "test_valid.png",
			outputFile: "test_valid_output.png",
			intensity:  0.5,
		},
		{
			name:        "Test invalid file",
			inputFile:   "test_invalid.txt",
			outputFile:  "test_invalid_output.png",
			intensity:   0.5,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := godotify.GoDotify(tc.inputFile, tc.outputFile, godotify.DottyConfig{Intensity: tc.intensity})

			if tc.expectedErr && err == nil {
				t.Errorf("error expected but got none")
			} else if !tc.expectedErr && err != nil {
				t.Errorf("didnt expect an error but got one: %v", err)
			}

			if !tc.expectedErr {
				if _, err := os.Stat(tc.outputFile); os.IsNotExist(err) {
					t.Errorf("didnt make output file")
				}
			}
		})
	}
}
