package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintTo(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}
	want := "Hello, World!"

	hello.PrintTo(fakeTerminal, want)
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
