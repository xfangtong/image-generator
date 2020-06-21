package imagegenerator

import (
	components "github.com/xfangtong/image-generator/components"
	"github.com/xfangtong/image-generator/resources"
)

type (
	// ImageTemplate 图片模板
	ImageTemplate struct {
		// Position 在区域中的位置方式
		Position components.Position `json:"position"`
		// Repeat 背景重复方式
		Repeat string `json:"repeat"`
		// Size 背景尺寸方式
		Size components.Size `json:"size"`
		// BackgroundPadding 背景边距
		BackgroundPadding components.Padding `json:"backgroundPadding"`
		// Padding 内容边距
		Padding components.Padding `json:"padding"`
		// Background 背景图片
		Background resources.Resource `json:"background"`
		// BackgroundColor 背景颜色
		BackgroundColor components.Gradient `json:"backgroundColor"`
		// Width 宽度
		Width components.Dimension `json:"width"`
		// Height 高度
		Height components.Dimension `json:"height"`
		// Component 组件列表
		Components []components.ComponentDefine `json:"components"`
	}
)
