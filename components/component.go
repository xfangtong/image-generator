package components

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"sync"
	"sync/atomic"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	igcolor "github.com/xfangtong/image-generator/color"
)

type (
	// ComponentDefine 组件定义
	ComponentDefine struct {
		// Type 类型
		Type string `json:"type"`
		// Level 层次
		Level int `json:"level"`
		// Area 放置区域
		Area Rectangle `json:"area"`
		// Position 在区域中的位置方式
		Position Position `json:"position"`
		// Repeat 重复方式
		Repeat string `json:"repeat"`
		// Size 组件尺寸方式
		Size Size `json:"size"`
		// Padding 间隔空白
		Padding Padding `json:"padding"`
		// BackgroundColor 背景颜色
		BackgroundColor igcolor.Color `json:"backgroundColor"`
		// ComponentData 组件数据，根据组件类型具有不同的结构
		ComponentData map[string]interface{} `json:"componentData"`
	}

	// Component 组件接口
	Component interface {
		Draw(c *DrawContext, config interface{}) error
		Measure(c *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error)
		ConfigType() interface{}
	}

	// DrawContext 上下文
	DrawContext struct {
		GraphicContext *draw2dimg.GraphicContext
		Image          image.Image
		Width          int
		Height         int
		CurrentLeft    int
		CurrentTop     int
	}
)

type componentItem struct {
	name    string
	creator func() Component
}

var (
	componentMu sync.Mutex
	components  atomic.Value
)

// RegisterComponent 注册组件类型
// name 组件名称
// c 组件创建器
func RegisterComponent(name string, c func() Component) {
	componentMu.Lock()
	cl, _ := components.Load().([]componentItem)
	components.Store(append(cl, componentItem{name, c}))
	componentMu.Unlock()
}

// 获取组建对象
func sniff(name string) componentItem {
	cl, _ := components.Load().([]componentItem)
	for _, c := range cl {
		if name == c.name {
			return c
		}
	}
	return componentItem{}
}

func offsetRect(rect *image.Rectangle, x int, y int) {
	rect.Min.X = rect.Min.X + x
	rect.Max.X = rect.Max.X + x
	rect.Min.Y = rect.Min.Y + y
	rect.Max.Y = rect.Max.Y + y
}

// DrawComponent 绘制组件
func (c *DrawContext) DrawComponent(component ComponentDefine) error {
	var err error = nil
	componentDrawer := sniff(component.Type)
	if componentDrawer.name != component.Type {
		return fmt.Errorf("no support component type %s", component.Type)
	}

	drawer := componentDrawer.creator()
	cConfig := drawer.ConfigType()
	cBytes, _ := json.Marshal(component.ComponentData)
	json.Unmarshal(cBytes, cConfig)

	//parentRect := image.Rect(0, 0, c.Width, c.Height)

	// 组件位置测量
	rect, mRect, err := c.measureComponent(drawer, component, cConfig)
	if err != nil {
		return err
	}

	// autoX, autoY := false, false
	// if rect.Min.X < 0 {
	// 	rect.Min.X = c.CurrentLeft
	// 	autoX = true
	// }
	// if rect.Min.Y < 0 {
	// 	rect.Min.Y = c.CurrentTop
	// 	autoY = true
	// }

	pdt, pdr, pdb, pdl, err := component.Padding.Parse(c.Width, c.Height)
	if err != nil {
		return err
	}

	// 测量组件尺寸
	// mRect := image.Rect(0, 0, rect.Dx(), rect.Dy())
	// if rect.Max.X < 0 {
	// 	mRect.Max.X = c.Width - c.CurrentLeft
	// }
	// if rect.Max.Y < 0 {
	// 	mRect.Max.Y = c.Height - c.CurrentTop
	// }
	// mRect, err = drawer.Measure(mRect, cConfig)
	// if err != nil {
	// 	return fmt.Errorf("component measure fail: %s", err.Error())
	// }

	// if rect.Max.X < 0 {
	// 	rect.Max.X = rect.Min.X + mRect.Dx() + pdr + pdl
	// }
	// if rect.Max.Y < 0 {
	// 	rect.Max.Y = rect.Min.Y + mRect.Dy() + pdt + pdb
	// }

	// // 需要由下级组件测量尺寸
	// if rect.Max.X < 0 || rect.Max.Y < 0 {
	// 	tmpRect := image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
	// 	actualRect, err := drawer.Measure(tmpRect, component)
	// 	if err != nil {
	// 		return fmt.Errorf("component measure fail: %s", err.Error())
	// 	}
	// 	rect = actualRect
	// 	rect.Max.X = rect.Max.X + pdr + pdl
	// 	rect.Max.Y = rect.Max.Y + pdt + pdb
	// }

	// 绘制组件图像
	cRect := image.Rect(0, 0, mRect.Dx(), mRect.Dy())
	// cImg := image.NewRGBA(cRect)
	cImg := image.NewRGBA(cRect)

	bgColor, err := component.BackgroundColor.Parse()
	if err != nil {
		return err
	}

	cgc := draw2dimg.NewGraphicContext(cImg)
	cgc.Save()
	cgc.SetFillColor(bgColor)
	cgc.Clear()
	cgc.Restore()

	// 绘制组件
	// 实际绘制区域（扣除边距）
	cDrawRect := image.Rect(rect.Min.X+pdl, rect.Min.Y+pdt, rect.Max.X-pdr, rect.Max.Y-pdb)

	// if cDrawRect.Max.X > c.Width {
	// 	if autoX && autoY {
	// 		offsetRect(&cDrawRect, -cDrawRect.Min.X, cDrawRect.Dy())
	// 		offsetRect(&rect, -rect.Min.X, rect.Dy())
	// 	} else if autoY {
	// 		offsetRect(&cDrawRect, 0, cDrawRect.Dy())
	// 		offsetRect(&rect, 0, rect.Dy())
	// 	}

	// }

	cContext := &DrawContext{GraphicContext: cgc, Image: cImg, Width: cDrawRect.Dx(), Height: cDrawRect.Dy(), CurrentLeft: 0, CurrentTop: 0}
	err = drawer.Draw(cContext, cConfig)
	if err != nil {
		return err
	}

	// 组件实际尺寸
	aRect, err := component.Size.Parse(cDrawRect, mRect)
	if err != nil {
		return err
	}

	// 外层画布（不含间距）
	bgRect := image.Rect(0, 0, cDrawRect.Dx(), cDrawRect.Dy())
	bgImg := image.NewRGBA(bgRect)
	isNewImage := true

	// 组件绘制区域
	x, y, err := component.Position.Parse(bgRect, aRect)
	if err != nil {
		return err
	}

	// 缩放
	sx, sy := float64(aRect.Dx())/float64(mRect.Dx()), float64(aRect.Dy())/float64(mRect.Dy())

	// 将组件绘制到上层
	if component.Repeat == RepeatNO {
		m := draw2d.NewTranslationMatrix(float64(x), float64(y))
		m.Scale(sx, sy)
		draw2dimg.DrawImage(cImg, bgImg, m, draw.Over, draw2dimg.LinearFilter)
	} else {
		l, t := 0, 0
		ml, mt := bgRect.Max.X, bgRect.Max.Y
		if component.Repeat == RepeatX {
			t = int(y)
			mt = t + 1
		}
		if component.Repeat == RepeatY {
			l = int(x)
			ml = l + 1
		}

		for i := t; i < mt; {

			for j := l; j < ml; {
				m := draw2d.NewTranslationMatrix(float64(j), float64(i))
				m.Scale(sx, sy)

				draw2dimg.DrawImage(cImg, bgImg, m, draw.Over, draw2dimg.LinearFilter)
				j += aRect.Dx()
			}
			i += aRect.Dy()
		}

	}

	if isNewImage {
		c.GraphicContext.Save()
		c.GraphicContext.SetFillColor(bgColor)
		c.GraphicContext.ClearRect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
		c.GraphicContext.Restore()
		draw2dimg.DrawImage(bgImg, c.Image.(draw.Image), draw2d.NewTranslationMatrix(float64(pdl+rect.Min.X), float64(pdt+rect.Min.Y)), draw.Over, draw2dimg.LinearFilter)
	}

	c.CurrentLeft = rect.Max.X
	c.CurrentTop = rect.Min.Y

	return nil
}

func (c *DrawContext) measureComponent(cp Component, cd ComponentDefine, config interface{}) (image.Rectangle, image.Rectangle, error) {
	var err error = nil

	mr := image.Rect(0, 0, 0, 0)

	parentRect := image.Rect(0, 0, c.Width, c.Height)
	// 组件在父级的位置
	rect, err := cd.Area.Parse(parentRect)
	if err != nil {
		return mr, mr, err
	}

	autoX, autoY := false, false
	if rect.Min.X < 0 {
		rect.Min.X = c.CurrentLeft
		autoX = true
	}
	if rect.Min.Y < 0 {
		rect.Min.Y = c.CurrentTop
		autoY = true
	}

	pdt, pdr, pdb, pdl, err := cd.Padding.Parse(c.Width, c.Height)
	if err != nil {
		return mr, mr, err
	}

	cConfig := cp.ConfigType()
	cBytes, _ := json.Marshal(cd.ComponentData)
	json.Unmarshal(cBytes, cConfig)

	// 测量组件尺寸
	mRect := image.Rect(0, 0, 0, 0)
	if rect.Max.X < 0 {
		mRect.Max.X = c.Width - c.CurrentLeft
	} else {
		mRect.Max.X = rect.Dx()
	}
	if rect.Max.Y < 0 {
		mRect.Max.Y = c.Height - c.CurrentTop
	} else {
		mRect.Max.Y = rect.Dy()
	}

	mRect, err = cp.Measure(c.Clone(), mRect, cConfig)
	if err != nil {
		return mr, mr, fmt.Errorf("component measure fail: %s", err.Error())
	}

	if rect.Max.X < 0 {
		rect.Max.X = rect.Min.X + mRect.Dx() + pdr + pdl
	}
	if rect.Max.Y < 0 {
		rect.Max.Y = rect.Min.Y + mRect.Dy() + pdt + pdb
	}

	if autoX && autoY && rect.Max.X > c.Width {
		breakLine := false
		if rect.Dx() <= c.Width {
			breakLine = true
		} else if rect.Min.X >= ((c.Width * 2) / 3) {
			// x 坐标超过 2/3时需换行
			breakLine = true
		}

		if breakLine {
			offsetRect(&rect, -rect.Min.X, rect.Dy())
		}
	}

	return rect, mRect, nil
}

// Clone 创建DrawContext副本
func (c *DrawContext) Clone() *DrawContext {
	return &DrawContext{
		GraphicContext: c.GraphicContext,
		Width:          c.Width,
		Height:         c.Height,
		Image:          c.Image,
		CurrentLeft:    c.CurrentLeft,
		CurrentTop:     c.CurrentTop,
	}
}

// MeasureComponent 组件位置测量
func (c *DrawContext) MeasureComponent(component ComponentDefine) (image.Rectangle, image.Rectangle, error) {

	componentDrawer := sniff(component.Type)
	if componentDrawer.name != component.Type {
		return image.Rect(0, 0, 0, 0), image.Rect(0, 0, 0, 0), fmt.Errorf("no support component type %s", component.Type)
	}

	drawer := componentDrawer.creator()
	cConfig := drawer.ConfigType()
	cBytes, _ := json.Marshal(component.ComponentData)
	json.Unmarshal(cBytes, cConfig)

	// 组件位置测量
	return c.measureComponent(drawer, component, cConfig)
}
