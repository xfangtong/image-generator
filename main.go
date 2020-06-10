package main

// type (
// 	TestInner struct {
// 		Text     string `json:"text"`
// 		FontSize int
// 	}
// 	TestConfig struct {
// 		Left  int
// 		Top   int
// 		Color string
// 		Data  TestInner
// 	}
// 	TestEntity struct {
// 		Level     int        `json:"level"`
// 		Component string     `json:"component"`
// 		Config    TestConfig `json:"config"`
// 	}

// 	TestComponent struct {
// 		Level     int                    `json:"level"`
// 		Component string                 `json:"component"`
// 		Config    map[string]interface{} `json:"config"`
// 	}
// )

func main() {
	// te := TestEntity{
	// 	Level:     1,
	// 	Component: "Text",
	// 	Config: TestConfig{
	// 		Left:  0,
	// 		Top:   0,
	// 		Color: "#FFF",
	// 		Data: TestInner{
	// 			Text:     "Hello",
	// 			FontSize: 12,
	// 		},
	// 	},
	// }

	// bs, _ := json.Marshal(te)
	// //	jsonText := string(bytes)

	// tc := TestComponent{}
	// json.Unmarshal(bs, &tc)

	// bs, _ = json.Marshal(tc)
	// jsonText := string(bs)

	// json2Bytes, _ := json.Marshal(tc.Config)

	// tec := TestConfig{}
	// json.Unmarshal(json2Bytes, &tec)

	// //log.Infof(nil, "json %s", jsonText)
	// println(jsonText)

	// tmp, _ := template.New("imageg").Parse("hello, {{ .test }}")
	// w := bytes.NewBuffer(make([]byte, 0))
	// tmp.Execute(w, map[string]string{
	// 	"test": "zhangsn",
	// })

	// println(string(w.Bytes()))

	// c, _ := colors.ParseHEX("#ff00ffff")
	// rgba := c.ToRGBA()

	// c2 := color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: uint8(rgba.A * 255)}

	// println(fmt.Sprintf("%v", c2))
}
