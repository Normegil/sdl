package games

import "time"

// FPSCounter sends back the number of fps displayed per second
type FPSCounter interface {
	FPS() float32
}

// UnlimitedFPSCounter is a counter counting the frame per seconds currently displayed. The value is recalculated every frame
type UnlimitedFPSCounter struct {
	lastTick time.Time
}

// NewFPSCounter will build a new counter
func NewFPSCounter() *UnlimitedFPSCounter {
	return &UnlimitedFPSCounter{
		lastTick: time.Now(),
	}
}

// FPS returns the number of FPS estimated for the last frame
func (s *UnlimitedFPSCounter) FPS() float32 {
	now := time.Now()
	timePassed := now.UnixNano() - s.lastTick.UnixNano()
	fps := float32(time.Second) / float32(timePassed)
	s.lastTick = now
	return fps
}

// LimitedFPSCounter is a counter counting the frame per seconds currently displayed. The value is recalculated depending on RefreshRate
type LimitedFPSCounter struct {
	*UnlimitedFPSCounter
	RefreshRate time.Duration

	lastCalculated time.Time
	result         float32
}

// NewLimitedFPSCounter will build a new counter, with default settings
func NewLimitedFPSCounter() *LimitedFPSCounter {
	return &LimitedFPSCounter{
		UnlimitedFPSCounter: &UnlimitedFPSCounter{lastTick: time.Now()},
		RefreshRate:         time.Second / 2,
		lastCalculated:      time.Now(),
	}
}

// FPS returns the number of FPS estimated since the last time the FPS was calculated. It refresh itself depending on RefreshRate
func (l *LimitedFPSCounter) FPS() float32 {
	timeSpentSinceLastCalculated := time.Duration(time.Now().UnixNano() - l.lastCalculated.UnixNano())
	if timeSpentSinceLastCalculated > l.RefreshRate || 0 == l.result {
		l.result = l.UnlimitedFPSCounter.FPS()
		l.lastCalculated = time.Now()
	}
	l.lastTick = time.Now()
	return l.result
}
