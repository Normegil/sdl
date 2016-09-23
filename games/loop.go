package games

import (
	"time"

	"github.com/normegil/log"
)

// LoopCtrl used to control the main game loop
type LoopCtrl struct {
	FPS  FPSControls
	Quit bool
}

// FPSControls used to control the number of frames per second in the loop
type FPSControls struct {
	Number int
	Capped bool
}

// MainLoop hold the data needed to launch the main loop of a game
type MainLoop struct {
	Log     log.AgnosticLogger
	Control LoopCtrl
}

// Launch the main game loop, with fps capped if needed. It will run the given function every frame
func (m MainLoop) Launch(toExec func(LoopCtrl) (LoopCtrl, error)) error {
	m.Log.With(log.Structure{"Controls": m.Control}).Log(log.DEBUG, "Launching main loop")
	ctrl := m.Control
	var err error
	for !ctrl.Quit {
		beforeLoop := time.Now()
		ctrl, err = toExec(ctrl)
		if nil != err {
			return err
		}
		if ctrl.FPS.Capped && !ctrl.Quit {
			afterLoop := time.Now()
			time.Sleep(timeToSleep(ctrl.FPS.Number, beforeLoop, afterLoop))
		}
	}
	return nil
}

func timeToSleep(framePerSeconds int, beforeLoop, afterLoop time.Time) time.Duration {
	spentTime := time.Duration(toMilliseconds(afterLoop) - toMilliseconds(beforeLoop))
	theoriticalTimeToWait := time.Duration(1000 / framePerSeconds)
	timeToWait := theoriticalTimeToWait - spentTime
	return timeToWait * time.Millisecond
}

func toMilliseconds(t time.Time) int64 {
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}
