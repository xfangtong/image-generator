package components

import (
	"image"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

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
	assert.Equal(t, string(ms[8]), "    ")
}

func TestMeasureText(t *testing.T) {
	s := "a国 b\r\n家、🏠\t"
	ms := splitText(s)

	bg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	gc := draw2dimg.NewGraphicContext(bg)

	draw2d.SetFontFolder("../fonts")

	gc.SetFontSize(12)

	measureText(gc, ms, 100)
}
