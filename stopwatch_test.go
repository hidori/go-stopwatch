package stopwatch

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStopwatch(t *testing.T) {
	tests := []struct {
		name string
		want *Stopwatch
	}{
		{
			name: "Success: returns a new instance.",
			want: &Stopwatch{
				start: nil,
				end:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStopwatch()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestStopwatch_Start(t *testing.T) {
	tests := []struct {
		name   string
		assert func(t *testing.T, sw *Stopwatch, now time.Time) bool
	}{
		{
			name: "Success: the value close to current time is set in the `start` field.",
			assert: func(t *testing.T, sw *Stopwatch, now time.Time) bool {
				if !assert.NotNil(t, sw.start) {
					return false
				}

				if !assert.Nil(t, sw.end) {
					return false
				}

				if !assert.True(t, math.Abs(float64(sw.start.Sub(now))) < float64(time.Second)) {
					return false
				}

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := NewStopwatch()

			sw.Start()
			if tt.assert != nil {
				if !tt.assert(t, sw, time.Now()) {
					return
				}
			}
		})
	}
}

func TestStopwatch_Stop(t *testing.T) {
	tests := []struct {
		name   string
		assert func(t *testing.T, sw *Stopwatch, now time.Time) bool
	}{
		{
			name: "Success: the value close to current time is set in the `end` field.",
			assert: func(t *testing.T, sw *Stopwatch, now time.Time) bool {
				if !assert.Nil(t, sw.start) {
					return false
				}

				if !assert.NotNil(t, sw.end) {
					return false
				}

				if !assert.True(t, math.Abs(float64(sw.end.Sub(now))) < float64(time.Second)) {
					return false
				}

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := NewStopwatch()

			sw.Stop()
			if tt.assert != nil {
				if !tt.assert(t, sw, time.Now()) {
					return
				}
			}
		})
	}
}

func TestStopwatch_Duration(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(t *testing.T, sw *Stopwatch) error
		assert func(t *testing.T, got time.Duration) bool
	}{
		{
			name: "Success: when start=nil and end=nil, returns 0s.",
			setup: func(t *testing.T, sw *Stopwatch) error {
				sw.start = nil
				sw.end = nil
				time.Sleep(1 * time.Second)

				return nil
			},
			assert: func(t *testing.T, got time.Duration) bool {
				return assert.Equal(t, time.Duration(0), got)
			},
		},
		{
			name: "Success: when start=not nil and end=nil, returns 1s+",
			setup: func(t *testing.T, sw *Stopwatch) error {
				sw.start = func() *time.Time { t := time.Now(); return &t }()
				sw.end = nil
				time.Sleep(1 * time.Second)

				return nil
			},
			assert: func(t *testing.T, got time.Duration) bool {
				return assert.True(t, got >= (1*time.Second))
			},
		},
		{
			name: "Success: when start=not nil and end=not nil, returns 500ms",
			setup: func(t *testing.T, sw *Stopwatch) error {
				now := func() *time.Time { t := time.Now(); return &t }()
				sw.start = now
				sw.end = func() *time.Time { t := now.Add(500 * time.Millisecond); return &t }()
				time.Sleep(1 * time.Second)

				return nil
			},
			assert: func(t *testing.T, got time.Duration) bool {
				return assert.Equal(t, 500*time.Millisecond, got)
			},
		},
		{
			name: "Success: when start=nil and end=not nil, returns 0s",
			setup: func(t *testing.T, sw *Stopwatch) error {
				sw.start = nil
				sw.end = func() *time.Time { t := time.Now(); return &t }()
				time.Sleep(1 * time.Second)

				return nil
			},
			assert: func(t *testing.T, got time.Duration) bool {
				return assert.Equal(t, time.Duration(0), got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := NewStopwatch()

			if tt.setup != nil {
				err := tt.setup(t, sw)
				if err != nil {
					t.Errorf("fail to tt.setup() err=%v", err)
					return
				}
			}

			got := sw.Duration()
			if tt.assert != nil {
				if !tt.assert(t, got) {
					return
				}
			}
		})
	}
}
