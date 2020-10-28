package hyperdeck

import (
	"time"

	"github.com/a-contre-plongee/hyperdeck"
	"github.com/johnsudaar/acp/devices/types/timer"
)

const (
	CountDown = "countdown"
	CountUp   = "countup"
)

func (h *Hyperdeck) TimecodeSources() []string {
	return []string{
		CountDown,
		CountUp,
	}
}

func (h *Hyperdeck) Timecode(source string) (timer.TimerValue, error) {
	h.cacheLock.Lock()
	if h.currentClip == 0 {
		h.cacheLock.Unlock()
		return timer.NewNAValue(), nil
	}
	clip, found := h.clipCache[h.currentClip]
	timecode := h.timecode
	lastRefreshedAt := h.lastRefreshedAt
	playing := h.playing
	h.cacheLock.Unlock()

	if !found {
		return timer.NewNAValue(), nil
	}

	clipTC := timecode.Duration()
	clipTC -= clip.StartAt.Duration()
	if playing {
		clipTC += time.Since(lastRefreshedAt)
	}
	if clipTC >= clip.Duration.Duration() {
		clipTC = clip.Duration.Duration()
	}

	if source == CountUp {
		return timer.NewDurationValue(clipTC), nil
	}

	if source == CountDown {
		return timer.NewDurationValue(clip.Duration.Duration() - clipTC), nil
	}

	return timer.NewNAValue(), nil
}

func (h *Hyperdeck) timers() {
	for {
		time.Sleep(1 * time.Second)
		// Mutex protection
		h.stoppingLock.Lock()
		stopping := h.stopping
		h.stoppingLock.Unlock()
		if stopping {
			return
		}

		h.refreshTransportInfo()

		h.ensureClipInCache()
	}
}

func (h *Hyperdeck) onSlotEvent(slot hyperdeck.Slot) {
	if slot.Status != hyperdeck.SlotStatusEmpty && slot.Status != hyperdeck.SlotStatusMounted {
		return
	}

	h.cacheLock.Lock()
	defer h.cacheLock.Unlock()
	// Disk changed, reset clip cache
	h.clipCache = make(map[int]hyperdeck.Clip)

	h.unsafeRefreshDiskCache()

	if h.currentClip == 0 {
		return
	}
	_, ok := h.clipCache[h.currentClip]
	if !ok {
		h.currentClip = 0
	}
}

func (h *Hyperdeck) onTransportEvent(transport hyperdeck.Transport) {
	h.cacheLock.Lock()
	defer h.cacheLock.Unlock()

	refreshTransportInfo := false

	if transport.ClipID != 0 && transport.ClipID != h.currentClip {
		h.currentClip = transport.ClipID
		refreshTransportInfo = true
	}

	if transport.Status == hyperdeck.TransportStatusPlay {
		refreshTransportInfo = true
	}
	if transport.Status == hyperdeck.TransportStatusStopped {
		refreshTransportInfo = true
	}

	if refreshTransportInfo {
		h.refreshTransportInfo()
	}
}

func (h *Hyperdeck) ensureClipInCache() {
	h.cacheLock.Lock()
	defer h.cacheLock.Unlock()

	_, inCache := h.clipCache[h.currentClip]
	if inCache {
		return
	}

	h.unsafeRefreshDiskCache()
}

func (h *Hyperdeck) unsafeRefreshDiskCache() {
	log := h.log
	client := h.Client()
	if client == nil {
		return
	}

	clips, err := client.ClipsGet()
	if err != nil {
		log.WithError(err).Error("fail to list slots")
		return
	}

	for _, clip := range clips {
		h.clipCache[clip.ID] = clip
	}
}

func (h *Hyperdeck) refreshTransportInfo() {
	log := h.log
	client := h.client
	if client == nil {
		return
	}

	transport, err := client.TransportInfo()
	if err != nil {
		log.WithError(err).Error("fail to get transport info")
		return
	}
	h.cacheLock.Lock()
	h.playing = transport.Status == hyperdeck.TransportStatusPlay
	h.currentClip = transport.ClipID
	h.timecode = transport.Timecode
	h.lastRefreshedAt = time.Now()
	h.cacheLock.Unlock()

}
