package components

import (
	"fmt"
	"image"
	"math"
	"strings"
	"unicode"

	"github.com/fogleman/gg"
)

type (
	// TextComponentDefine 文本组件
	TextComponentDefine struct {
		ShapeComponentDefine
		Text       string  `json:"text"`
		FontPath   string  `json:"fontPath"`
		FontSize   float64 `json:"fontSize"`
		LineHeight float64 `json:"lineHeight"`
	}

	// TextComponent 文本组件
	TextComponent struct {
		_drawItmes [][]drawTextItem
		_w, _h     float64
	}

	drawTextItem struct {
		x, y, w, h, b float64
		text          string
	}
)

func splitText(s string) [][]rune {
	t := strings.Replace(s, "\r\n", "\n", -1)
	t = strings.Replace(t, "\r", "\n", -1)
	ts := []rune(t)

	text := make([][]rune, 0)
	word := ""
	for _, tx := range ts {
		if unicode.IsSpace(tx) || tx < 0 || tx > 256 {
			if word != "" {
				text = append(text, []rune(word))
			}
			// if tx == '\t' {
			// 	text = append(text, []rune("    "))
			// } else {
			text = append(text, []rune{tx})
			//}

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

func measureText(gc *gg.Context, text [][]rune, w float64, lineHeight float64) [][]drawTextItem {
	x, y := 0.0, 0.0
	dl := make([]drawTextItem, 0)

	_, norHeight := gc.MeasureString("国")

	for _, rs := range text {
		if len(rs) == 1 && rs[0] == '\n' {
			// 直接换行
			dl = append(dl, drawTextItem{
				x:    0,
				y:    y,
				w:    0,
				h:    norHeight,
				text: "",
			})
			y++
			x = 0
			continue
		}
		tw, th := gc.MeasureString(string(rs))

		if (x + tw) > w {
			//换行
			if tw > w {
				//强拆
				sr := 0
				for i := sr; i < len(rs); i++ {
					s := string(rs[sr : i+1])
					sw, sh := gc.MeasureString(s)
					if (x + sw) > w {
						s = string(rs[sr:i])
						sw, sh = gc.MeasureString(s)
						sr = i
						dl = append(dl, drawTextItem{
							x:    x,
							y:    y,
							w:    sw,
							h:    sh,
							text: s,
						})
						y++
						x = 0
					}

					if i == (len(rs) - 1) {
						dl = append(dl, drawTextItem{
							x:    x,
							y:    y,
							w:    sw,
							h:    sh,
							text: s,
						})
						x += sw
					}
				}

			} else {
				y++
				x = 0
				dl = append(dl, drawTextItem{
					x:    x,
					y:    y,
					w:    tw,
					h:    th,
					text: string(rs),
				})
				x += tw
			}
		} else {

			dl = append(dl, drawTextItem{
				x:    x,
				y:    y,
				w:    tw,
				h:    th,
				text: string(rs),
			})
			x += tw
		}

	}

	lines := make([][]drawTextItem, 0)

	// lh := 0.0
	y = 0
	top := 0.0
	si := 0

	processLine := func(line []drawTextItem, top float64) ([]drawTextItem, float64) {
		lh := 0.0
		for _, di := range line {
			lh = math.Max(lh, di.h)
		}
		// 行高
		alh := lh
		if lineHeight > alh {
			alh = lineHeight
		}

		newItems := make([]drawTextItem, 0)
		for _, item := range line {
			item.h = alh
			item.y = (top + alh) - (alh-lh)/2.0
			item.b = top + alh
			newItems = append(newItems, item)
		}

		return newItems, top + alh

	}

	for idx, di := range dl {
		if di.y != y {
			items := dl[si:idx]
			var newItems []drawTextItem
			newItems, top = processLine(items, top)
			lines = append(lines, newItems)
			y = di.y
			si = idx
		}
		if idx == len(dl)-1 {
			items := dl[si : idx+1]
			var newItems []drawTextItem
			newItems, top = processLine(items, top)
			lines = append(lines, newItems)
		}

	}

	return lines
}

// Draw 绘制
func (c *TextComponent) Draw(dc *DrawContext, config interface{}) error {
	cd := config.(*TextComponentDefine)

	var err error = nil

	dc.GraphicContext.Push()
	defer dc.GraphicContext.Pop()
	if err = dc.GraphicContext.LoadFontFace(cd.FontPath, cd.FontSize); err != nil {
		return err
	}

	if fc, err := cd.FillColor.Parse(); err == nil {
		dc.GraphicContext.SetColor(fc)
	} else {
		return err
	}

	lines := c._drawItmes
	for _, l := range lines {
		for _, t := range l {
			dc.GraphicContext.DrawString(t.text, t.x, t.y)
		}
	}

	return err
}

// Measure 测量
func (c *TextComponent) Measure(dc *DrawContext, rect image.Rectangle, config interface{}) (image.Rectangle, error) {
	cd := config.(*TextComponentDefine)

	var err error = nil
	rt := image.Rect(0, 0, 0, 0)

	w := rect.Dx()
	if w <= 0 {
		return rt, fmt.Errorf("必须确定宽度")
	}

	if err = dc.GraphicContext.LoadFontFace(cd.FontPath, cd.FontSize); err != nil {
		return rt, err
	}

	texts := splitText(cd.Text)
	lines := measureText(dc.GraphicContext, texts, float64(w), cd.LineHeight)
	c._drawItmes = lines

	mw, mh := 0.0, 0.0
	for _, l := range lines {

		for _, t := range l {
			mw = math.Max(mw, t.x+t.w)
			mh = math.Max(mh, t.b)
		}
	}
	c._w = mw
	c._h = mh
	return image.Rect(0, 0, int(mw), int(mh)), nil
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
