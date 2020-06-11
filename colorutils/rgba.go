package colorutils

import (
	"image/color"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	rgbaString                    = "rgba(%d,%d,%d,%g)"
	rgbaCaptureRegexString        = "^rgba\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0\\.[0-9]*|[01])\\s*\\)$"
	rgbaCaptureRegexPercentString = "^rgba\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0\\.[0-9]*|[01])\\s*\\)$"
)

var (
	rgbaCaptureRegex        = regexp.MustCompile(rgbaCaptureRegexString)
	rgbaCapturePercentRegex = regexp.MustCompile(rgbaCaptureRegexPercentString)
)

// ParseRGBA 解析RGBA格式
func ParseRGBA(s string) (color.Color, error) {

	s = strings.ToLower(s)

	var isPercent bool

	vals := rgbaCaptureRegex.FindAllStringSubmatch(s, -1)

	if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {

		vals = rgbaCapturePercentRegex.FindAllStringSubmatch(s, -1)

		if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {
			return nil, ErrBadColor
		}

		isPercent = true
	}

	r, _ := strconv.ParseUint(vals[0][1], 10, 8)
	g, _ := strconv.ParseUint(vals[0][2], 10, 8)
	b, _ := strconv.ParseUint(vals[0][3], 10, 8)
	a, _ := strconv.ParseFloat(vals[0][4], 64)

	if isPercent {
		r = uint64(math.Floor(float64(r)/100*255 + .5))
		g = uint64(math.Floor(float64(g)/100*255 + .5))
		b = uint64(math.Floor(float64(b)/100*255 + .5))
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a * 255)}, nil
}
