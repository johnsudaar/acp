package graph

import "github.com/johnsudaar/acp/devices"

type DevicePort struct {
	ID   string `json:"id"`
	Port string `json:"port"`
}

type Device struct {
	Type   string         `json:"type"`
	Device devices.Device `json:"-"`
}

type Graph struct {
	Devices map[string]Device `json:"devices"`
}
