package arrays

import (
	"reflect"
	"testing"
)

func TestSummAll(t *testing.T) {
	// compareSlices := func(a, b []int) bool {
	// 	t.Helper()
	// 	if len(a) != len(b) {
	// 		return false
	// 	}
	// 	for i, el := range a {
	// 		if el != b[i] {
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }

	testCases := []struct {
		desc string
		in   [][]int
		want []int
	}{
		{
			desc: "",
			in: [][]int{
				[]int{1, 2},
				[]int{2, 1},
			},
			want: []int{3, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := sumAll(tC.in)
			// if !compareSlices(got, tC.want) {
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("want %v, got %v for %v", tC.want, got, tC.in)
			}
		})
	}
}

func Test(t *testing.T) {
	got := sumAllTails([]int{12, 13}, []int{2, 1, 13}, []int{1}, []int{})
	want := []int{13, 14, 0, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
