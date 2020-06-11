package colorutils

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

var (
	// ErrBadColor 颜色格式错误
	ErrBadColor = errors.New("Parsing of color failed, Bad Color")
)

const (
	hexRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	hexAlphaFormat = "#%02x%02x%02x%02x"
	hexFormat      = "#%02x%02x%02x"
	hexShortFormat = "#%1x%1x%1x"
	hexToRGBFactor = 17
)

var (
	hexRegex = regexp.MustCompile(hexRegexString)
)

// ParseHEX 解析十六进制格式颜色
func ParseHEX(s string) (color.Color, error) {

	s = strings.ToLower(s)

	if !hexRegex.MatchString(s) {
		return nil, ErrBadColor
	}

	var r, g, b, a uint8

	if len(s) == 4 {
		fmt.Sscanf(s, hexShortFormat, &r, &g, &b)
		r *= hexToRGBFactor
		g *= hexToRGBFactor
		b *= hexToRGBFactor
		a = 255
	} else if len(s) == 7 {
		fmt.Sscanf(s, hexFormat, &r, &g, &b)
		a = 255
	} else {
		fmt.Sscanf(s, hexAlphaFormat, &r, &g, &b, &a)
	}

	return color.RGBA{R: r, G: g, B: b, A: a}, nil
}
