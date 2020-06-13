package components

import "image"

type (
	// Rectangle 区域
	Rectangle struct {
		Left   Dimension `json:"left"`
		Top    Dimension `json:"top"`
		Right  Dimension `json:"right"`
		Bottom Dimension `json:"bottom"`
	}
)

// Parse 解析
func (r Rectangle) Parse(refRect image.Rectangle) (image.Rectangle, error) {

	result := image.Rect(0, 0, 0, 0)

	l, err := r.Left.Measure(float64(refRect.Dx()))
	if err != nil {
		return result, err
	}
	result.Min.X = int(l)
	t, err := r.Top.Measure(float64(refRect.Dy()))
	if err != nil {
		return result, err
	}
	result.Min.Y = int(t)

	ri, err := r.Right.Measure(float64(refRect.Dx()))
	if err != nil {
		return result, err
	}
	result.Max.X = int(ri)
	bo, err := r.Bottom.Measure(float64(refRect.Dy()))
	if err != nil {
		return result, err
	}
	result.Max.Y = int(bo)

	return result, nil
}

// ParseNoError 解析，忽略错误
func (r Rectangle) ParseNoError(refRect image.Rectangle) image.Rectangle {
	result, _ := r.Parse(refRect)

	return result
}
