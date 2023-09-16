package main

import "testing"

func TestMessage(t *testing.T) {

	msg := "Kamalesh"

	if msg != "Kamalesh" {
		t.Errorf("got %q, wanted %q", msg, "Kamalesh")

	}
}

func TestMain(t *testing.T) {
	main()
}
