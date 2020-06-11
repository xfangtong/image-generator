package imagegenerator

import (
	"sync"
	"sync/atomic"
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
		// ComponentData 组件数据，根据组件类型具有不同的结构
		ComponentData map[string]interface{} `json:"componentData"`
	}
)

type componentItem struct {
	name string
	draw func() error
}

var (
	componentMu sync.Mutex
	components  atomic.Value
)

// RegisterComponent 注册组件类型
// name 组件名称
// draw 绘制方法
func RegisterComponent(name string, draw func() error) {
	componentMu.Lock()
	cl, _ := components.Load().([]componentItem)
	components.Store(append(cl, componentItem{name, draw}))
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
