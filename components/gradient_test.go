package components

import (
	"fmt"
	"image"
	"testing"

	"github.com/fogleman/gg"
)

func TestLinearGradientParse(t *testing.T) {
	g1 := "linear-gradient(100 100, rgba(2,0,36,1) 0%, rgba(9,9,121,1) 41%, rgba(8,40,141,1) 55%, rgba(0,212,255,1) 100%)"

	gg, _ := _parseLinearGradient(g1, 100.0, 100.0)

	fmt.Println(gg)
}

func TestRadialGradientParse(t *testing.T) {
	g1 := "radial-gradient(75 75 10 75 75 200, rgba(34,193,195,1) 0%, rgba(253,187,45,1) 100%)"

	g, _ := _parseRadialGradient(g1, 100.0, 100.0)

	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	gc := gg.NewContextForRGBA(img)
	gc.SetFillStyle(g)
	gc.DrawRectangle(0, 0, 200, 200)
	gc.Fill()

	gc.SavePNG("../test/radial_gradient.png")
}
