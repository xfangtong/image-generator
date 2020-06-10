package colorutils

import (
	"errors"
	"image/color"
	"strings"
)

var (
	// ErrBadColor 颜色格式错误
	ErrBadColor = errors.New("Parsing of color failed, Bad Color")
)

// Parse 颜色解析
func Parse(s string) (color.Color, error) {

	if len(s) < 4 {
		return nil, ErrBadColor
	}

	s = strings.ToLower(s)

	if s[:1] == "#" {
		return parseHEX(s)
	} else if s[:4] == "rgba" {
		return parseRGBA(s)
	} else if s[:3] == "rgb" {
		return parseRGB(s)
	}

	return nil, ErrBadColor
}
