package tty

import (
	"fmt"
	"testing"
)

func TestTTYColorize(t *testing.T) {
	exp := fmt.Sprintf("%s%s%s%s", fgBlue, bgCyan, reset, reset)
	cc := FgBlue()
	cc.Add(BgCyan())
	if got := cc.Start() + cc.End(); got != exp {
		t.Errorf("expected to be %q got %q", exp, got)
	}
}
