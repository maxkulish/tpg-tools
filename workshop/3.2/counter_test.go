package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()

	c := counter.NewCounter
	c.Input = bytes.NewBufferString("1\n2\n3")
	want := 3
	got := c.Lines()

	if want != got {
		t.Error("want %q, got %q", want, got)
	}
}
