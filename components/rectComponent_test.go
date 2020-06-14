package components

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	_ "image/jpeg"
	_ "image/png"

	"github.com/llgcode/draw2d/draw2dimg"
)

var rect = ComponentDefine{
	Type:  "rect",
	Level: 1,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "0",
	BackgroundColor: "#00ffff",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ff0000",
		"strokeColor": "#00ff00",
		"lineWidth":   1,
		"Width":       150,
		"Height":      100,
		"Radius":      20,
	},
}

func TestDrawRectCenter(t *testing.T) {

	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.White)
	gc.Clear()

	cd := rect
	cd.Size = "100% 100%"
	//cd.Repeat = RepeatXY

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_rect_center.png")
	png.Encode(f, bg)

	f.Close()

}
