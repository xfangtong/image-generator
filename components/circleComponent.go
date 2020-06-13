package components

import (
	"image"
	"math"
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

	gc := dc.GraphicContext
	gc.Save()
	gc.SetLineWidth(float64(cd.ShapeComponentDefine.LineWidth))
	gc.SetFillColor(fc)
	gc.SetStrokeColor(sc)
	gc.BeginPath()
	sx := cd.Radius
	gc.MoveTo(float64(sx), float64(0))
	gc.ArcTo(float64(sx), float64(0), float64(cd.Radius), float64(cd.Radius), (1/2)*math.Pi, (1/2)*math.Pi+math.Pi*2)
	gc.Close()
	gc.Stroke()
	gc.Fill()

	gc.Restore()

	return nil
}

// Measure 测量
func (c *CircleComponent) Measure(rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*CircleComponentDefine)

	return image.Rect(0, 0, cd.Radius*2+cd.ShapeComponentDefine.LineWidth/2, cd.Radius*2+cd.ShapeComponentDefine.LineWidth/2), nil
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
