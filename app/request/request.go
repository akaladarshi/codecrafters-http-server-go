package request

import (
	"bytes"
	"fmt"
)

const (
	CRLF = "\r\n"
)

type Request struct {
	header *Header
	data   []byte
}

func NewRequest(req []byte) (*Request, error) {
	splitData := bytes.Split(req, []byte(CRLF))
	if len(splitData) < 1 {
		return nil, fmt.Errorf("invalid request")
	}

	// length represents actual array size
	dataLen := len(splitData) - 1

	// if their exist non-empty last element then last element will be body data
	// otherwise it will only contain header data
	header, err := NewHeader(splitData[:dataLen])
	if err != nil {
		return nil, fmt.Errorf("failed to get header: %w", err)
	}

	return &Request{
		header: header,
		data:   splitData[dataLen],
	}, nil
}

func (r *Request) GetHeader() *Header {
	return r.header
}

func (r *Request) Data() []byte {
	return r.data
}
