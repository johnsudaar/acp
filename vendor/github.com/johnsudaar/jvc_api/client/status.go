package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
	errgo "gopkg.in/errgo.v1"
)

type AEStatus string
type CamStatus string
type ExposureStatus string
type FocusStatus string
type GainStatus string
type IrisStatus string
type ProtectStatus string
type SlotStatus string
type StreamingStatus string
type ShutterStatus string
type WhiteBalanceStatus string
type Mode string
type RecMode string

const (
	CamStatusNoCard    CamStatus = "NoCard"
	CamStatusStop      CamStatus = "Stop"
	CamStatusStandby   CamStatus = "Standby"
	CamStatusRecording CamStatus = "Rec"
	CamStatusRecPause  CamStatus = "RecPause"

	ExposureStatusAuto            ExposureStatus = "Auto"
	ExposureStatusManual          ExposureStatus = "Manual"
	ExposureStatusIrisPriority    ExposureStatus = "IrisPriority"
	ExposureStatusShutterPriority ExposureStatus = "ShutterPriority"

	IrisStatusAuto       IrisStatus = "Auto"
	IrisStatusManual     IrisStatus = "Manual"
	IrisStatusAutoAELock IrisStatus = "AutoAELock"

	GainStatusManualL   GainStatus = "ManualL"
	GainStatusManualR   GainStatus = "ManualR"
	GainStatusManualH   GainStatus = "ManualH"
	GainStatusAlc       GainStatus = "Alc"
	GainStatusAlcAELock GainStatus = "AlcAELock"
	GainStatusLolux     GainStatus = "Lolux"
	GainStatusVariable  GainStatus = "Variable"

	ModeCamera       Mode = "Camera"
	ModeThumbnail    Mode = "Thumbnail"
	ModePlay         Mode = "Play"
	ModeReview       Mode = "Review"
	ModeUSB          Mode = "USB"
	ModeEditMetadata Mode = "EditMetadata"

	RecModeNormal   RecMode = "Normal"
	RecModePre      RecMode = "Pre"
	RecModeClip     RecMode = "Clip"
	RecModeFrame    RecMode = "Frame"
	RecModeInterval RecMode = "Interval"
	RecModeVariable RecMode = "Variable"

	AEStatusOff    AEStatus = "AeOff"
	AEStatusOn     AEStatus = "AeOn"
	AEStatusOnFace AEStatus = "AeOnFace"

	ShutterStatusOff       ShutterStatus = "Off"
	ShutterStatusManual    ShutterStatus = "Manual"
	ShutterStatusStep      ShutterStatus = "Step"
	ShutterStatusVariable  ShutterStatus = "Variable"
	ShutterStatusEEI       ShutterStatus = "Eei"
	ShutterStatusEEIAELock ShutterStatus = "EeiAELock"

	WhilteBalanceStatusPreset    WhiteBalanceStatus = "Preset"
	WhilteBalanceStatusPresetA   WhiteBalanceStatus = "PresetA"
	WhilteBalanceStatusPresetB   WhiteBalanceStatus = "PresetB"
	WhilteBalanceStatusFaw       WhiteBalanceStatus = "Faw"
	WhilteBalanceStatusFawAELock WhiteBalanceStatus = "FawAELock"

	FocusStatusAutoFace      FocusStatus = "AFFace"
	FocusStatusAuto          FocusStatus = "AF"
	FocusStatusManualOnePush FocusStatus = "MFOnePush"
	FocusStatusManual        FocusStatus = "MF"
	FocusStatusManualFace    FocusStatus = "MFFace"

	StreamingStatusStop     StreamingStatus = "Stop"
	StreamingStatusStopping StreamingStatus = "Stopping"
	StreamingStatusStart    StreamingStatus = "Start"
	StreamingStatusStarting StreamingStatus = "Starting"
	StreamingStatusWaiting  StreamingStatus = "Waiting"
	StreamingStatusError    StreamingStatus = "Error"

	SlotStatusSelect      SlotStatus = "Select"
	SlotStatusNoSelect    SlotStatus = "NoSelect"
	SlotStatusSelectRec   SlotStatus = "SelectRec"
	SlotStatusNoSelectRec SlotStatus = "NoSelectRec"
	SlotStatusInvalid     SlotStatus = "Invalid"
	SlotStatusError       SlotStatus = "Error"
)

type CameraStatus struct {
	Camera       CameraDetailedStatus       `json:"Camera"`
	Iris         IrisDetailedStatus         `json:"Iris"`
	Gain         GainDetailedStatus         `json:"Gain"`
	AELevel      AELevelDetailedStatus      `json:"AeLevel"`
	Shutter      ShutterDefailedStatus      `json:"Shutter"`
	WhiteBalance WhiteBalanceDetailedStatus `json:"Whb"`
	Zoom         ZoomDetailedStatus         `json:"Zoom"`
	Focus        FocusDetailedStatus        `json:"Focus"`
	CharaterMix  CharacterMixDetailedStatus `json:"CharaterMix"`
	TallyLamp    TallyLampDetailedStatus    `json:"TallyLamp"`
	SlotA        SlotDetailedStatus         `json:"SlotA"`
	SlotB        SlotDetailedStatus         `json:"SlotB"`

	FullAutoStatus struct {
		Status Status `json:"Status"`
	} `json:"FullAutoStatus"`

	ExposureStatus struct {
		Status ExposureStatus `json:"Status"`
	}

	MasterBlack struct {
		Value string `json:"Value"`
	} `json:"MasterBlack"`

	Detail struct {
		Value string `json:"Value"`
	} `json:"Detail"`

	Streaming struct {
		Status StreamingStatus `json:"Status"`
	} `json:"Streaming"`

	DispTV struct {
		Status Status `json:"Status"`
	} `json:"Disptv"`
}

type CameraDetailedStatus struct {
	Status            CamStatus   `json:"Status"`
	Mode              Mode        `json:"Mode"`
	RecordingMode     RecMode     `json:"RecMode"`
	Timecode          string      `json:"TC"`
	AspectRatio       AspectRatio `json:"AspectRatio"`
	WebAccess         Status      `json:"WebAccess"`
	VideoOutputStatus Status      `json:"VideoOutputStatus"`
	MenuStatus        Status      `json:"MenuStatus"`
}

func (c CameraDetailedStatus) TCTime() (time.Duration, error) {
	seconds, err := strconv.Atoi(c.Timecode)
	if err != nil {
		return 0, errgo.Notef(err, "invalid timecode")
	}

	return time.Duration(seconds) * time.Second, nil
}

type IrisDetailedStatus struct {
	Status IrisStatus `json:"Status"`
	Value  string     `json:"Value"`
}

type GainDetailedStatus struct {
	Status GainStatus `json:"Status"`
	Value  string     `json:"Value"`
}

type AELevelDetailedStatus struct {
	Status AEStatus `json:"Status"`
	Adjust Status   `json:"Adjust"`
	Value  string   `json:"Value"`
}

type ShutterDefailedStatus struct {
	Status ShutterStatus `json:"Status"`
	Value  string        `json:"Value"`
}

type WhiteBalanceDetailedStatus struct {
	Status                 WhiteBalanceStatus `json:"Status"`
	Value                  string             `json:"Value"`
	WhitePaintRedScale     int                `json:"WhPRScale"`
	WhitePaintBlueScale    int                `json:"WhPBScale"`
	WhitePaintRedPosition  int                `json:"WhPRPosition"`
	WhitePaintBluePosition int                `json:"WhPBPosition"`
	WhitePaintRedValue     string             `json:"WhPRValue"`
	WhitePaintBlueValue    string             `json:"WhPBValue"`
}

type ZoomDetailedStatus struct {
	Position     string `json:"Position"`
	DisplayValue string `json:"DisplayValue"`
}

type FocusDetailedStatus struct {
	Status FocusStatus `json:"Status"`
	Value  string      `json:"Value"`
}

type CharacterMixDetailedStatus struct {
	SDI   Status `json:"Sdi"`
	HDMI  Status `json:"Hdmi"`
	Video Status `json:"Video"`
}

type TallyLampDetailedStatus struct {
	Priority    TallyPriority   `json:"Priority"`
	Lighting    Status          `json:"Lighting"`
	StudioTally TallyIndication `json:"StudioTally"`
}

type SlotDetailedStatus struct {
	Status        SlotStatus    `json:"Status"`
	Protect       ProtectStatus `json:"Protect"`
	Remain        string        `json:"Remain"`
	ClipNum       int           `json:"ClipNum"`
	RemainWarning int           `json:"RemainWarning"`
}

func (c HTTPClient) GetCamStatus() (CameraStatus, error) {
	var camStatus CameraStatus
	resp, err := c.makeRequest("GetCamStatus", nil)
	if err != nil {
		return camStatus, errors.Wrap(err, "fail to call HTTP API")
	}

	err = json.Unmarshal(resp.Data, &camStatus)
	if err != nil {
		return camStatus, errors.Wrap(err, "invalid response")
	}

	return camStatus, nil
}
