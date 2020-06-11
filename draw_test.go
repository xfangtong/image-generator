package imagegenerator

import (
	"image/png"
	"os"
	"testing"
)

func TestEmpty(t *testing.T) {
	tmp := ImageTemplate{
		Width:           300,
		Height:          300,
		BackgroundColor: "#ff0000",
	}

	img, _ := GenerateImage(tmp)

	f, _ := os.Create("./test/empty.png")
	defer f.Close()
	png.Encode(f, img)
}
