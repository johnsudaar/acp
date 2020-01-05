package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const DefaultPort = 60030

const (
	pcktStart byte = 0x02
	pcktEnd   byte = 0x03
)

type Bus string
type Input string

const (
	BusA    Bus = "00"
	BusB    Bus = "01"
	BusPgm  Bus = "02"
	BusPvw  Bus = "03"
	BusKeyF Bus = "04"
	BusKeyS Bus = "05"
	BusPinP Bus = "10"
	BusAux  Bus = "12"

	InputXPT1  Input = "00"
	InputXPT2  Input = "01"
	InputXPT3  Input = "02"
	InputXPT4  Input = "03"
	InputXPT5  Input = "04"
	InputXPT6  Input = "05"
	InputXPT7  Input = "06"
	InputXPT8  Input = "07"
	InputXPT9  Input = "08"
	InputXPT10 Input = "09"

	Input1 Input = "50"
	Input2 Input = "51"
	Input3 Input = "52"
	Input4 Input = "53"
	Input5 Input = "54"

	InputColorBars    Input = "70"
	InputColor        Input = "71"
	InputBlack        Input = "72"
	InputFrameMemory1 Input = "73"
	InputFrameMemory2 Input = "74"
	InputPGM          Input = "77"
	InputPVW          Input = "78"
	InputKeyOut       Input = "79"
	InputClean        Input = "80"
	InputMultiView    Input = "81"
)

type Client struct {
	ip   string
	port int
}

func New(ip string) Client {
	return Client{
		ip:   ip,
		port: DefaultPort,
	}
}

func (c Client) sendCommand(command string, params []string, expectResponse bool) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte(pcktStart)
	buffer.WriteString(fmt.Sprintf("%s:%s", command, strings.Join(params, ":")))
	buffer.WriteByte(pcktEnd)

	conn, err := net.Dial("tcp", c.addr())
	if err != nil {
		return "", errors.Wrap(err, "fail to dial switcher")
	}
	defer conn.Close()

	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		return "", errors.Wrap(err, "fail to send packet")
	}

	if !expectResponse {
		return "", nil
	}

	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", errors.Wrap(err, "fail to read device response")
	}

	return string(resp), nil
}

func (c Client) addr() string {
	return net.JoinHostPort(c.ip, strconv.Itoa(c.port))
}

func (c Client) SwitchBus(bus Bus, input Input) error {
	_, err := c.sendCommand("SBUS", []string{string(bus), string(input)}, false)
	if err != nil {
		return errors.Wrap(err, "fail to switch bus")
	}

	return nil
}
