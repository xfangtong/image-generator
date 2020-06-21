package components

import (
	"image"

	"github.com/fogleman/gg"
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

	lt, rt, rb, lb, err := cd.Radius.Parse(cd.Width, cd.Height)
	if err != nil {
		return nil
	}

	lw := float64(cd.ShapeComponentDefine.LineWidth)
	gc := dc.GraphicContext
	gc.Push()
	defer gc.Pop()
	err = cd.SetContextParameter(dc, float64(cd.Width), float64(cd.Height))
	if err != nil {
		return err
	}

	w, h := float64(cd.Width), float64(cd.Height)
	ltr, rtr, rbr, lbr := float64(lt), float64(rt), float64(rb), float64(lb)

	offsetLine := lw / 2.0

	gc.MoveTo(ltr, offsetLine)
	gc.LineTo(w-rtr, offsetLine)
	gc.DrawArc(w-rtr, rtr, rtr-offsetLine, gg.Radians(270), gg.Radians(360))
	gc.LineTo(w-offsetLine, h-rbr)
	gc.DrawArc(w-rbr, h-rbr, rbr-offsetLine, 0, gg.Radians(90))
	gc.LineTo(lbr, h-offsetLine)
	gc.DrawArc(lbr, h-lbr, lbr-offsetLine, gg.Radians(90), gg.Radians(180))
	gc.LineTo(offsetLine, ltr)
	gc.DrawArc(ltr, ltr, ltr-offsetLine, gg.Radians(180), gg.Radians(270))

	gc.FillPreserve()
	gc.StrokePreserve()

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
