package imagegenerator

type (
	// Rectangle 区域
	Rectangle struct {
		Left   Dimension `json:"left"`
		Top    Dimension `json:"top"`
		Right  Dimension `json:"right"`
		Bottom Dimension `json:"bottom"`
	}
)
