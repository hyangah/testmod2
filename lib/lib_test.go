package lib

import (
	"testing"
)

func TestHello(t *testing.T) {
	if Hello != "HELLO WORLD" {
		t.Fatalf("unexpected greetings %q", Hello)
	}
}
