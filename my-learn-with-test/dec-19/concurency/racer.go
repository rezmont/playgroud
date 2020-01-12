package concurency

import (
	"errors"
	"net/http"
	"time"
)

var TimeOutError = errors.New("No response error")

func Racer(url1, url2 string) (string, error) {
	c1 := ping(url1)
	c2 := ping(url2)
	c3 := timeout(10 * time.Millisecond)
	select {
	case <-c1:
		return url1, nil
	case <-c2:
		return url2, nil
	case <-c3:
		return "", TimeOutError
	case <-time.After(10 * time.Millisecond):
		return "", TimeOutError
	}

}

func ping(u string) chan struct{} {
	c := make(chan struct{})
	go func() {
		http.Get(u)
		close(c)
	}()
	return c
}

func timeout(t time.Duration) chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(t)
		close(c)
	}()
	return c
}
