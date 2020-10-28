package hyperdeck

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	Slot1       = 1
	Slot2       = 2
	SlotCurrent = 0
)

type DiskClips []DiskClip
type DiskClip struct {
	ID       int
	Slot     int
	Name     string
	Codec    string
	Format   string
	Duration Timecode
}

// DiskList will list clips in slot. If the slot is set to 0 it will list all slots
func (c *Client) DiskList(slot int) (DiskClips, error) {
	cmd := fmt.Sprintf("disk list")
	if slot != 0 {
		cmd = fmt.Sprintf("%s: slot id: %v", cmd, slot)
	}
	cmd += "\n"

	res, err := c.Send([]byte(cmd))
	if err != nil {
		return nil, errors.Wrap(err, "fail to send cmd")
	}

	if IsError(res) {
		parsedError, err := ParseError(res)
		if err != nil {
			return nil, errors.Wrap(err, "fail to parse error")
		} else {
			return nil, parsedError
		}
	}

	buff := bytes.NewBuffer(res)
	// Ignore first line (206 disk list:)
	_, err = buff.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "fail parse response (header)")
	}

	curSlot := -1
	clips := make([]DiskClip, 0)
	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			return nil, errors.Wrap(err, "fail to parse response (body)")
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			return clips, nil
		}

		if strings.HasPrefix(line, "slot id: ") {
			id := strings.TrimPrefix(line, "slot id: ")
			curSlot, err = strconv.Atoi(id)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to parse current slot: %s", line)
			}
		} else {
			clip, err := ParseDiskClip(line)
			if err != nil {
				return nil, errors.Wrap(err, "fail to parse current clip")
			}
			if curSlot == -1 {
				return nil, fmt.Errorf("Clip received before slot")
			}
			clip.Slot = curSlot
			clips = append(clips, clip)
		}
	}
}

func ParseDiskClip(line string) (DiskClip, error) {
	var clip DiskClip
	values := strings.SplitN(line, ":", 2)

	if len(values) != 2 {
		return clip, fmt.Errorf("invalid clip line: %s", line)
	}
	id, err := strconv.Atoi(values[0])
	if err != nil {
		return clip, errors.Wrapf(err, "Invalid clip ID: %s", values[0])
	}

	clip.ID = id

	metadata := strings.Split(values[1], " ")
	if len(metadata) < 5 {
		return clip, fmt.Errorf("invalid metadata: %s", values[1])
	}
	clip.Name = metadata[1]
	clip.Codec = metadata[2]
	clip.Format = metadata[3]
	duration, err := ParseTimecode(metadata[4])
	if err != nil {
		return clip, errors.Wrap(err, "invalid timecode")
	}
	clip.Duration = duration
	return clip, nil
}
