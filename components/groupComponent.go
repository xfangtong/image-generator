package components

import (
	"image"
	"image/draw"
	"math"
	"sort"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
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
	di   *drawItem
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

	levelMap := make(map[int][]*componentDefineWrapper)
	levelList := make([]int, 0)
	clTemp := make([]*componentDefineWrapper, 0)
	for i, c := range cs {
		wItem := &componentDefineWrapper{Item: c, Seq: i}
		clTemp = append(clTemp, wItem)
		m, b := levelMap[c.Level]
		if !b {
			m = make([]*componentDefineWrapper, 0)
			levelList = append(levelList, c.Level)
		}
		levelMap[c.Level] = append(m, wItem)
	}

	for _, c := range clTemp {
		di, err := dc.createComponentImage(c.Item)
		if err != nil {
			return err
		}
		c.di = &di
	}

	sort.Ints(levelList)

	for _, l := range levelList {
		items := levelMap[l]
		for _, ci := range items {
			bgColor, err := ci.Item.BackgroundColor.Parse()
			if err != nil {
				return err
			}
			di := ci.di
			br, dr := di.bgRect, di.drawRect
			dc.GraphicContext.Save()
			dc.GraphicContext.SetFillColor(bgColor)
			dc.GraphicContext.ClearRect(br.Min.X, br.Min.Y, br.Max.X, br.Max.Y)
			draw2dimg.DrawImage(di.image, dc.Image.(draw.Image), draw2d.NewTranslationMatrix(float64(dr.Min.X), float64(dr.Min.Y)), draw.Over, draw2dimg.LinearFilter)
			dc.GraphicContext.Restore()
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
