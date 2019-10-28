package integers

import (
	"fmt"
	"testing"
)

func TestThis(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		sum int
	}{
		{
			2, 2, 4,
		},
		{
			2, -2, 0,
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			s := add(tC.a, tC.b)
			if tC.sum != s {
				t.Errorf("wanted %v, got %v", tC.sum, s)
			}
		})
	}
}

func ExampleAdd() {
	sum := add(1, 109)
	fmt.Println(sum)
	// outpu 110
}
