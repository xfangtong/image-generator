package components

import (
	"image"
)

type (
	// LineComponentDefine 线条
	LineComponentDefine struct {
		ShapeComponentDefine
		X int `json:"x"`
		Y int `json:"y`
	}

	// LineComponent 线条
	LineComponent struct {
	}
)

// Draw 绘制
func (c *LineComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*LineComponentDefine)

	var err error = nil
	gc := dc.GraphicContext

	gc.Push()
	defer gc.Pop()
	if err = cd.SetContextParameter(dc); err != nil {
		return err
	}

	gc.MoveTo(0, 0)
	gc.LineTo(float64(cd.X-cd.LineWidth), float64(cd.Y-cd.LineWidth))

	gc.Stroke()

	return nil
}

// Measure 测量
func (c *LineComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*LineComponentDefine)

	return image.Rect(0, 0, cd.X, cd.Y), nil
}

// ConfigType 配置类型
func (c *LineComponent) ConfigType() interface{} {
	return &LineComponentDefine{}
}

func init() {
	RegisterComponent("line", func() Component {
		return &LineComponent{}
	})
}
