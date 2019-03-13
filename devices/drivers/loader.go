package drivers

import (
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/drivers/atem"
	jvc "github.com/johnsudaar/acp/devices/drivers/jvc_hm_660"
	tally "github.com/johnsudaar/acp/devices/drivers/tally_rasp"
	"github.com/johnsudaar/acp/devices/drivers/tallyrecorder"
)

func LoadDrivers() {
	devices.RegisterType("ATEM", atem.NewLoader())
	devices.RegisterType("JVC_HM_660", jvc.NewLoader())
	devices.RegisterType("TALLY_RASP", tally.NewLoader())
	devices.RegisterType("TALLY_REC", tallyrecorder.NewLoader())
}
