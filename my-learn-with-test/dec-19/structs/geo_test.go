package structs

import (
	"math"
	"testing"
)

func TestPerim(t *testing.T) {
	testCases := []struct {
		desc string
		r    rectangle
		want float32
	}{
		{
			desc: "",
			r:    rectangle{w: 10, h: 20},
			want: 60,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := perimeter(tC.r)
			if tC.want != got {
				t.Errorf("want %f, got %f for %v", tC.want, got, tC.r)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		desc string
		s    shape
		want float64
	}{
		{
			desc: "circle",
			s:    circle{10},
			want: 314.1592,
		},
		{
			desc: "rectangle",
			s:    rectangle{10, 20},
			want: 200,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.s.area()
			diff := math.Abs(tC.want - got)
			if diff > 0.0001 {
				t.Errorf("want %f, got %f for %#v", tC.want, got, tC.s)
			}
		})
	}
}
