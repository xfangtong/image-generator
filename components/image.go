package components

import (
	"image"

	resources "github.com/xfangtong/image-generator/resources"
)

type (
	// ImageComponentDefine 图像组件定义
	ImageComponentDefine struct {
		// URL 图像地址
		URL resources.Resource `json:"url"`
	}

	// ImageComponent 图像组件
	ImageComponent struct {
		Component Component
	}
)

func (c *ImageComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*ImageComponentDefine)
	reader, err := cd.URL.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	cc, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	dc.GraphicContext.DrawImage(cc)

	return nil
}

func (c *ImageComponent) Measure(rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*ImageComponentDefine)
	reader, err := cd.URL.Open()
	if err != nil {
		return image.Rect(0, 0, 0, 0), err
	}
	defer reader.Close()

	cc, _, err := image.DecodeConfig(reader)
	if err != nil {
		return image.Rect(0, 0, 0, 0), err
	}

	return image.Rect(0, 0, cc.Width, cc.Height), nil
}

func (c *ImageComponent) ConfigType() interface{} {
	return &ImageComponentDefine{}
}

// Draw(c *DrawContext, config interface{}) error
// 		Measure(rect image.Rectangle, config interface{}) (image.Rectangle, error)
// 		ConfigType() interface{}
func init() {
	RegisterComponent("image", func() Component {
		return &ImageComponent{}
	})
}
