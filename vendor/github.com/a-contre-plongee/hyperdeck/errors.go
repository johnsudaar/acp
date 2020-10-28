package hyperdeck

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type HyperdeckError struct {
	StatusCode int
	Status     string
	Message    string
}

func (e HyperdeckError) Error() string {
	return fmt.Sprintf("%s\n%s", e.Status, e.Message)
}

func ParseError(payload []byte) (HyperdeckError, error) {
	er := HyperdeckError{}

	buff := bytes.NewBuffer(payload)

	header, err := buff.ReadString('\n')
	if err != nil {
		return er, errors.Wrap(err, "fail to read first line")
	}

	code, err := strconv.Atoi(header[:3])
	if err != nil {
		return er, errors.Wrapf(err, "invalid error code: %s", header[:3])
	}

	er.StatusCode = code
	er.Status = header

	for {
		line, err := buff.ReadString('\n')
		if err != nil && err != io.EOF {
			return er, errors.Wrap(err, "fail to read message")
		}
		if line == "\n" || (err != nil && err == io.EOF) {
			er.Message = strings.TrimSpace(er.Message)
			return er, nil
		}
		er.Message += line
	}
}
