package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	wg.Add(1)
		go updateMessage("Kamalesh")
		defer wg.Done()
	wg.Wait()
	if msg != "Kamalesh" {
		t.Error("Expected Kamalesh, got ", msg)
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = "Kamalesh"
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Kamalesh") {
		t.Error("Expected to find Kamalesh")
	}

}

func TestMain(t *testing.T) {
	main()
}
