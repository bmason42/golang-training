package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShapes(t *testing.T) {
	circle := Circle{center: Point{1, 2}, radius: 4}
	rect := NewRectangle(0, 0, 100, 100)
	expectedArea := PI * 4 * 4
	testArea(t, expectedArea, &circle)
	expectedArea = 100 * 100
	testArea(t, expectedArea, rect)
	expectedArea = 100 * 100
	testAreaNeg(t, 2, rect)
}

func testArea(t *testing.T, expected float64, shape Shape) {
	calculatedArea := shape.CalArea()
	assert.Equal(t, expected, calculatedArea, "Shape area did not match")
}
func testAreaNeg(t *testing.T, expected float64, shape Shape) {
	calculatedArea := shape.CalArea()
	assert.NotEqualf(t, expected, calculatedArea, "Shape area did not match")
}
