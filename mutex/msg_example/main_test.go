package main

import "testing"

func Test_UpdateMessage(t *testing.T) {
	msg = "Kamalesh"
	wg.Add(1)
	go UpdateMessage("Hi I am Folks")
	wg.Wait()
	if msg != "Hi I am Folks" {
		t.Error("Expected message to be Hi I am Folks but got:", msg)
	}
}


