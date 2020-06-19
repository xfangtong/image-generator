package components

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	_ "image/jpeg"
	_ "image/png"

	"github.com/fogleman/gg"
)

var avatar = ComponentDefine{
	Type:  "avatar",
	Level: 1,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "100%",
		Bottom: "100%",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "0",
	BackgroundColor: "#000000",
	ComponentData: map[string]interface{}{
		"fillColor":   "transparent",
		"strokeColor": "transparent",
		"lineWidth":   10,
		"width":       100,
		"size":        "auto auto",
		"url":         "local://../images/avatar.jpg",
	},
}

func TestDrawAvatarCenter(t *testing.T) {

	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := gg.NewContextForImage(bg)
	gc.SetColor(color.White)
	gc.Clear()

	cd := avatar
	cd.Size = "100% 100%"
	cd.Repeat = RepeatXY

	dc := &DrawContext{
		GraphicContext: gc,
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_avatar_center.png")
	png.Encode(f, gc.Image())

	f.Close()

}
