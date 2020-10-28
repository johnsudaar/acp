package hyperdeck

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Timecode struct {
	Hours   int
	Minutes int
	Seconds int
	Frames  int
}

func (t Timecode) String() string {
	return fmt.Sprintf("%02d:%02d:%02d:%02d", t.Hours, t.Minutes, t.Seconds, t.Frames)
}

func (t *Timecode) FromString(timecode string) error {
	fields := strings.Split(timecode, ":")
	if len(fields) != 4 {
		return fmt.Errorf("invalid timecode: %s", timecode)
	}

	hours, err := strconv.Atoi(fields[0])
	if err != nil {
		return errors.Wrapf(err, "fail to parse timecode (hour): %s", timecode)
	}
	minutes, err := strconv.Atoi(fields[1])
	if err != nil {
		return errors.Wrapf(err, "fail to parse timecode (minute): %s", timecode)
	}
	seconds, err := strconv.Atoi(fields[2])
	if err != nil {
		return errors.Wrapf(err, "fail to parse timecode (second): %s", timecode)
	}
	frames, err := strconv.Atoi(fields[3])
	if err != nil {
		return errors.Wrapf(err, "fail to parse timecode (frame): %s", timecode)
	}

	t.Hours = hours
	t.Minutes = minutes
	t.Seconds = seconds
	t.Frames = frames
	return nil
}

func (t Timecode) Duration() time.Duration {
	var res time.Duration
	res += time.Duration(t.Hours) * time.Hour
	res += time.Duration(t.Minutes) * time.Minute
	res += time.Duration(t.Seconds) * time.Second
	res += time.Duration(t.Frames) * (time.Second / 25)
	return res
}

func ParseTimecode(timecode string) (Timecode, error) {
	res := &Timecode{}
	err := res.FromString(timecode)
	if err != nil {
		return *res, errors.Wrap(err, "fail to parse timecode")
	}
	return *res, nil
}
