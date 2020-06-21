package components

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/fogleman/gg"
	"github.com/go-playground/assert/v2"
)

var text = ComponentDefine{
	Type:  "text",
	Level: 1,
	Area: Rectangle{
		Left:   "0",
		Top:    "0",
		Right:  "200",
		Bottom: "100%",
	},
	Position:        "center",
	Size:            "contain",
	Repeat:          RepeatNO,
	Padding:         "0",
	BackgroundColor: "#000000",
	ComponentData: map[string]interface{}{
		//"fillColor": "#ff0000",
		"fillColor": "linear-gradient(100% 0, #FF4E50 0%, #F9D423 100%)",
		//"strokeColor": "#00ff00",
		//"lineWidth":   10,
		"text":       "HELLO，\r\n你好，\r你在做什么？this is a long text begintextbegintextbegintextbegin",
		"fontPath":   "../fonts/zkgdh.ttf",
		"fontSize":   24,
		"lineHeight": 32,
	},
}

func TestSplitText(t *testing.T) {
	s := "abc"
	ms := splitText(s)
	assert.Equal(t, string(ms[0]), s)

	s = "abc word"
	ms = splitText(s)
	assert.Equal(t, len(ms), 3)
	assert.Equal(t, string(ms[0]), "abc")
	assert.Equal(t, string(ms[1]), " ")
	assert.Equal(t, string(ms[2]), "word")

	s = "a国 b\r\n家、🏠\t"
	ms = splitText(s)
	assert.Equal(t, len(ms), 9)
	assert.Equal(t, string(ms[0]), "a")
	assert.Equal(t, string(ms[1]), "国")
	assert.Equal(t, string(ms[2]), " ")
	assert.Equal(t, string(ms[3]), "b")
	assert.Equal(t, string(ms[4]), "\n")
	assert.Equal(t, string(ms[5]), "家")
	assert.Equal(t, string(ms[6]), "、")
	assert.Equal(t, string(ms[7]), "🏠")
	assert.Equal(t, string(ms[8]), "\t")
}

func TestDrawText(t *testing.T) {
	bg := image.NewRGBA(image.Rect(0, 0, 400, 500))
	gc := gg.NewContextForRGBA(bg)
	gc.SetColor(color.Black)
	gc.Clear()

	cd := text
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
	f, _ := os.Create("../test/component_text_center.png")
	png.Encode(f, gc.Image())

	f.Close()

}
