package imagegenerator

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/xfangtong/image-generator/components"
)

func TestBackgroundRepeat(t *testing.T) {
	tmp := ImageTemplate{
		Width:           "400",
		Height:          "500",
		BackgroundColor: "#ff0000",
		Background:      "local://./images/dog.png",
		Repeat:          components.RepeatXY,
		Size:            "100% 100%",
		Padding:         "0",
		Position:        "left top",
	}

	img, _ := GenerateImage(tmp)

	f, _ := os.Create("./test/img_template_background_repeat.png")
	defer f.Close()
	png.Encode(f, img)
}

func TestPoster(t *testing.T) {
	avatar := components.ComponentDefine{
		Type:  "avatar",
		Level: 1,
		Area: components.Rectangle{
			Left:   "0",
			Top:    "1280",
			Right:  "150",
			Bottom: "1440",
		},
		Position: "center",
		Size:     "100% 100%",
		ComponentData: map[string]interface{}{
			"url":   "local://./images/avatar.jpg",
			"width": 100,
			"size":  "contain",
		},
	}
	name := components.ComponentDefine{
		Type:  "text",
		Level: 1,
		Area: components.Rectangle{
			Left:   "150",
			Top:    "1280",
			Right:  "auto",
			Bottom: "1440",
		},
		Position: "center",
		Size:     "100% 100%",
		ComponentData: map[string]interface{}{
			"text":      "张三 13100000000",
			"fontPath":  "./fonts/zkgdh.ttf",
			"fontSize":  38,
			"fillColor": "#03732A",
		},
	}
	qrcode := components.ComponentDefine{
		Type:  "qrcode",
		Level: 1,
		Area: components.Rectangle{
			Left:   "520",
			Top:    "1280",
			Right:  "720",
			Bottom: "1440",
		},
		Position: "center",
		Size:     "100% 100%",
		ComponentData: map[string]interface{}{
			"content": "张三 13100000000",
			"color":   "#03732A",
			"width":   120,
		},
	}
	tmp := ImageTemplate{
		Width:           "auto",
		Height:          "+160",
		BackgroundColor: "#ffffff",
		Background:      "local://./images/dw.jpg",
		Repeat:          components.RepeatNO,
		Size:            "100% 100%",
		Padding:         "0",
		Position:        "left top",
		Components: []components.ComponentDefine{
			avatar, name, qrcode,
		},
	}

	img, _ := GenerateImage(tmp)

	f, _ := os.Create("./test/img_template_poster.png")
	defer f.Close()
	png.Encode(f, img)
}

func TestTemplatePoster(t *testing.T) {

	jsonBytes, _ := ioutil.ReadFile("./examples/poster.json")
	jsonText := string(jsonBytes)

	img, _ := GenerateImageFromTemplate(jsonText, map[string]string{
		"url":    "local://./images/dw.jpg",
		"avatar": "local://./images/avatar.jpg",
		"name":   "王五",
		"tel":    "13098899899",
	})

	f, _ := os.Create("./test/img_template_poster2.png")
	defer f.Close()
	png.Encode(f, img)
}
