package imagegenerator

type (
	// ImageTemplate 图片模板
	ImageTemplate struct {
		// Background 背景
		Background Resource `json:"background"`
		// BackgroundColor 背景颜色
		BackgroundColor Color `json:"backgroundColor"`
		// Width 宽度，如果设置为0，使用背景宽度
		Width int `json:"width"`
		// Height 高度，如果设置为0，使用背景高度
		Height int `json:"height"`
		// Component 组件列表
		Components []ComponentDefine `json:"components"`
	}
)
