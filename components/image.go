package components

import (
	"github.com/llgcode/draw2d/draw2dimg"
	imagegenerator "github.com/xfangtong/image-generator"
)

type (
	// ImageComponentDefine 图像组件定义
	ImageComponentDefine struct {
		// URL 图像地址
		URL imagegenerator.Resource `json:"url"`
	}
)

func drawImageComponent(gc *draw2dimg.GraphicContext, c interface{}, context *imagegenerator.DrawContext) error {
	return nil
}

func init() {

}
