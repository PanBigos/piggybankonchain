package clock

import "time"

// Nower is a function that can return the current time.
// In Clock, Now() will use time.Now by default, but a Nower can be set using WithNower in NewClock
// to customize the return value for Now() in tests.
type Nower func() time.Time

// Clock abstracts important time-related concerns in the beacon chain:
//   - provides a time.Now() construct that can be overridden in tests
type Clock struct{ now Nower }

// Now provides a value for time.Now() that can be overridden in tests.
func (g *Clock) Now() time.Time {
	return g.now()
}

// Since is a shorthand for Now().Sub(t)
func (g *Clock) Since(t time.Time) time.Duration {
	return g.now().Sub(t)
}

// ClockOpt is a functional option to change the behavior of a clock value made by NewClock.
// It is primarily intended as a way to inject an alternate time.Now() callback (WithNower) for testing.
type ClockOpt func(*Clock)

// WithNower allows tests in particular to inject an alternate implementation of time.Now (vs using system time)
func WithNower(n Nower) ClockOpt {
	return func(g *Clock) {
		g.now = n
	}
}

// NewClock constructs a Clock.
// The WithNower ClockOpt can be used in tests to specify an alternate `time.Now` implementation,
func NewClock(opts ...ClockOpt) *Clock {
	c := &Clock{}
	for _, o := range opts {
		o(c)
	}
	if c.now == nil {
		c.now = time.Now
	}
	return c
}
