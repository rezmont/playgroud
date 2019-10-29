package structs

import "math"

type rectangle struct {
	w float32
	h float32
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func perimeter(r rectangle) float32 {
	return (r.w + r.h) * 2
}

func (r rectangle) area() float64 {
	return (float64)(r.w * r.h)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
