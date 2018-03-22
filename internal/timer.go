package internal

import "time"

type Timer struct {
	begin time.Time
	end time.Time
}

func (t *Timer) Start() {
	t.begin = time.Now()
}

func (t *Timer) End() {
	t.end = time.Now()
}

func (t Timer) CostInMs() int64 {
	cost := t.end.UnixNano() - t.begin.UnixNano()
	return cost / 10e6
}
