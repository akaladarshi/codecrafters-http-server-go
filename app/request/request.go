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

	header, err := NewHeader(splitData)
	if err != nil {
		return nil, fmt.Errorf("failed to get header: %w", err)
	}

	return &Request{
		header: header,
	}, nil
}

func (r *Request) GetHeader() *Header {
	return r.header
}

//func (r *Request) GetPath() string {
//	return r.header.GetPath()
//}
