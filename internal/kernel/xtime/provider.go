package xtime

import "time"

type Provider interface {
	Now() time.Time
}

type ActualProvider struct{}

func (p ActualProvider) Now() time.Time {
	return time.Now()
}
