package di_mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCounter(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		spySleeper := &SpySleeper{Calls: 0}

		buffer := &bytes.Buffer{}
		CountDown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("wanted %s, but got %s", want, got)
		}

		wantCalls := 4

		if spySleeper.Calls != wantCalls {
			t.Errorf("sleep calls; wanted %d, but got %d", wantCalls, spySleeper.Calls)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Test(t *testing.T) {
	t.Run("configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
