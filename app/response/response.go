package response

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/protocol"
	"net"
)

const (
	CRLF = "\r\n"
)

type Response struct {
	header *Header
}

func NewResponse(ver protocol.Protocol, code int) *Response {
	header := NewResponseHeader(ver, code)
	return &Response{
		header: header,
	}
}

func (r *Response) GetProcessData() []byte {
	return []byte(fmt.Sprintf("%s %d%s%s", r.header.version, r.header.status, CRLF, CRLF))
}

func (r *Response) WriteResponse(conn net.Conn) error {
	_, err := conn.Write(r.GetProcessData())
	if err != nil {
		return fmt.Errorf("failed to write response to connection: %w", err)
	}

	return nil
}
