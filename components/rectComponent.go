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
		// Radius 圆角
		Radius int
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

	lw := float64(cd.ShapeComponentDefine.LineWidth)
	gc := dc.GraphicContext
	gc.Save()
	gc.SetLineWidth(lw)
	gc.SetFillColor(fc)
	gc.SetStrokeColor(sc)
	gc.SetFillRule(draw2d.FillRuleWinding)
	gc.BeginPath()

	w, h, r := float64(cd.Width), float64(cd.Height), float64(cd.Radius)

	gc.MoveTo(r, 0)
	gc.LineTo(w-r, 0)

	//
	//r = r - lw
	if r > 0 {
		gc.ArcTo(w-r/2.0, r/2.0, r/2.0, r/2.0, (math.Pi/180.0)*270.0, 0)
	}

	//gc.LineTo(w, r)

	gc.LineTo(w, h-r)

	gc.LineTo(w-r, h)
	gc.LineTo(r, h)

	gc.LineTo(0, h-r)
	gc.LineTo(0, r)

	gc.LineTo(r, 0)

	gc.FillStroke()

	gc.Restore()

	return nil
}

// Measure 测量
func (c *RectComponent) Measure(rect image.Rectangle, config interface{}) (image.Rectangle, error) {
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
