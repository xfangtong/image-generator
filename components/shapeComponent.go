package components

import (
	"github.com/fogleman/gg"
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
		LineWidth  int       `json:"lineWidth"`
		Dash       []float64 `json:"dash"`
		DashOffset float64   `json:"dashOffset"`
	}
)

// SetContextParameter 设置基础参数
func (shape *ShapeComponentDefine) SetContextParameter(dc *DrawContext) error {
	var err error = nil
	fc, err := shape.FillColor.Parse()
	if err != nil {
		return err
	}
	sc, err := shape.StrokeColor.Parse()
	if err != nil {
		return err
	}

	lw := float64(shape.LineWidth)
	gc := dc.GraphicContext

	gc.SetColor(fc)
	gc.SetLineWidth(lw)

	if len(shape.Dash) > 0 {
		gc.SetDash(shape.Dash...)
	}
	if shape.DashOffset != 0 {
		gc.SetDashOffset(shape.DashOffset)
	}

	gc.SetStrokeStyle(gg.NewSolidPattern(sc))
	gc.SetFillRuleWinding()

	return err
}
