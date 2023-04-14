package pkg

const PI = float64(3.1415926536)

type Point struct {
	x int
	y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) SetY(y int) {
	p.y = y
}

type Circle struct {
	center Point
	radius int
}

func (t *Circle) CalArea() float64 {
	return PI * float64(t.radius*t.radius)
}

type Rectangle struct {
	leftUpper  Point
	rightLower Point
}

func NewRectangle(x1 int, y1 int, x2 int, y2 int) *Rectangle {
	ret := new(Rectangle)
	ret.leftUpper.SetX(x1)
	ret.leftUpper.SetY(y1)
	ret.rightLower.SetX(x2)
	ret.rightLower.SetY(y2)
	return ret
}
func (t *Rectangle) CalArea() float64 {
	width := t.rightLower.X() - t.leftUpper.X()
	height := t.rightLower.Y() - t.leftUpper.Y()
	return float64(width * height)
}

type Shape interface {
	CalArea() float64
}
