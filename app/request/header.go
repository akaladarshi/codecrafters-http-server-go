package request

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/protocol"
)

const (
	Protocol = "HTTP"
)

type Header struct {
	version protocol.Protocol
	method  Method
	path    string
}

func NewHeader(data []byte) (*Header, error) {
	splitData := bytes.Split(data, []byte(" "))
	if len(splitData) < 2 {
		return nil, errInvalidHeaderData
	}

	return parseHeader(splitData)
}

func (h *Header) GetProtocolVersion() protocol.Protocol {
	return h.version
}

func (h *Header) GetMethod() Method {
	return h.method
}

func parseHeader(data [][]byte) (*Header, error) {
	var (
		method  Method
		version protocol.Protocol
		path    string
		err     error
	)

	method, err = CreateMethod(string(data[0]))
	if err != nil {
		return nil, fmt.Errorf("failed to parse method: %w", err)
	}

	switch len(data) {
	case 2: // method, protocol
		version, err = protocol.CreateVersion(string(data[1]))
		if err != nil {
			return nil, fmt.Errorf("failed to parse protocol: %w", err)
		}
	case 3: // method, path, protocol
		path = string(data[1])
		version, err = protocol.CreateVersion(string(data[2]))
		if err != nil {
			return nil, fmt.Errorf("failed to parse protocol: %w", err)
		}
	}

	return &Header{
		version: version,
		method:  method,
		path:    path,
	}, nil

}

func (h *Header) IsPathExist() bool {
	return h.path != ""
}

func (h *Header) GetPath() string {
	if h.IsPathExist() {
		return h.path
	}

	return ""
}
