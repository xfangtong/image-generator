package components

import (
	"fmt"
	"testing"
)

func TestGradientParse1(t *testing.T) {
	g1 := "linear-gradient(100 100, rgba(2,0,36,1) 0%, rgba(9,9,121,1) 41%, rgba(8,40,141,1) 55%, rgba(0,212,255,1) 100%)"

	gg, _ := _parseLinearGradient(g1, 100.0, 100.0)

	fmt.Println(gg)
}
