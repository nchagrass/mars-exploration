package main

import (
	"github.com/nchagrass/mars-exploration/internal/bootstrap"
	"io/ioutil"
	"os"
	"testing"
)

func TestApp_ExploreMars(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	expectedOutput, _ := ioutil.ReadFile("./expectedoutputsample-1.txt")

	bootstrap.New("./inputsample-1.txt")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(expectedOutput) != string(out) {
		t.Fatalf("Failed to explore mars, got %s, want %s", string(out), string(expectedOutput))
	}
}
