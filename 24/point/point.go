package point

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point { // конструктор
	return &Point{x, y}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}
