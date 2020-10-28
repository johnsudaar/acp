package timer

import (
	"sync"
	"time"

	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/realtime"
)

const (
	ClockFormat = "15:04:05"
)

type Timer struct {
	startedAt time.Time
	pausedAt  time.Time
	paused    bool
	started   bool
	timer     models.Timer

	stopping bool
	stopLock *sync.Mutex

	oldType models.TimerType

	timerLock *sync.RWMutex
	realtime  realtime.Realtime
	graph     graph.Graph
}

func NewTimer(t models.Timer, realtime realtime.Realtime, graph graph.Graph) *Timer {
	return &Timer{
		started:   false,
		timer:     t,
		realtime:  realtime,
		graph:     graph,
		oldType:   t.Type,
		stopping:  false,
		timerLock: &sync.RWMutex{},
		stopLock:  &sync.Mutex{},
	}
}

func (t *Timer) UpdateTimer(timer models.Timer) {
	t.timerLock.Lock()
	t.timer = timer
	if t.oldType != timer.Type {
		t.started = false
		t.paused = false
	}

	t.timer = timer
	t.oldType = timer.Type
	t.timerLock.Unlock()

	t.refresh()
}

func (t *Timer) Do(action models.TimerAction) error {
	t.timerLock.Lock()
	defer t.timerLock.Unlock()
	switch action.Action {
	case models.TimerActionStart:
		if t.paused {
			t.startedAt = t.startedAt.Add(time.Since(t.pausedAt))
		} else {
			t.startedAt = time.Now()
		}
		t.started = true
		t.paused = false
	case models.TimerActionStop:
		t.started = false
		t.paused = false
	case models.TimerActionPause:
		t.pausedAt = time.Now()
		t.started = false
		t.paused = true
	case models.TimerActionReset:
		t.startedAt = time.Now()
		t.started = false
		t.paused = false
	}
	return nil
}

func (t *Timer) Start() {
	go t.runLoop()
}

func (t *Timer) Stop() {
	t.stopLock.Lock()
	t.stopping = true
	t.stopLock.Unlock()
}

func (t *Timer) Value() string {
	t.timerLock.RLock()
	defer t.timerLock.RUnlock()
	timer := t.timer
	if timer.Type == models.TimerTypeClock {
		loc, _ := time.LoadLocation("Europe/Paris")
		return time.Now().In(loc).Format(ClockFormat)
	}

	if timer.Type == models.TimerTypeCountDown {
		if t.started {
			return FormatDuration(timer.Duration.Duration - time.Since(t.startedAt))
		} else if t.paused {
			return FormatDuration(timer.Duration.Duration - time.Since(t.startedAt) + time.Since(t.pausedAt))
		} else {
			return FormatDuration(timer.Duration.Duration)
		}
	}

	if timer.Type == models.TimerTypeCountUp {
		if t.started {
			return FormatDuration(time.Since(t.startedAt))
		} else if t.paused {
			return FormatDuration(time.Since(t.startedAt) - time.Since(t.pausedAt))
		} else {
			return FormatDuration(time.Duration(0))
		}
	}

	if timer.Type == models.TimerTypeExternal {
		return t.ExternalDuration()
	}

	return "-00:00:00"
}

func (t *Timer) Expired() bool {
	t.timerLock.RLock()
	defer t.timerLock.RUnlock()
	if !t.started {
		return false
	}
	if t.timer.Type == models.TimerTypeCountDown {
		return time.Since(t.startedAt) > t.timer.Duration.Duration
	}
	return false
}

func (t *Timer) Direction() models.TimerDirection {
	return models.TimerDirectionDown
}

func (t *Timer) refresh() {
	t.timerLock.RLock()
	timer := t.timer
	t.timerLock.RUnlock()

	// TODO: Check error
	t.realtime.Publish(RealtimeUpdateChannel, timer.ID.Hex(), RealtimeUpdateEvent{
		Value:   t.Value(),
		Expired: t.Expired(),
		Type:    timer.Type,
	})
}

func (t *Timer) runLoop() {
	for {
		t.stopLock.Lock()
		stopping := t.stopping
		t.stopLock.Unlock()
		if stopping {
			return
		}

		t.refresh()
		time.Sleep(100 * time.Millisecond)
	}
}
