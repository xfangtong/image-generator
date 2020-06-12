package imagegenerator

import (
	"errors"
	"image"
	"strings"
)

// Position 表示位置
type Position string

var (
	xPosition = map[string]string{"left": "0%", "center": "50%", "right": "100%"}
	yPosition = map[string]string{"top": "0%", "center": "50%", "bottom": "100%"}
)

// Parse 解析位置
func (p Position) Parse(canvasRect image.Rectangle, imgRect image.Rectangle) (float64, float64, error) {
	s := string(p)
	y, x := "center", "center"
	vals := strings.Split(s, " ")
	if len(vals) > 2 {
		return 0, 0, errors.New("invalid format")
	} else if len(vals) == 1 {
		y = vals[0]
	} else {
		y = vals[0]
		x = vals[1]
	}
	y = strings.ToLower(y)
	x = strings.ToLower(x)

	y1, ok := yPosition[y]
	if ok {
		y = y1
	}
	x1, ok := xPosition[x]
	if ok {
		x = x1
	}

	dx := canvasRect.Dx() - imgRect.Dx()
	dy := canvasRect.Dy() - imgRect.Dy()

	fy, err := Dimension(y).Measure(float64(dy))
	if err != nil {
		return 0, 0, err
	}
	fx, err := Dimension(x).Measure(float64(dx))
	if err != nil {
		return 0, 0, err
	}

	return fx, fy, nil

}

// ParseNoError 解析位置，忽略错误
func (p Position) ParseNoError(canvasRect image.Rectangle, imgRect image.Rectangle) (float64, float64) {
	x, y, _ := p.Parse(canvasRect, imgRect)
	return x, y
}
