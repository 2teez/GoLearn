package main

import (
	"testing"

	"github.com/practice/msg/msg"
)

func TestMsgGreetingPackage(t *testing.T) {

	expected, got := msg.Greeting("Hello, World"), "Hello, World"
	if expected != got {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}
