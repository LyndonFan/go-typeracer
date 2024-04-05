package main

import "time"

type Timer struct {
	StartTime time.Time
	EndTime   time.Time
	Running   bool
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.StartTime = time.Now()
	t.Running = true
}

func (t *Timer) Stop() {
	t.EndTime = time.Now()
	t.Running = false
}

func (t *Timer) Duration() time.Duration {
	if t.Running {
		return time.Since(t.StartTime)
	}
	return t.EndTime.Sub(t.StartTime)
}

func (t *Timer) Reset() {
	t.StartTime = time.Time{}
	t.EndTime = time.Time{}
	t.Running = false
}
