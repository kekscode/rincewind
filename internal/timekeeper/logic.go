package timekeeper

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID     uuid.UUID
	Slices []TimeSlice
}

type TimeSlice struct {
	Label    string
	Duration time.Duration
	Begin    time.Time
	End      time.Time
}

func New(label string, duration time.Duration) *TimeSlice {
	return &TimeSlice{
		Label:    label,
		Duration: duration,
	}
}

func (ts *TimeSlice) Start() {
	ts.Begin = time.Now()
	ts.End = ts.Begin.Add(ts.Duration)
}

func (ts *TimeSlice) TimeLeft() time.Duration {
	if ts.Expired() {
		return time.Duration(0)
	}
	return time.Until(ts.End)
}

func (ts *TimeSlice) TimeElapsed() time.Duration {
	if ts.Expired() {
		return ts.Duration
	}
	return time.Since(ts.Begin)
}

func (ts *TimeSlice) Expired() bool {
	return time.Now().After(ts.End)
}
