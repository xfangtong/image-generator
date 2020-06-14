package components

import (
	"image"
	"math"

	"github.com/llgcode/draw2d"
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
	gc.Save()
	gc.SetLineWidth(lw)
	gc.SetFillColor(fc)
	gc.SetStrokeColor(sc)
	gc.SetFillRule(draw2d.FillRuleWinding)
	gc.BeginPath()

	x := float64(cd.Radius)
	gc.ArcTo(float64(x), float64(x), float64(cd.Radius)-lw, float64(cd.Radius)-lw, 0, 2*math.Pi)
	gc.FillStroke()

	gc.Restore()

	return nil
}

// Measure 测量
func (c *CircleComponent) Measure(rect image.Rectangle, config interface{}) (image.Rectangle, error) {
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
