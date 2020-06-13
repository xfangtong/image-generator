package components

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMeasure(t *testing.T) {
	d := Dimension("10")
	assert.Equal(t, 10.0, d.MeasureNoError(0))

	d = Dimension("10.2")
	assert.Equal(t, 10.2, d.MeasureNoError(0))

	d = Dimension("1%")
	assert.Equal(t, 1.0, d.MeasureNoError(100))

}
