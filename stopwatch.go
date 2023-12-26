package stopwatch

import "time"

type Stopwatch struct {
	start *time.Time
	end   *time.Time
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{}
}

func (sw *Stopwatch) Start() {
	t := time.Now()
	sw.start = &t
	sw.end = nil
}

func (sw *Stopwatch) Stop() {
	t := time.Now()
	sw.end = &t
}

func (sw *Stopwatch) Duration() time.Duration {
	if sw.start != nil {
		if sw.end == nil {
			return time.Since(*sw.start)
		}

		return sw.end.Sub(*sw.start)
	}

	return 0
}
