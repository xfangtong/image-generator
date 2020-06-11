package imagegenerator

import (
	"errors"
	"regexp"
	"strconv"
)

// Dimension 尺寸
type Dimension string

var (
	regPercent *regexp.Regexp = regexp.MustCompile("^(.*)%$")
	// ErrBadDimensionFormat 错误的尺寸
	ErrBadDimensionFormat = errors.New("bad dimension format")
)

// Measure 计算实际值
func (d Dimension) Measure(ref float64) (float64, error) {
	s := string(d)
	// 绝对值，直接返回
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
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
