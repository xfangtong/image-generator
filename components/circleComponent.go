package components

import (
	"image"
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

	lw := float64(cd.ShapeComponentDefine.LineWidth)

	gc := dc.GraphicContext

	gc.Push()
	defer gc.Pop()
	err = cd.SetContextParameter(dc)
	if err != nil {
		return err
	}

	x := float64(cd.Radius)
	gc.DrawCircle(x, x, x-lw/2)

	gc.FillPreserve()
	gc.StrokePreserve()

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
