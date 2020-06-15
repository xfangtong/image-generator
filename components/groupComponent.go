package components

import (
	"image"
	"math"
	"sort"
)

type (
	// GroupComponentDefine 组合组件
	GroupComponentDefine struct {
		// Components 子组件
		Components []ComponentDefine `json:"components"`
	}

	// GroupComponent 组合组件
	GroupComponent struct {
	}
)

type componentDefineWrapper struct {
	Item ComponentDefine
	Seq  int
}

type componentList []componentDefineWrapper

func (cl componentList) Len() int {
	return len(cl)
}

func (cl componentList) Less(i, j int) bool {
	return cl[i].Item.Level > cl[j].Item.Level || (cl[i].Item.Level == cl[j].Item.Level && cl[i].Seq > cl[j].Seq)
}

func (cl componentList) Swap(i, j int) {
	cl[i], cl[j] = cl[j], cl[i]
}

// Draw 绘制
func (c *GroupComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*GroupComponentDefine)
	cs := cd.Components

	clTemp := make([]componentDefineWrapper, 0)
	for i, c := range cs {
		clTemp = append(clTemp, componentDefineWrapper{Item: c, Seq: i})
	}
	cl := componentList(clTemp)
	sort.Sort(cl)

	for _, c := range cl {
		err := dc.DrawComponent(c.Item)
		if err != nil {
			return err
		}

	}

	return nil
}

// Measure 测量
func (c *GroupComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {

	cd := config.(*GroupComponentDefine)
	w, h := 0.0, 0.0
	cs := cd.Components
	mr := image.Rect(0, 0, 0, 0)

	clTemp := make([]componentDefineWrapper, 0)
	for i, c := range cs {
		clTemp = append(clTemp, componentDefineWrapper{Item: c, Seq: i})
	}
	cl := componentList(clTemp)
	sort.Sort(cl)

	for _, c := range cl {
		r1, _, err := dc.MeasureComponent(c.Item)
		if err != nil {
			return mr, err
		}
		dc.CurrentLeft = r1.Max.X
		dc.CurrentTop = r1.Min.Y
		w = math.Max(w, float64(r1.Max.X))
		h = math.Max(h, float64(r1.Max.Y))
	}

	mr.Max.X = int(w)
	mr.Max.Y = int(h)

	return mr, nil
}

// ConfigType 配置类型
func (c *GroupComponent) ConfigType() interface{} {
	return &GroupComponentDefine{}
}

func init() {
	RegisterComponent("group", func() Component {
		return &GroupComponent{}
	})
}
