package components

import (
	"errors"
	"strings"
)

// Padding 间距 上 右 下 左
type Padding string

// Parse Padding解析
func (p Padding) Parse(refW int, refH int) (int, int, int, int, error) {
	s := string(p)
	if s == "" {
		return 0, 0, 0, 0, nil
	}

	vals := strings.Split(s, " ")
	pt, pr, pb, pl := "0", "0", "0", "0"
	if len(vals) == 1 {
		pt = vals[0]
		pr = pt
		pb = pt
		pl = pt
	} else if len(vals) == 2 {
		pt = vals[0]
		pb = pt
		pr = vals[1]
		pl = pr
	} else if len(vals) == 3 {
		pt = vals[0]
		pr = vals[1]
		pb = vals[2]
	} else if len(vals) == 4 {
		pt = vals[0]
		pr = vals[1]
		pb = vals[2]
		pl = vals[3]
	} else {
		return 0, 0, 0, 0, errors.New("invalid padding format")
	}

	var err error = nil
	pti, pri, pbi, pli := 0.0, 0.0, 0.0, 0.0
	pti, err = Dimension(pt).Measure(float64(refH))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	pri, err = Dimension(pr).Measure(float64(refW))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	pbi, err = Dimension(pb).Measure(float64(refH))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	pli, err = Dimension(pl).Measure(float64(refW))
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return int(pti), int(pri), int(pbi), int(pli), nil
}

// ParseNoError 解析
func (p Padding) ParseNoError(refW int, refH int) (int, int, int, int) {
	pti, pri, pbi, pli, _ := p.Parse(refW, refH)
	return pti, pri, pbi, pli
}
