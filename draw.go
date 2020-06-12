package imagegenerator

import (
	"errors"
	"image"
	_ "image/jpeg" //注入解码
	_ "image/png"  //注入解码

	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	// ErrInvalidWidth 无效宽度
	ErrInvalidWidth = errors.New("invalid width")
	// ErrInvalidHeight 无效高度
	ErrInvalidHeight = errors.New("invalid hight")
)

type (
	// DrawContext 上下文
	DrawContext struct {
		GraphicContext *draw2dimg.GraphicContext
		Width          int
		Height         int
		CurrentLeft    int
		CurrentTop     int
	}
)

// GenerateImage 生成图像
func GenerateImage(t ImageTemplate) (image.Image, error) {
	if t.Width < 0 || (t.Width == 0 && t.Background == "") {
		return nil, ErrInvalidWidth
	}
	if t.Height < 0 || (t.Height == 0 && t.Background == "") {
		return nil, ErrInvalidHeight
	}

	var backgroundImage image.Image = nil
	if t.Background != "" {
		biReader, err := t.Background.Open()
		if err != nil {
			return nil, err
		}
		defer biReader.Close()
		backgroundImage, _, err = image.Decode(biReader)
		if err != nil {
			return nil, err
		}
		if t.Width == 0 {
			t.Width = backgroundImage.Bounds().Dx()
		}
		if t.Height == 0 {
			t.Height = backgroundImage.Bounds().Dy()
		}
	}

	//w, h := float64(t.Width), float64(t.Height)

	// 生成底图
	imgRect := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: t.Width, Y: t.Height}}
	img := image.NewRGBA(imgRect)

	gc := draw2dimg.NewGraphicContext(img)

	if t.BackgroundColor != "" {
		bgColor, err := t.BackgroundColor.Parse()
		if err != nil {
			return nil, errors.New("invalid background color value")
		}
		gc.Save()
		gc.SetFillColor(bgColor)
		gc.Clear()

		gc.Restore()
	}

	if backgroundImage != nil {
		//draw2dimg.DrawImage()
	}

	img2 := image.NewRGBA(image.Rect(0, 0, 150, 150))
	gc2 := draw2dimg.NewGraphicContext(img2)
	gc2.SetFillColor(Color("#00000000").ParseNoError())
	gc2.Clear()
	// gc2.BeginPath()
	// gc2.MoveTo(0.0, 0.0)
	// gc2.LineTo(150, 0)
	// gc2.LineTo(150, 150)
	// gc2.LineTo(0, 150)
	// gc2.Close()
	// gc2.Fill()

	//gc2.SetFillColor(Color("#00FF00").ParseNoError())

	//gc2.SetFontSize(24)
	//gc2.FillString("HELLO World!")

	//draw.Draw(img, image.Rect(0, 0, 150, 150), img2, image.Pt(0, 0), draw.Over)
	//gc.Save()

	return img, nil
}
