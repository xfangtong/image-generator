package components

import (
	"image"
	"math"

	"github.com/llgcode/draw2d"
)

type (
	// RectComponentDefine  矩形
	RectComponentDefine struct {
		ShapeComponentDefine
		// Width 宽度
		Width int
		// Height 高度
		Height int
		// Radius 圆角 左上 右上 右下 左下 | 给定两个值：左上 右下 右上 左下 | 给定一个值： 四角相同
		Radius Padding
	}

	// RectComponent 矩形
	RectComponent struct {
	}
)

// Draw 绘制
func (c *RectComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*RectComponentDefine)

	var err error = nil
	fc, err := cd.ShapeComponentDefine.FillColor.Parse()
	if err != nil {
		return err
	}
	sc, err := cd.ShapeComponentDefine.StrokeColor.Parse()
	if err != nil {
		return err
	}

	lt, rt, rb, lb, err := cd.Radius.Parse(cd.Width, cd.Height)
	if err != nil {
		return nil
	}

	lw := float64(cd.ShapeComponentDefine.LineWidth)
	gc := dc.GraphicContext
	gc.Save()
	gc.SetLineWidth(lw)
	gc.SetFillColor(fc)
	gc.SetStrokeColor(sc)
	gc.SetFillRule(draw2d.FillRuleWinding)
	gc.BeginPath()

	w, h := float64(cd.Width), float64(cd.Height)
	ltr, rtr, rbr, lbr := float64(lt), float64(rt), float64(rb), float64(lb)

	offsetLine := lw / 2.0
	angle := math.Pi / 2.0

	gc.MoveTo(ltr, offsetLine)
	gc.LineTo(w-rtr, offsetLine)
	gc.ArcTo(w-rtr, rtr, rtr-offsetLine, rtr-offsetLine, (math.Pi/2)*3, angle)
	gc.LineTo(w-offsetLine, h-rbr)
	gc.ArcTo(w-rbr, h-rbr, rbr-offsetLine, rbr-offsetLine, 0, angle)
	gc.LineTo(lbr, h-offsetLine)
	gc.ArcTo(lbr, h-lbr, lbr-offsetLine, lbr-offsetLine, math.Pi/2.0, angle)
	gc.LineTo(offsetLine, ltr)
	gc.ArcTo(ltr, ltr, ltr-offsetLine, ltr-offsetLine, math.Pi, angle)

	gc.FillStroke()

	gc.Restore()

	return nil
}

// Measure 测量
func (c *RectComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*RectComponentDefine)

	return image.Rect(0, 0, cd.Width, cd.Height), nil
}

// ConfigType 配置类型
func (c *RectComponent) ConfigType() interface{} {
	return &RectComponentDefine{}
}

func init() {
	RegisterComponent("rect", func() Component {
		return &RectComponent{}
	})
}
