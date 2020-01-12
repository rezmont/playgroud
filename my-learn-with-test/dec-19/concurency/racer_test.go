package concurency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("when response comes back in time", func(t *testing.T) {
		fastServer := makeTestServer(1 * time.Millisecond)
		slowServer := makeTestServer(20 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("when timesout", func(t *testing.T) {
		fastServer := makeTestServer(15 * time.Millisecond)
		slowServer := makeTestServer(20 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err != TimeOutError {
			t.Errorf("got %+v, wanted %+v", err, TimeOutError)
		}
	})
}

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
