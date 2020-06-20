package components

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Dimension 尺寸
type Dimension string

// AutoValue 自动值
const AutoValue int = math.MaxInt32

var (
	regPercent *regexp.Regexp = regexp.MustCompile("^(.*)%$")
	// ErrBadDimensionFormat 错误的尺寸
	ErrBadDimensionFormat = errors.New("bad dimension format")
)

// Measure 计算实际值
func (d Dimension) Measure(ref float64) (float64, error) {
	s := strings.ToLower(string(d))
	// auto 值为-1
	if s == "auto" {
		return float64(AutoValue), nil
	}

	ext := 0
	if strings.HasPrefix(s, "+") {
		s = strings.TrimPrefix(s, "+")
		ext = 1
	}

	if strings.HasPrefix(s, "-") {
		s = strings.TrimPrefix(s, "-")
		ext = -1
	}

	// 绝对值，直接返回
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
		if ext != 0 {
			return ref + float64(ext)*v, err
		}
		return v, err
	}

	perVals := regPercent.FindAllStringSubmatch(s, -1)
	if perVals != nil && len(perVals) > 0 && len(perVals[0]) == 2 {
		p, err := strconv.ParseFloat(perVals[0][1], 64)
		if err != nil {
			return 0, ErrBadDimensionFormat
		}

		return (p / 100) * ref, nil
	}

	return 0, ErrBadDimensionFormat
}

// MeasureNoError 计算实际值
func (d Dimension) MeasureNoError(ref float64) float64 {
	v, _ := d.Measure(ref)
	return v
}
