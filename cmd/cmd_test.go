package main

import (
	"testing"
	"os/exec"
)

func TestTest(t *testing.T) {
	if Hello() != "Hello World" {
		t.Fatalf("unexpected greetings %q", Hello())
	}
}
