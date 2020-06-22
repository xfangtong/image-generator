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

var qc = ComponentDefine{
	Type:  "qrcode",
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
	Padding:         "10",
	BackgroundColor: "#000000",
	ComponentData: map[string]interface{}{
		"width":    150,
		"content":  "这是一个带LOGO的二维码，这是一个带LOGO的二维码，这是一个带LOGO的二维码，这是一个带LOGO的二维码，这是一个带LOGO的二维码，这是一个带LOGO的二维码，这是一个带LOGO的二维码",
		"logo":     "local://../images/avatar.jpg",
		"logoSize": 48,
		"isCircle": true,
		//"color":   "#ff00ff",
		//"color": "linear-gradient(100% 0, #ec008c 0%, #fc6767 100%)",
		"color": "radial-gradient(75 75 0 75 75 200, rgba(34,193,195,1) 0%, rgba(253,187,45,1) 100%)",
	},
}

func TestDrawQrcodeCenter(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := gg.NewContextForImage(bg)
	gc.SetColor(color.White)
	gc.Clear()

	cd := qc
	cd.Size = "100% 100%"
	cd.Repeat = RepeatNO

	dc := &DrawContext{
		GraphicContext: gc,
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_qrcode_center.png")
	png.Encode(f, gc.Image())

	f.Close()

}
