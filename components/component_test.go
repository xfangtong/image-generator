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

var tcd = ComponentDefine{
	Type:  "image",
	Level: 1,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "400",
		Bottom: "500",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "0",
	BackgroundColor: "#ff0000",
	ComponentData: map[string]interface{}{
		"URL": "local://../test/dog.jpg",
	},
}

func TestDrawContainCenter(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_contain_center.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDrawCoverCenter(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Size = "cover"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_cover_center.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDrawAutoCenter(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Size = "auto"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_auto_center.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100Center(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Size = "100% 100%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_100_center.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100LeftTop(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "left top"
	cd.Size = "100% 100%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_100_lt.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100LeftBottom(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "left bottom"
	cd.Size = "100% 100%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_100_lb.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100RightTop(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "right top"
	cd.Size = "100% 100%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_100_rt.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100RightBottom(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "right bottom"
	cd.Size = "100% 100%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_100_rb.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100CenterRepeatX(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "center"
	cd.Size = "100% 100%"
	cd.Repeat = RepeatX

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_center_repeatx.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100CenterRepeatY(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "center"
	cd.Size = "100% 100%"
	cd.Repeat = RepeatY

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_center_repeaty.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw100RepeatXY(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "center"
	cd.Size = "100% 100%"
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
	f, _ := os.Create("../test/component_repeatxy.png")
	png.Encode(f, bg)

	f.Close()

}

func TestDraw50PaddingRepeatXY(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := draw2dimg.NewGraphicContext(bg)
	gc.SetFillColor(color.Black)
	gc.Clear()

	cd := tcd
	cd.Position = "center"
	cd.Size = "50% 50%"
	cd.Repeat = RepeatXY
	cd.Padding = "10"
	cd.Area.Left = "10%"
	cd.Area.Top = "20%"

	dc := &DrawContext{
		GraphicContext: draw2dimg.NewGraphicContext(bg),
		Image:          bg,
		Width:          400,
		Height:         500,
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	dc.DrawComponent(cd)
	f, _ := os.Create("../test/component_padding_repeatxy.png")
	png.Encode(f, bg)

	f.Close()

}
