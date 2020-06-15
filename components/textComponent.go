package components

import (
	"image"
	"strings"
	"unicode"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type (
	// TextComponentDefine 文本组件
	TextComponentDefine struct {
		ShapeComponentDefine
		Text     string
		FontSize int
	}

	// TextComponent 文本组件
	TextComponent struct {
	}
)

func splitText(s string) [][]rune {
	t := strings.Replace(s, "\r\n", "\n", -1)
	ts := []rune(t)

	text := make([][]rune, 0)
	word := ""
	for _, tx := range ts {
		if unicode.IsSpace(tx) || tx < 0 || tx > 256 {
			if word != "" {
				text = append(text, []rune(word))
			}
			if tx == '\t' {
				text = append(text, []rune("    "))
			} else {
				text = append(text, []rune{tx})
			}

			word = ""
		} else {
			word = word + string([]rune{tx})
		}
	}

	if word != "" {
		text = append(text, []rune(word))
	}

	return text
}

func measureText(gc *draw2dimg.GraphicContext, text [][]rune, w float64) {
	y := 0.0
	for _, rs := range text {
		_, _, r, b := gc.GetStringBounds(string(rs))
		si := 0
		if r > w {
			for idx := range rs {
				_, _, r, b = gc.GetStringBounds(string(rs[si : idx+1]))
				if r > w {
					y = b
					si = idx
					gc.MoveTo(0.0, y)
				}
			}
		}

	}
}

// Draw 绘制
func (c *TextComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*TextComponentDefine)

	var err error = nil
	fc, err := cd.ShapeComponentDefine.FillColor.Parse()
	if err != nil {
		return err
	}
	sc, err := cd.ShapeComponentDefine.StrokeColor.Parse()
	if err != nil {
		return err
	}

	lw := float64(cd.ShapeComponentDefine.LineWidth)
	gc := dc.GraphicContext
	gc.Save()
	gc.SetLineWidth(lw)
	gc.SetFillColor(fc)
	gc.SetStrokeColor(sc)
	gc.SetFillRule(draw2d.FillRuleWinding)
	gc.BeginPath()

	//x := float64(cd.Radius)
	//gc.ArcTo(float64(x), float64(x), float64(cd.Radius)-lw/2, float64(cd.Radius)-lw/2, 0, 2*math.Pi)
	gc.FillStroke()

	gc.Restore()

	return nil
}

// Measure 测量
func (c *TextComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	td := config.(*TextComponentDefine)
	dc.GraphicContext.Save()
	dc.GraphicContext.SetFontSize(float64(td.FontSize))
	dc.GraphicContext.MoveTo(0, 0)
	// w, h := rect.Dx(), rect.Dy()
	//w := float64(rect.Dx())
	//y := 0.0

	//text := splitText(td.Text)

	//dc.GraphicContext.getstr

	return image.Rect(0, 0, 0, 0), nil
}

// ConfigType 配置类型
func (c *TextComponent) ConfigType() interface{} {
	return &TextComponentDefine{}
}

func init() {
	RegisterComponent("text", func() Component {
		return &TextComponent{}
	})
}
