package client

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type StreamingSettings struct {
	Resolution string `json:"Resolution"`
	Framerate  string `json:"Framerate"`
	Bitrate    string `json:"Bitrate"`
}

func (c HTTPClient) GetStreamingSettings() (StreamingSettings, error) {
	var settings StreamingSettings
	resp, err := c.makeRequest("GetStreamingSettings", nil)
	if err != nil {
		return settings, errors.Wrap(err, "fail to call HTTP API")
	}

	err = json.Unmarshal(resp.Data, &settings)
	if err != nil {
		return settings, errors.Wrap(err, "invalid response")
	}
	return settings, nil
}

type setStreamingResolutionParams struct {
	Resolution string `json:"Resolution"`
}

func (c HTTPClient) SetStreamingResolution(resolution string) error {
	_, err := c.makeRequest("SetStreamingResolution", setStreamingResolutionParams{
		Resolution: resolution,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}

type setStreamingFramerateParams struct {
	Framerate string `json:"Framerate"`
}

func (c HTTPClient) SetStreamingFramerate(framerate string) error {
	_, err := c.makeRequest("SetStreamingFramerate", setStreamingFramerateParams{
		Framerate: framerate,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}

type setStreamingBitrateParams struct {
	Bitrate string `json:"Bitrate"`
}

func (c HTTPClient) SetStreamingBitrate(bitrate string) error {
	_, err := c.makeRequest("SetStreamingBitrate", setStreamingBitrateParams{
		Bitrate: bitrate,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}
