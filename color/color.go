package color

import (
	imgcolor "image/color"
	"strings"
)

// Color 颜色
type Color string

var ()

func parseColor(s string) (imgcolor.Color, error) {
	if s == "" || strings.ToLower(s) == "transparent" {
		return imgcolor.Transparent, nil
	}
	if len(s) < 4 {
		return nil, ErrBadColor
	}

	s = strings.ToLower(s)

	if s[:1] == "#" {
		return ParseHEX(s)
	} else if s[:4] == "rgba" {
		return ParseRGBA(s)
	} else if s[:3] == "rgb" {
		return ParseRGB(s)
	}

	return nil, ErrBadColor
}

func (c Color) IsGradient() bool {
	return false
}

// Parse 颜色解析
func (c Color) Parse() (imgcolor.Color, error) {
	return parseColor(string(c))
}

// ParseNoError 颜色解析
func (c Color) ParseNoError() imgcolor.Color {
	color, _ := parseColor(string(c))
	return color
}
