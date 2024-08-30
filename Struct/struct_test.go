package main

import (
	"testing"
)

func TestPeriment(t *testing.T) {
	reactanfle := Reactanfle{10.0, 10.0}
	got := Perimenter(reactanfle)
	want := float64(40)

	if got != want {
		t.Errorf("got :(%.2f), want:(%.2f)", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	got := Area(float32(20), float32(30))
// 	want := float64(600)

// 	if got != want {
// 		t.Errorf("got :(%.2f), want:(%.2f)", got, want)
// 	}
// }
