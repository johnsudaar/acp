package hyperdeck

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Clips []Clip
type Clip struct {
	ID       int
	Name     string
	StartAt  Timecode
	Duration Timecode
}

func (c *Client) ClipsGet() (Clips, error) {
	res, err := c.Send([]byte("clips get\n"))
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
	// Ignore first line (205 clip info:)
	_, err = buff.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "fail parse response (header)")
	}
	// Ignore second line (clip count: N)
	_, err = buff.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "fail to parse response (count)")
	}

	clips := make(Clips, 0)

	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return clips, nil
			}
			return nil, errors.Wrap(err, "fail tp parse response (body)")
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			return clips, nil
		}

		clip, err := ParseClipLine(line)
		if err != nil {
			return nil, errors.Wrap(err, "fail to parse current clip")
		}

		clips = append(clips, clip)
	}
}

func ParseClipLine(line string) (Clip, error) {
	var clip Clip
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
	if len(metadata) < 4 {
		return clip, fmt.Errorf("invalid metadata: %s", values[1])
	}

	clip.Name = strings.Join(metadata[1:len(metadata)-2], " ")
	startAt, err := ParseTimecode(metadata[len(metadata)-2])
	if err != nil {
		return clip, errors.Wrap(err, "invalid startAt")
	}
	clip.StartAt = startAt

	duration, err := ParseTimecode(metadata[len(metadata)-1])
	if err != nil {
		return clip, errors.Wrap(err, "invalid duration")
	}
	clip.Duration = duration

	return clip, nil
}
