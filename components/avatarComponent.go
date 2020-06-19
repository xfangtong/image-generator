package components

import (
	"image"

	resources "github.com/xfangtong/image-generator/resources"
)

type (
	// AvatarComponentDefine 头像
	AvatarComponentDefine struct {
		ShapeComponentDefine
		// URL 图像地址
		URL resources.Resource `json:"url"`
		// Width 宽度
		Width int `json:"width"`
		// Size 头像尺寸
		Size Size `json:"size"`
	}
	// AvatarComponent 头像组件
	AvatarComponent struct {
	}
)

// Draw 绘制
func (c *AvatarComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*AvatarComponentDefine)

	var err error = nil

	gc := dc.GraphicContext
	gc.Push()
	defer gc.Pop()

	err = cd.SetContextParameter(dc)
	if err != nil {
		return err
	}

	imgReader, err := cd.URL.Open()
	if err != nil {
		return err
	}
	defer imgReader.Close()

	img, _, err := image.Decode(imgReader)
	if err != nil {
		return err
	}

	mRect := image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())
	drawRect := image.Rect(0, 0, cd.Width-cd.LineWidth*2, cd.Width-cd.LineWidth*2)

	// 组件实际尺寸
	aRect, err := cd.Size.Parse(drawRect, mRect)
	if err != nil {
		return err
	}

	x, y, err := Position("center").Parse(drawRect, aRect)
	if err != nil {
		return err
	}

	cx, cy := float64(cd.Width)/2.0, float64(cd.Width)/2.0

	gc.DrawCircle(cx, cy, cx-float64(cd.LineWidth/2))
	gc.FillPreserve()
	if cd.LineWidth > 0 {
		gc.Stroke()
	}
	gc.ClearPath()

	gc.DrawCircle(cx, cy, cx-float64(cd.LineWidth))
	gc.Clip()
	gc.AsMask()

	sx, sy := float64(aRect.Dx())/float64(mRect.Dx()), float64(aRect.Dy())/float64(mRect.Dy())
	gc.Translate(x+float64(cd.LineWidth), y+float64(cd.LineWidth))
	gc.Scale(sx, sy)
	gc.DrawImage(img, 0, 0)

	return err
}

// Measure 测量
func (c *AvatarComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*AvatarComponentDefine)

	return image.Rect(0, 0, cd.Width, cd.Width), nil
}

// ConfigType 配置类型
func (c *AvatarComponent) ConfigType() interface{} {
	return &AvatarComponentDefine{}
}

func init() {
	RegisterComponent("avatar", func() Component {
		return &AvatarComponent{}
	})
}
