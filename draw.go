package imagegenerator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg" //注入解码
	_ "image/png"  //注入解码
	"strconv"
	"text/template"

	"github.com/fogleman/gg"
	"github.com/xfangtong/image-generator/components"
)

var (
	// ErrInvalidWidth 无效宽度
	ErrInvalidWidth = errors.New("invalid width")
	// ErrInvalidHeight 无效高度
	ErrInvalidHeight = errors.New("invalid hight")
)

// GenerateImage 生成图像
func GenerateImage(t ImageTemplate) (image.Image, error) {
	var err error = nil
	var bgImg image.Config
	hasBgImg := false
	refW, refH := 0.0, 0.0
	if t.Background != "" {
		if bgReader, err := t.Background.Open(); err == nil {
			if bgImg, _, err = image.DecodeConfig(bgReader); err != nil {
				bgReader.Close()
				return nil, err
			}
			bgReader.Close()
			hasBgImg = true
			refW = float64(bgImg.Width)
			refH = float64(bgImg.Height)
		} else {
			return nil, err
		}

	}

	w, err := t.Width.Measure(refW)
	if err != nil || w <= 0 {
		return nil, ErrInvalidWidth
	}
	h, err := t.Height.Measure(refH)
	if err != nil || h <= 0 {
		return nil, ErrInvalidHeight
	}

	isAutoHeight := false
	if !hasBgImg && w != float64(components.AutoValue) && h == float64(components.AutoValue) {
		h = w * 10
		isAutoHeight = true
	}

	if w == float64(components.AutoValue) && h == float64(components.AutoValue) && !hasBgImg {
		return nil, fmt.Errorf("宽度和高度不可同时为auto")
	}
	if w == float64(components.AutoValue) {
		w = refW
	}
	if h == float64(components.AutoValue) {
		h = refH
	}

	imgRect := image.Rect(0, 0, int(w), int(h))
	img := image.NewRGBA(imgRect)
	gc := gg.NewContextForRGBA(img)

	dw, dh := components.Dimension(strconv.FormatFloat(w, 'f', -1, 64)), components.Dimension(strconv.FormatFloat(h, 'f', -1, 64))

	bgDc := &components.DrawContext{
		GraphicContext: gc,
		Image:          img,
		Width:          int(w),
		Height:         int(h),
		CurrentLeft:    0,
		CurrentTop:     0,
	}

	err = t.BackgroundColor.SetFill(gc, w, h)
	if err != nil {
		return nil, err
	}
	//gc.SetColor(bgColor)
	//gc.Clear()
	gc.DrawRectangle(0, 0, w, h)
	gc.Fill()

	// 背景
	if hasBgImg {
		bgImgComponent := components.ComponentDefine{
			Type: "image",
			//BackgroundColor: t.BackgroundColor,
			Size:     t.Size,
			Position: t.Position,
			Padding:  t.BackgroundPadding,
			Repeat:   t.Repeat,
			Area: components.Rectangle{
				Left:   "0",
				Top:    "0",
				Right:  dw,
				Bottom: dh,
			},
			ComponentData: map[string]interface{}{
				"url": t.Background,
			},
		}

		if _, err = bgDc.Clone().DrawComponent(bgImgComponent); err != nil {
			return nil, err
		}
	}

	if len(t.Components) > 0 {
		// 绘制组件
		groupComponent := components.ComponentDefine{
			Type:            "group",
			BackgroundColor: "transparent",
			Size:            "100% 100%",
			Position:        "left top",
			Repeat:          components.RepeatNO,
			Padding:         t.Padding,
			Area: components.Rectangle{
				Left:   "0",
				Top:    "0",
				Right:  dw,
				Bottom: dh,
			},
			ComponentData: map[string]interface{}{
				"components": t.Components,
			},
		}
		rect, err := bgDc.DrawComponent(groupComponent)
		if err != nil {
			return nil, err
		}

		if isAutoHeight {
			_, _, b, _, _ := t.Padding.Parse(int(w), rect.Max.Y)
			rect.Max.X = int(w)
			rect.Max.Y = bgDc.ActualHeight + b

			aImg := image.NewRGBA(rect)
			agc := gg.NewContextForRGBA(aImg)
			agc.DrawImage(bgDc.Image, 0, 0)
			img = aImg
		}
	}

	return img, nil
}

//GenerateImageFromTemplate 从模版生成图像
func GenerateImageFromTemplate(text string, values map[string]interface{}) (image.Image, error) {
	t := template.New("image-template")
	tpl, err := t.Parse(text)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, values)
	if err != nil {
		return nil, err
	}

	imgTpl := &ImageTemplate{}
	if err = json.Unmarshal(buffer.Bytes(), imgTpl); err != nil {
		return nil, err
	}

	return GenerateImage(*imgTpl)

}
