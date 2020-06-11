package imagegenerator

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xfangtong/image-generator/colorutils"
)

func TestParseHEX(t *testing.T) {
	s := Color("#fff")
	c, _ := s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("#ffffff")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("#ffffff00")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 0}, c)

	s = Color("#ffff0")
	c, err := s.Parse()
	assert.Equal(t, colorutils.ErrBadColor, err)

}

func TestParseRGB(t *testing.T) {
	s := Color("rgb(255,255,255)")
	c, _ := s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("rgb(100%,100%,100%)")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("rgb(50%,50%,50%)")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 128, G: 128, B: 128, A: 255}, c)

}

func TestParseRGBA(t *testing.T) {
	s := Color("rgba(255,255,255,1)")
	c, _ := s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("rgba(100%,100%,100%,1)")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, c)

	s = Color("rgba(50%,50%,50%,1)")
	c, _ = s.Parse()
	assert.Equal(t, color.RGBA{R: 128, G: 128, B: 128, A: 255}, c)

}
