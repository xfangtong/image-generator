package components

import (
	"image"
	"image/color"

	"github.com/skip2/go-qrcode"
	igcolor "github.com/xfangtong/image-generator/color"
)

type (
	// QrcodeComponentDefine 二维码
	QrcodeComponentDefine struct {
		Width   int           `json:"width"`
		Content string        `json:"content"`
		Color   igcolor.Color `json:"color"`
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

	fc, err := cd.Color.Parse()
	if err != nil {
		return err
	}

	gc.Push()
	defer gc.Pop()

	qc, err := qrcode.New(cd.Content, qrcode.Highest)
	if err != nil {
		return err
	}
	qc.BackgroundColor = color.Transparent
	qc.ForegroundColor = fc
	qc.DisableBorder = true

	img := qc.Image(cd.Width)
	gc.DrawImage(img, 0, 0)

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
