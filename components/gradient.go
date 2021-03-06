package components

import (
	"fmt"
	"image/color"
	"regexp"
	"strings"

	"github.com/fogleman/gg"
	igcolor "github.com/xfangtong/image-generator/color"
)

// Gradient 渐变
type Gradient string

const (
	linearGradientRegexString = "^linear-gradient\\(([^ ]+) ([^,]+), ?(.*)\\)$"
	radialGradientRegexString = "^radial-gradient\\(([^ ]+) ([^ ]+) ([^ ]+) ([^ ]+) ([^ ]+) ([^,]+), ?(.*)\\)$"
	colorStopsRegexString     = "([^%]*) ([0-9\\.]+%)"
)

var (
	linearGradientRegex = regexp.MustCompile(linearGradientRegexString)
	radialGradientRegex = regexp.MustCompile(radialGradientRegexString)
	colorStopsRegex     = regexp.MustCompile(colorStopsRegexString)
)

func isLinearGradient(s string) bool {
	s = strings.ToLower(s)

	if !linearGradientRegex.MatchString(s) {
		return false
	}

	return true
}

func isRadialGradient(s string) bool {
	s = strings.ToLower(s)

	if !radialGradientRegex.MatchString(s) {
		return false
	}

	return true
}

func _parseColorStops(gradient gg.Gradient, clist string) error {
	cValues := colorStopsRegex.FindAllStringSubmatch(clist, -1)

	for _, v := range cValues {
		if len(v) != 3 {
			return fmt.Errorf("错误的颜色点格式")
		}
		sc := v[1]
		sc = strings.TrimPrefix(sc, ",")
		sc = strings.TrimPrefix(sc, " ")
		sp := v[2]
		dp, err := Dimension(sp).Measure(1.0)
		if err != nil {
			return fmt.Errorf("错误的颜色点位置")
		}
		c, err := igcolor.Color(sc).Parse()
		if err != nil {
			return fmt.Errorf("错误的颜色格式")
		}

		gradient.AddColorStop(dp, c)

	}
	return nil
}

func _parseLinearGradient(s string, refW float64, refH float64) (gg.Gradient, error) {
	values := linearGradientRegex.FindAllStringSubmatch(s, -1)
	if len(values) != 1 && len(values[0]) != 4 {
		return nil, fmt.Errorf("错误的渐变格式")
	}

	w, h := values[0][1], values[0][2]
	dw, err := Dimension(w).Measure(refW)
	if err != nil {
		return nil, err
	}
	dh, err := Dimension(h).Measure(refH)
	if err != nil {
		return nil, err
	}

	gradient := gg.NewLinearGradient(0, 0, dw, dh)

	clist := values[0][3]

	err = _parseColorStops(gradient, clist)
	if err != nil {
		return nil, err
	}

	return gradient, err
}

func _parseRadialGradient(s string, refW float64, refH float64) (gg.Gradient, error) {
	values := radialGradientRegex.FindAllStringSubmatch(s, -1)
	if len(values) != 1 && len(values[0]) != 8 {
		return nil, fmt.Errorf("错误的渐变格式")
	}

	x0, y0, r0, x1, y1, r1 := values[0][1], values[0][2], values[0][3], values[0][4], values[0][5], values[0][6]
	dx0, err := Dimension(x0).Measure(refW)
	if err != nil {
		return nil, err
	}
	dy0, err := Dimension(y0).Measure(refH)
	if err != nil {
		return nil, err
	}
	dr0, err := Dimension(r0).Measure(refW)
	if err != nil {
		return nil, err
	}
	dx1, err := Dimension(x1).Measure(refW)
	if err != nil {
		return nil, err
	}
	dy1, err := Dimension(y1).Measure(refH)
	if err != nil {
		return nil, err
	}
	dr1, err := Dimension(r1).Measure(refW)
	if err != nil {
		return nil, err
	}

	gradient := gg.NewRadialGradient(dx0, dy0, dr0, dx1, dy1, dr1)

	clist := values[0][7]

	err = _parseColorStops(gradient, clist)
	if err != nil {
		return nil, err
	}

	return gradient, err
}

func _parseGradient(s string, refW float64, refH float64) (gg.Pattern, error) {
	if isLinearGradient(s) {
		return _parseLinearGradient(s, refW, refH)
	}
	if isRadialGradient(s) {
		return _parseRadialGradient(s, refW, refH)
	}
	c, err := igcolor.Color(s).Parse()
	if err != nil {
		return nil, err
	}
	return gg.NewSolidPattern(c), nil

}

// IsGradient 是否为渐变
func (g Gradient) IsGradient() bool {
	s := string(g)
	return isLinearGradient(s) || isRadialGradient(s)
}

// GetColor 获取实颜色
func (g Gradient) GetColor() (color.Color, error) {
	return igcolor.Color(string(g)).Parse()
}

// SetFill 设置填充模式
func (g Gradient) SetFill(ctx *gg.Context, refW float64, refH float64) error {
	s := string(g)
	pattern, err := _parseGradient(s, refW, refH)
	if err != nil {
		return err
	}

	ctx.SetFillStyle(pattern)

	return nil
}

// SetStork 设置填充模式
func (g Gradient) SetStork(ctx *gg.Context, refW float64, refH float64) error {
	s := string(g)
	pattern, err := _parseGradient(s, refW, refH)
	if err != nil {
		return err
	}

	ctx.SetStrokeStyle(pattern)

	return nil
}
