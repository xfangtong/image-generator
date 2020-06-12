package imagegenerator

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/llgcode/draw2d/draw2dimg"
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
		// ComponentData 组件数据，根据组件类型具有不同的结构
		ComponentData map[string]interface{} `json:"componentData"`
	}

	// Component 组件接口
	Component interface {
		Draw(c *DrawContext, define ComponentDefine) error
		Measure(rect image.Rectangle, define ComponentDefine) (image.Rectangle, error)
	}
)

type componentItem struct {
	name      string
	component Component
}

var (
	componentMu sync.Mutex
	components  atomic.Value
)

// RegisterComponent 注册组件类型
// name 组件名称
// draw 绘制方法
func RegisterComponent(name string, c Component) {
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

// DrawComponent 绘制组件
func (c *DrawContext) DrawComponent(component ComponentDefine) error {
	var err error = nil
	componentDrawer := sniff(component.Type)
	if componentDrawer.name != component.Type {
		return fmt.Errorf("no support component type %s", component.Type)
	}

	ct := reflect.TypeOf(componentDrawer)
	if ct.Kind() == reflect.Ptr {
		ct = ct.Elem()
	}
	drawer := reflect.New(ct).Interface().(Component)
	if drawer == nil {
		return fmt.Errorf("创建组件实例失败")
	}

	parentRect := image.Rect(0, 0, c.Width, c.Height)
	// 组件在父级的位置
	rect, err := component.Area.Parse(parentRect)
	if err != nil {
		return err
	}

	if rect.Min.X < 0 {
		rect.Min.X = c.CurrentLeft
	}
	if rect.Min.Y < 0 {
		rect.Min.Y = c.CurrentTop
	}

	// 需要由下级组件测量尺寸
	if rect.Max.X < 0 || rect.Max.Y < 0 {
		tmpRect := image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
		actualRect, err := drawer.Measure(tmpRect, component)
		if err != nil {
			return fmt.Errorf("component measure fail: %s", err.Error())
		}
		rect = actualRect
	}

	// 绘制组件图像
	cRect := image.Rect(0, 0, rect.Dx(), rect.Dy())
	cImg := image.NewRGBA(cRect)
	cgc := draw2dimg.NewGraphicContext(cImg)
	cgc.Save()
	cgc.SetFillColor(color.Transparent)
	cgc.Clear()
	cgc.Restore()

	cContext := &DrawContext{GraphicContext: cgc, Width: cRect.Dx(), Height: cRect.Dy(), CurrentLeft: 0, CurrentTop: 0}
	err = drawer.Draw(cContext, component)
	if err != nil {
		return err
	}

	return nil
}
