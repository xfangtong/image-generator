package components

import (
	"image"

	"github.com/fogleman/gg"
)

type (
	// CircleComponentDefine 圆
	CircleComponentDefine struct {
		ShapeComponentDefine
		Radius int `json:"radius"`
	}

	// CircleComponent 圆
	CircleComponent struct {
	}
)

// Draw 绘制
func (c *CircleComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*CircleComponentDefine)

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

	gc.Push()
	gc.SetColor(fc)
	gc.SetLineWidth(lw)
	// gc.SetFillColor(fc)
	// gc.SetStrokeColor(sc)
	// gc.SetFillRule(draw2d.FillRuleWinding)
	// gc.BeginPath()

	x := float64(cd.Radius)
	gc.DrawCircle(x, x, x-lw/2)

	//	gc.ArcTo(float64(x), float64(x), float64(cd.Radius)-lw/2, float64(cd.Radius)-lw/2, 0, 2*math.Pi)
	gc.SetStrokeStyle(gg.NewSolidPattern(sc))
	//	gc.Fill()
	//gc.Stroke()
	//gc.Fill()
	gc.FillPreserve()
	gc.StrokePreserve()
	//	gc.FillStroke()

	gc.Pop()

	return nil
}

// Measure 测量
func (c *CircleComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*CircleComponentDefine)

	return image.Rect(0, 0, cd.Radius*2, cd.Radius*2), nil
}

// ConfigType 配置类型
func (c *CircleComponent) ConfigType() interface{} {
	return &CircleComponentDefine{}
}

func init() {
	RegisterComponent("circle", func() Component {
		return &CircleComponent{}
	})
}
