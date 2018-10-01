package main

import (
	"testing"
)

func TestTest(t *testing.T) {
	if Hello() != "HELLO WORLD" {
		t.Fatalf("unexpected greetings %q", Hello())
	}
}
