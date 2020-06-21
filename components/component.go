package components

import (
	"encoding/json"
	"fmt"
	"image"
	"sync"
	"sync/atomic"

	"github.com/fogleman/gg"
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
		BackgroundColor Gradient `json:"backgroundColor"`
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
		//GraphicContext *draw2dimg.GraphicContext
		GraphicContext *gg.Context
		Image          image.Image
		Width          int
		Height         int
		CurrentLeft    int
		CurrentTop     int
		AutoWidth      bool
		AutoHeight     bool
	}
)

type (
	componentItem struct {
		name    string
		creator func() Component
	}
	drawItem struct {
		image    image.Image
		drawRect image.Rectangle
		bgRect   image.Rectangle
	}
)

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

func (c *DrawContext) createComponentImage(component ComponentDefine) (drawItem, error) {
	var err error = nil
	componentDrawer := sniff(component.Type)
	if componentDrawer.name != component.Type {
		return drawItem{}, fmt.Errorf("no support component type %s", component.Type)
	}

	drawer := componentDrawer.creator()
	cConfig := drawer.ConfigType()
	cBytes, _ := json.Marshal(component.ComponentData)
	json.Unmarshal(cBytes, cConfig)

	// 组件位置测量
	rect, mRect, err := c.measureComponent(drawer, component, cConfig)
	if err != nil {
		return drawItem{}, err
	}

	pdt, pdr, pdb, pdl, err := component.Padding.Parse(c.Width, c.Height)
	if err != nil {
		return drawItem{}, err
	}

	// 绘制组件图像
	cRect := image.Rect(0, 0, mRect.Dx(), mRect.Dy())
	// cImg := image.NewRGBA(cRect)
	cImg := image.NewRGBA(cRect)

	// bgColor, err := component.BackgroundColor.Parse()
	// if err != nil {
	// 	return drawItem{}, err
	// }

	cgc := gg.NewContextForRGBA(cImg)
	cgc.Push()
	err = component.BackgroundColor.SetFill(cgc, float64(mRect.Dx()), float64(mRect.Dy()))
	cgc.DrawRectangle(0, 0, float64(mRect.Dx()), float64(mRect.Dy()))
	cgc.Fill()
	//cgc.SetColor(bgColor)
	//cgc.Clear()
	cgc.Pop()

	// 绘制组件
	// 实际绘制区域（扣除边距）
	cDrawRect := image.Rect(rect.Min.X+pdl, rect.Min.Y+pdt, rect.Max.X-pdr, rect.Max.Y-pdb)

	cContext := &DrawContext{GraphicContext: cgc, Image: cImg, Width: cDrawRect.Dx(), Height: cDrawRect.Dy(), CurrentLeft: 0, CurrentTop: 0}

	err = drawer.Draw(cContext, cConfig)
	if err != nil {
		return drawItem{}, err
	}
	//cgc = cContext.GraphicContext

	// 组件实际尺寸
	aRect, err := component.Size.Parse(cDrawRect, mRect)
	if err != nil {
		return drawItem{}, err
	}

	// 外层画布（不含间距）
	bgRect := image.Rect(0, 0, cDrawRect.Dx(), cDrawRect.Dy())
	bgImg := image.NewRGBA(bgRect)
	bggc := gg.NewContextForRGBA(bgImg)

	// 组件绘制区域
	x, y, err := component.Position.Parse(bgRect, aRect)
	if err != nil {
		return drawItem{}, err
	}

	// 缩放
	sx, sy := float64(aRect.Dx())/float64(mRect.Dx()), float64(aRect.Dy())/float64(mRect.Dy())

	// 将组件绘制到上层
	if component.Repeat == RepeatNO {
		bggc.Translate(float64(x), float64(y))
		bggc.Scale(sx, sy)
		bggc.DrawImage(cImg, 0, 0)

		//m := draw2d.NewTranslationMatrix(float64(x), float64(y))
		//m.Scale(sx, sy)

		//draw2dimg.DrawImage(cImg, bgImg, m, draw.Over, draw2dimg.LinearFilter)
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
				bggc.Push()
				bggc.Translate(float64(j), float64(i))
				bggc.Scale(sx, sy)
				bggc.DrawImage(cImg, 0, 0)
				// m := draw2d.NewTranslationMatrix(float64(j), float64(i))
				// m.Scale(sx, sy)

				// draw2dimg.DrawImage(cImg, bgImg, m, draw.Over, draw2dimg.LinearFilter)
				j += aRect.Dx()
				bggc.Pop()
			}
			i += aRect.Dy()
		}

	}

	c.CurrentLeft = rect.Max.X
	c.CurrentTop = rect.Min.Y

	return drawItem{image: bgImg, drawRect: cDrawRect, bgRect: rect}, nil

}

// DrawComponent 绘制组件
func (c *DrawContext) DrawComponent(component ComponentDefine) error {

	di, err := c.createComponentImage(component)
	if err != nil {
		return err
	}

	//bgColor := component.BackgroundColor.ParseNoError()

	c.GraphicContext.Push()
	defer c.GraphicContext.Pop()
	err = component.BackgroundColor.SetFill(c.GraphicContext, float64(di.bgRect.Dx()), float64(di.bgRect.Dy()))
	if err != nil {
		return err
	}
	//c.GraphicContext.SetColor(bgColor)
	c.GraphicContext.DrawRectangle(float64(di.bgRect.Min.X), float64(di.bgRect.Min.Y), float64(di.bgRect.Dx()), float64(di.bgRect.Dy()))
	//c.GraphicContext.SetColor(bgColor)
	c.GraphicContext.Fill()
	c.GraphicContext.Translate(float64(di.drawRect.Min.X), float64(di.drawRect.Min.Y))
	c.GraphicContext.DrawImage(di.image, 0, 0)
	//c.GraphicContext.Pop()
	c.Image = c.GraphicContext.Image()

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
	if rect.Min.X == AutoValue {
		rect.Min.X = c.CurrentLeft
		autoX = true
	}
	if rect.Min.Y == AutoValue {
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
	if rect.Max.X == AutoValue {
		mRect.Max.X = c.Width - c.CurrentLeft
	} else {
		mRect.Max.X = rect.Dx()
	}
	if rect.Max.Y == AutoValue {
		mRect.Max.Y = c.Height - c.CurrentTop
	} else {
		mRect.Max.Y = rect.Dy()
	}

	nc := c.Clone()
	nc.AutoHeight = autoY
	nc.AutoWidth = autoX

	mRect, err = cp.Measure(nc, mRect, cConfig)
	if err != nil {
		return mr, mr, fmt.Errorf("component measure fail: %s", err.Error())
	}

	if rect.Max.X == AutoValue {
		rect.Max.X = rect.Min.X + mRect.Dx() + pdr + pdl
	}
	if rect.Max.Y == AutoValue {
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
