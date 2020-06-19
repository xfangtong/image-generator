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

var a1 = ComponentDefine{
	Type:  "avatar",
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

var t1 = ComponentDefine{
	Type:  "text",
	Level: 1,
	Area: Rectangle{
		Left:   "auto",
		Top:    "auto",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "0",
	BackgroundColor: "#000000",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ff0000",
		"strokeColor": "#00ff00",
		"text":        "张三000err",
		"fontPath":    "../fonts/zkgdh.ttf",
		"fontSize":    24,
		"lineHeight":  100,
	},
}

var cl = []ComponentDefine{
	a1, t1,
}

var group = ComponentDefine{
	Type:  "group",
	Level: 1,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "100%",
		Bottom: "100%",
	},
	Position:        "center",
	Size:            "100% 100%",
	Repeat:          RepeatNO,
	Padding:         "10",
	BackgroundColor: "#333",
	ComponentData: map[string]interface{}{
		"components": cl,
	},
}

func TestDrawGroup(t *testing.T) {

	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := gg.NewContextForImage(bg)
	gc.SetColor(color.White)
	gc.Clear()

	cd := group
	cd.Size = "100% 100%"
	//cd.Repeat = RepeatXY

	dc := &DrawContext{
		GraphicContext: gc,
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_group.png")
	png.Encode(f, gc.Image())

	f.Close()

}
