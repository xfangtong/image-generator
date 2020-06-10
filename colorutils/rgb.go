package colorutils

import (
	"image/color"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	rgbString                    = "rgb(%d,%d,%d)"
	rgbCaptureRegexString        = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	rgbCaptureRegexPercentString = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*\\)$"
)

var (
	rgbCaptureRegex        = regexp.MustCompile(rgbCaptureRegexString)
	rgbCapturePercentRegex = regexp.MustCompile(rgbCaptureRegexPercentString)
)

// parseRGB 解析RGB格式
func parseRGB(s string) (color.Color, error) {

	s = strings.ToLower(s)

	var isPercent bool
	vals := rgbCaptureRegex.FindAllStringSubmatch(s, -1)

	if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {

		vals = rgbCapturePercentRegex.FindAllStringSubmatch(s, -1)

		if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {
			return nil, ErrBadColor
		}

		isPercent = true
	}

	r, _ := strconv.ParseUint(vals[0][1], 10, 8)
	g, _ := strconv.ParseUint(vals[0][2], 10, 8)
	b, _ := strconv.ParseUint(vals[0][3], 10, 8)

	if isPercent {
		r = uint64(math.Floor(float64(r)/100*255 + .5))
		g = uint64(math.Floor(float64(g)/100*255 + .5))
		b = uint64(math.Floor(float64(b)/100*255 + .5))
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}
