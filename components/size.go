package components

import (
	"errors"
	"image"
	"math"
	"strings"
)

// Size 尺寸
type Size string

// Parse 解析位置
func (p Size) Parse(canvasRect image.Rectangle, imgRect image.Rectangle) (image.Rectangle, error) {
	s := string(p)
	s = strings.ToLower(s)

	// contain 缩放背景图片以完全装入背景区 cover 缩放背景图片以完全覆盖背景区
	if s == "contain" || s == "cover" {
		rateX := float64(canvasRect.Dx()) / float64(imgRect.Dx())
		rateY := float64(canvasRect.Dy()) / float64(imgRect.Dy())

		rate := 1.0
		if s == "contain" {
			rate = math.Min(rateX, rateY)
		} else {
			rate = math.Max(rateX, rateY)
		}

		return image.Rect(int(rate*float64(imgRect.Min.X)), int(rate*float64(imgRect.Min.Y)), int(rate*float64(imgRect.Max.X)), int(rate*float64(imgRect.Max.Y))), nil
	}

	rect := image.Rect(imgRect.Min.X, imgRect.Min.Y, imgRect.Max.X, imgRect.Max.Y)

	x, y := "auto", "auto"
	vals := strings.Split(s, " ")
	if len(vals) > 2 {
		return image.Rect(0, 0, 0, 0), errors.New("invalid format")
	} else if len(vals) == 1 {
		x = vals[0]
	} else {
		x = vals[0]
		y = vals[1]
	}

	// auto 使用画布的值. 百分比使用图像百分比
	if x == "auto" {
		rect.Max.X = canvasRect.Dx() + rect.Min.X
	} else {
		x1, err := Dimension(x).Measure(float64(imgRect.Dx()))
		if err != nil {
			return rect, err
		}
		rect.Max.X = int(x1 + float64(rect.Min.X))
	}

	if y == "auto" {
		rect.Max.Y = canvasRect.Dy() + rect.Min.Y
	} else {
		y1, err := Dimension(y).Measure(float64(imgRect.Dy()))
		if err != nil {
			return rect, err
		}
		rect.Max.Y = int(y1 + float64(rect.Min.Y))
	}

	return rect, nil

}

// ParseNoError 解析位置，忽略错误
func (p Size) ParseNoError(canvasRect image.Rectangle, imgRect image.Rectangle) image.Rectangle {
	rect, _ := p.Parse(canvasRect, imgRect)
	return rect
}
