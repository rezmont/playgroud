package di_mocks

import (
	"fmt"
	"io"
)

func Greet(b io.Writer, s string) {
	fmt.Fprintf(b, "Hello, %s", s)
}
