package colorutils

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHEX(t *testing.T) {
	s := "#fff"
	c, _ := Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "#ffffff"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "#ffffff00"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 0}, c)

	s = "#ffff0"
	c, err := Parse(s)
	assert.Equal(t, ErrBadColor, err)

}

func TestParseRGB(t *testing.T) {
	s := "rgb(255,255,255)"
	c, _ := Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "rgb(100%,100%,100%)"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "rgb(50%,50%,50%)"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 128, G: 128, B: 128, A: 255}, c)

}

func TestParseRGBA(t *testing.T) {
	s := "rgba(255,255,255,1)"
	c, _ := Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "rgba(100%,100%,100%,1)"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = "rgba(50%,50%,50%,1)"
	c, _ = Parse(s)
	assert.Equal(t, color.RGBA{R: 128, G: 128, B: 128, A: 255}, c)

}
