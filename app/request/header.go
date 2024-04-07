package request

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/protocol"
	"strings"
)

const (
	Protocol = "HTTP"
)

type Header struct {
	version    protocol.Protocol
	method     Method
	path       string
	headerData map[string]string
}

func NewHeader(data [][]byte) (*Header, error) {
	version, method, path, err := GetStatusLine(data[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get status line: %w", err)
	}

	headerData, err := ParseHeaderData(data[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to parse header data: %w", err)
	}

	return &Header{
		version:    version,
		method:     method,
		path:       path,
		headerData: headerData,
	}, nil
}

func GetStatusLine(data []byte) (protocol.Protocol, Method, string, error) {
	splitData := bytes.Split(data, []byte(" "))
	if len(splitData) < 2 {
		return "", "", "", errInvalidHeaderData
	}

	return parseStatusLine(splitData)
}

func ParseHeaderData(data [][]byte) (map[string]string, error) {
	headerData := make(map[string]string)
	for _, v := range data {
		keyValueData := strings.SplitN(string(v), ":", 2)
		if len(keyValueData) < 2 {
			continue
		}

		headerData[strings.ToLower(keyValueData[0])] = strings.TrimSpace(keyValueData[1])
	}

	return headerData, nil
}

func (h *Header) GetProtocolVersion() protocol.Protocol {
	return h.version
}

func (h *Header) GetMethod() Method {
	return h.method
}

func parseStatusLine(data [][]byte) (protocol.Protocol, Method, string, error) {
	var (
		method  Method
		version protocol.Protocol
		path    string
		err     error
	)

	method, err = CreateMethod(string(data[0]))
	if err != nil {
		return "", "", "", fmt.Errorf("failed to parse method: %w", err)
	}

	switch len(data) {
	case 2: // method, protocol
		version, err = protocol.CreateVersion(string(data[1]))
		if err != nil {
			return "", "", "", fmt.Errorf("failed to parse protocol: %w", err)
		}
	case 3: // method, path, protocol
		path = string(data[1])
		version, err = protocol.CreateVersion(string(data[2]))
		if err != nil {
			return "", "", "", fmt.Errorf("failed to parse protocol: %w", err)
		}
	}

	return version, method, path, nil
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

func (h *Header) GetHeaderVal(key string) string {
	return h.headerData[key]
}
