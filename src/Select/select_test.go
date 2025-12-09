package selectPRA

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	/*
		httptest.NewServerはテスト用サーバ
	*/
	// slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(20 * time.Millisecond)
	// 	w.WriteHeader(http.StatusOK)
	// }))

	// fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))

	// slowURL := "http://www.facebook.com"
	// fastURL := "http://www.quii.co.uk"

	// want := fastURL
	// got := Racer(slowURL, fastURL)

	// if got != want {
	// 	t.Errorf("got %q, want %q", got, want)
	// }

	// slowServer.Close()
	// fastServer.Close()

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
