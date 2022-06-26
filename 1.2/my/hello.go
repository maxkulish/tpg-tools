package hello

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer, s string) {
	fmt.Fprint(w, s)
}
