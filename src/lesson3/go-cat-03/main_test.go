package main

import (
	"testing"
)

func Test_cat(t *testing.T) {
	src := strings.NewReader("foo")
	dst := new(bytes.Buffer)

	cat(src, dst)

	expected := "foo"
	observed := dst.String()

	if expected != observed {
		t.Errorf(`
expected: %s
observed: %s
`, expected, observed)
	}

}

func Test_determineSources(t *testing.T) {
}
