package components

import (
	igcolor "github.com/xfangtong/image-generator/color"
)

type (
	// ShapeComponentDefine 形状组件
	ShapeComponentDefine struct {
		// FillColor 填充颜色
		FillColor igcolor.Color `json:"fillColor"`
		// StrokeColor 线条颜色
		StrokeColor igcolor.Color `json:"strokeColor"`
		// LineWidth 线条宽度
		LineWidth int `json:"lineWidth"`
	}
)