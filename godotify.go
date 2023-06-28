package godotify

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

type Config struct {
	Intensity float64
}

func GoDotify(inputFile string, outputFile string, config Config) error {
	
	//openfile
	infile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer infile.Close()

	//decode
	src, err := imaging.Decode(infile)
	if err != nil {
		return err
	}

	//to make it easier for humans (... reverse intensity val 0 - 1)
	intensity := 1.1 - config.Intensity
	if intensity < 0 {
		intensity = 0
	} else if intensity > 1 {
		intensity = 1
	}

	//make img smaller
	smaller := imaging.Resize(src, int(float64(src.Bounds().Dx())*0.1*intensity), 0, imaging.Lanczos)

	//make img bigger again
	godotified := imaging.Resize(smaller, src.Bounds().Dx(), src.Bounds().Dy(), imaging.NearestNeighbor)

	//openfile (output)
	outfile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	//encode and write output file
	ext := strings.ToLower(filepath.Ext(outputFile))
	if ext == ".png" {
		err = png.Encode(outfile, godotified)
	} else if ext == ".jpg" || ext == ".jpeg" {
		err = jpeg.Encode(outfile, godotified, nil)
	} else {
		return image.ErrFormat
	}
	if err != nil {
		return err
	}

	return nil
}