package drivers

import (
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/drivers/atem"
	"github.com/johnsudaar/acp/devices/drivers/discord"
	"github.com/johnsudaar/acp/devices/drivers/hs50"
	"github.com/johnsudaar/acp/devices/drivers/hyperdeck"
	jvc "github.com/johnsudaar/acp/devices/drivers/jvc_hm_660"
	jvcremote "github.com/johnsudaar/acp/devices/drivers/jvc_remote"
	restream "github.com/johnsudaar/acp/devices/drivers/restream_chat"
	"github.com/johnsudaar/acp/devices/drivers/smartview"
	tally "github.com/johnsudaar/acp/devices/drivers/tally_rasp"
	"github.com/johnsudaar/acp/devices/drivers/tallyrecorder"
)

func LoadDrivers() {
	devices.RegisterType("ATEM", atem.NewLoader())
	devices.RegisterType("JVC_HM_660", jvc.NewLoader())
	devices.RegisterType("JVC_REMOTE", jvcremote.NewLoader())
	devices.RegisterType("TALLY_RASP", tally.NewLoader())
	devices.RegisterType("TALLY_REC", tallyrecorder.NewLoader())
	devices.RegisterType("SMARTVIEW_DUO", smartview.NewLoader())
	//devices.RegisterType("TALLY_BOX", tallybox.NewLoader())
	devices.RegisterType("HS_50", hs50.NewLoader())
	devices.RegisterType("HYPERDECK", hyperdeck.NewLoader())
	devices.RegisterType("RESTREAM", restream.NewLoader())
	devices.RegisterType("DISCORD", discord.NewLoader())
}
