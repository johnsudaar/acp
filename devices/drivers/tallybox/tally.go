package tallybox

import (
	"context"
	"errors"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices/types/tally"
	rpio "github.com/stianeikeland/go-rpio"
)

// Cam 1:
// - PVW: GPIO 4, Color: Yellow
// - PGM: GPIO 25, Color: Orange
// Cam 2:
// - PVW: GPIO 24, Color: Orange
// - PGM: GPIO 23, Color: Pink
// Cam 3:
// - PVW: GPIO 22, Color: Red
// - PGM: GPIO 27, Color: Brown
// Cam 4:
// - PVW: GPIO 18, Color: Blue
// - PGM: GPIO 17, Color: Green

var (
	previewPorts = []uint8{4, 24, 22, 18}
	programPorts = []uint8{25, 23, 27, 17}
)

func (t *tallybox) initGPIO() {
	for _, pin := range previewPorts {
		rpio.Pin(pin).Output()
	}

	for _, pin := range programPorts {
		rpio.Pin(pin).Output()
	}
}

func (t *tallybox) SendTally(ctx context.Context, port string, value tally.Value) {
	log := logger.Get(ctx)
	portIdx, err := t.portToIndex(port)
	if err != nil {
		log.WithError(err).Error("Fail to open port")
	}

	pvwPin := rpio.Pin(previewPorts[portIdx])
	pgmPin := rpio.Pin(programPorts[portIdx])

	switch value {
	case tally.Program:
		pvwPin.Low()
		pgmPin.High()
	case tally.Preview:
		pvwPin.High()
		pgmPin.Low()
	case tally.Off:
		pvwPin.Low()
		pgmPin.Low()
	}
}

func (t *tallybox) portToIndex(currentPort string) (uint8, error) {
	ports := t.OutputPorts()

	for idx, port := range ports {
		if port == currentPort {
			return uint8(idx), nil
		}
	}

	return 0, errors.New("Not found")
}
