func TestPeriment(t *testing.T) {
	got := Perimenter(10.0, 10, 0)
	want := 40.0

	if got != want {
		t.Errorf("got :(%q), want:(%q)", got, want)
	}
}