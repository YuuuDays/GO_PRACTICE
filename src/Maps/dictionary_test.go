package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a pen"}

	got := Search(dictionary, "test")
	want := "this is a pen"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
