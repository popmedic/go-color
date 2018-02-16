package colorize

import (
	"testing"
)

func TestColorize(t *testing.T) {
	exp := "123456"
	cc := NewColorize("1", "6", NewColorize("2", "5", NewColorize("3", "4")),
		NewColorize("4", "3", NewColorize("5", "2"), NewColorize("6", "1")))
	if got := cc.Start(); got != exp {
		t.Errorf("Start expected to be %q got %q", exp, got)
	}
	if got := cc.End(); got != exp {
		t.Errorf("End expected to be %q got %q", exp, got)
	}
	if got := cc.Color(""); got != exp+exp {
		t.Errorf("Color expected to be %q got %q", exp+exp, got)
	}
}
