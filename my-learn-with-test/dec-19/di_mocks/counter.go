package di_mocks

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func CountDown(w io.Writer, s Sleeper) {
	for index := countdownStart; index > 0; index-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", index)
	}
	s.Sleep()
	fmt.Fprintln(w, finalWord)
}

func main() {
	realSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, realSleeper)
}
