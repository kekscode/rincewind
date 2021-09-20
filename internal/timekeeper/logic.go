package timekeeper

import "time"

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

func (ts *TimeSlice) TimeLeft() time.Time {
	return time.Now()
}

func (ts *TimeSlice) Expired() bool {
	return true
}
