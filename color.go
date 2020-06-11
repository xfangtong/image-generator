package imagegenerator

import (
	"image/color"
	"strings"

	"github.com/xfangtong/image-generator/colorutils"
)

// Color 颜色
type Color string

func parseColor(s string) (color.Color, error) {
	if len(s) < 4 {
		return nil, colorutils.ErrBadColor
	}

	s = strings.ToLower(s)

	if s[:1] == "#" {
		return colorutils.ParseHEX(s)
	} else if s[:4] == "rgba" {
		return colorutils.ParseRGBA(s)
	} else if s[:3] == "rgb" {
		return colorutils.ParseRGB(s)
	}

	return nil, colorutils.ErrBadColor
}

// Parse 颜色解析
func (c Color) Parse() (color.Color, error) {
	return parseColor(string(c))
}

// ParseNoError 颜色解析
func (c Color) ParseNoError() color.Color {
	color, _ := parseColor(string(c))
	return color
}
