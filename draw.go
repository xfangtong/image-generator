package imagegenerator

type (
	// Rectangle 区域
	Rectangle struct {
		Left   Dimension `json:"left"`
		Top    Dimension `json:"top"`
		Right  Dimension `json:"right"`
		Bottom Dimension `json:"bottom"`
	}

	// Component 组件
	Component struct {
		// Type 类型
		Type string `json:"type"`
		// Level 层次
		Level int `json:"level"`
		// Area 放置区域
		Area Rectangle `json:"area"`
		// ComponentData 组件数据，根据组件类型具有不同的结构
		ComponentData map[string]interface{} `json:"componentData"`
	}

	// ImageTemplate 图片模板
	ImageTemplate struct {
		// Background 背景
		Background Resource `json:"background"`
		// BackgroundColor 背景颜色
		BackgroundColor Color `json:"backgroundColor"`
		// Components 组件列表
		Components []Component `json:"components"`
	}
)
