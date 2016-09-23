package games

import "time"

// FPSCounter is a counter counting the frame per seconds currently displayed.
type FPSCounter struct {
	lastTick time.Time
}

// NewFPSCounter will build a new counter
func NewFPSCounter() *FPSCounter {
	return &FPSCounter{
		lastTick: time.Now(),
	}
}

// FPS returns the number of FPS estimated for the last frame
func (s *FPSCounter) FPS() float32 {
	now := time.Now()
	timePassed := now.UnixNano() - s.lastTick.UnixNano()
	fps := float32(time.Second) / float32(timePassed)
	s.lastTick = now
	return fps
}

// LimitedFPSCounter define a FPSCounter with a refresh rate, to limit the update of the counter.
type LimitedFPSCounter struct {
	*FPSCounter
	RefreshRate time.Duration

	lastCalculated time.Time
	result         float32
}

// NewLimitedFPSCounter will build a new counter, with default settings
func NewLimitedFPSCounter() *LimitedFPSCounter {
	return &LimitedFPSCounter{
		FPSCounter:     &FPSCounter{lastTick: time.Now()},
		RefreshRate:    time.Second / 2,
		lastCalculated: time.Now(),
	}
}

// FPS returns the number of FPS estimated since the last time the FPS was calculated. It refresh itself depending on RefreshRate
func (l *LimitedFPSCounter) FPS() float32 {
	timeSpentSinceLastCalculated := time.Duration(time.Now().UnixNano() - l.lastCalculated.UnixNano())
	if timeSpentSinceLastCalculated > l.RefreshRate || 0 == l.result {
		l.result = l.FPSCounter.FPS()
		l.lastCalculated = time.Now()
	}
	l.lastTick = time.Now()
	return l.result
}
