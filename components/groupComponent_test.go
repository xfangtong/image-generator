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

var c1 = ComponentDefine{
	Type:  "circle",
	Level: 2,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "10 20 0 0",
	BackgroundColor: "",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ff0000",
		"strokeColor": "#00ff00",
		"lineWidth":   10,
		"radius":      50,
	},
}

var c2 = ComponentDefine{
	Type:  "circle",
	Level: 1,
	Area: Rectangle{
		Left:   "50",
		Top:    "50",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "10 20 0 0",
	BackgroundColor: "",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ff0000",
		"strokeColor": "#00ffff",
		"lineWidth":   10,
		"radius":      50,
	},
}

var c3 = ComponentDefine{
	Type:  "circle",
	Level: 1,
	Area: Rectangle{
		Left:   "100",
		Top:    "100",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "10 20 0 0",
	BackgroundColor: "",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ffff00",
		"strokeColor": "#00ff00",
		"lineWidth":   10,
		"radius":      50,
	},
}

var c4 = ComponentDefine{
	Type:  "circle",
	Level: 3,
	Area: Rectangle{
		Left:   "150",
		Top:    "150",
		Right:  "auto",
		Bottom: "auto",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "10 20 0 0",
	BackgroundColor: "",
	ComponentData: map[string]interface{}{
		"fillColor":   "#ff4400",
		"strokeColor": "#00ff00",
		"lineWidth":   10,
		"radius":      50,
	},
}

var cl = []ComponentDefine{
	c1, c2, c3, c4,
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
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.White)
	gc.Clear()

	cd := group
	cd.Size = "50% 50%"
	cd.Repeat = RepeatXY

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_group.png")
	png.Encode(f, bg)

	f.Close()

}
