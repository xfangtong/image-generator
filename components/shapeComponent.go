package components

type (
	// ShapeComponentDefine 形状组件
	ShapeComponentDefine struct {
		// FillColor 填充颜色
		FillColor Gradient `json:"fillColor"`
		// StrokeColor 线条颜色
		StrokeColor Gradient `json:"strokeColor"`
		// LineWidth 线条宽度
		LineWidth  int       `json:"lineWidth"`
		Dash       []float64 `json:"dash"`
		DashOffset float64   `json:"dashOffset"`
	}
)

// SetContextParameter 设置基础参数
func (shape *ShapeComponentDefine) SetContextParameter(dc *DrawContext, width float64, height float64) error {
	var err error = nil
	// fc, err := shape.FillColor.Parse()
	// if err != nil {
	// 	return err
	// }
	// sc, err := shape.StrokeColor.Parse()
	// if err != nil {
	// 	return err
	// }

	err = shape.FillColor.SetFill(dc.GraphicContext, width, height)
	if err != nil {
		return err
	}
	lw := float64(shape.LineWidth)
	gc := dc.GraphicContext

	//gc.SetColor(fc)
	gc.SetLineWidth(lw)

	if len(shape.Dash) > 0 {
		gc.SetDash(shape.Dash...)
	}
	if shape.DashOffset != 0 {
		gc.SetDashOffset(shape.DashOffset)
	}

	err = shape.StrokeColor.SetStork(dc.GraphicContext, width, height)
	//gc.SetStrokeStyle(gg.NewSolidPattern(sc))
	gc.SetFillRuleWinding()

	return err
}
