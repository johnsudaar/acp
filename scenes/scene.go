package scenes

import (
	"sync"

	"github.com/johnsudaar/acp/realtime"
)

const SceneActiveChannel = "scene_active"

type SceneActiveEvent struct {
	SceneID string `json:"scene_id"`
}

type Scenes interface {
	Active() string
	SetActive(a string)
}

type ScenesServer struct {
	active     string
	activeLock *sync.RWMutex
	realtime   realtime.Realtime
}

func New(r realtime.Realtime) Scenes {
	return &ScenesServer{
		active:     "",
		activeLock: &sync.RWMutex{},
		realtime:   r,
	}
}

func (s *ScenesServer) Active() string {
	s.activeLock.RLock()
	defer s.activeLock.RUnlock()

	return s.active
}

func (s *ScenesServer) SetActive(id string) {
	s.activeLock.Lock()
	s.active = id
	s.activeLock.Unlock()

	s.realtime.Publish(SceneActiveChannel, "controller", SceneActiveEvent{
		SceneID: s.active,
	})
}
