package components

import (
	"image"
	"image/color"

	"github.com/skip2/go-qrcode"
)

type (
	// QrcodeComponentDefine 二维码
	QrcodeComponentDefine struct {
		Width   int      `json:"width"`
		Content string   `json:"content"`
		Color   Gradient `json:"color"`
	}

	// QrcodeComponent 二维码
	QrcodeComponent struct {
	}
)

// Draw 绘制
func (c *QrcodeComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*QrcodeComponentDefine)

	var err error = nil
	gc := dc.GraphicContext

	// fc, err := cd.Color.Parse()
	// if err != nil {
	// 	return err
	// }

	gc.Push()
	defer gc.Pop()

	ig := cd.Color.IsGradient()

	var fc color.Color = color.Black
	if ig {
		// mask
		dc.GraphicContext.SetColor(color.Transparent)
		dc.GraphicContext.Clear()
		//dc.GraphicContext.SetRGB(0, 0, 0)
	} else {
		fc, err = cd.Color.GetColor()
		if err != nil {
			return err
		}
		//	dc.GraphicContext.SetColor(fc)
	}

	qc, err := qrcode.New(cd.Content, qrcode.Highest)
	if err != nil {
		return err
	}
	qc.BackgroundColor = color.Transparent
	qc.ForegroundColor = fc
	qc.DisableBorder = true

	img := qc.Image(cd.Width)
	gc.DrawImage(img, 0, 0)

	if ig {
		mask := gc.AsMask()
		dc.GraphicContext.Clear()
		err = cd.Color.SetFill(dc.GraphicContext, float64(cd.Width), float64(cd.Width))
		if err != nil {
			return err
		}
		dc.GraphicContext.SetMask(mask)
		dc.GraphicContext.DrawRectangle(0, 0, float64(cd.Width), float64(cd.Width))
		dc.GraphicContext.Fill()
	}

	return err
}

// Measure 测量
func (c *QrcodeComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*QrcodeComponentDefine)

	return image.Rect(0, 0, cd.Width, cd.Width), nil
}

// ConfigType 配置类型
func (c *QrcodeComponent) ConfigType() interface{} {
	return &QrcodeComponentDefine{}
}

func init() {
	RegisterComponent("qrcode", func() Component {
		return &QrcodeComponent{}
	})
}
