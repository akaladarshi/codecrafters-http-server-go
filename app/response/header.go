package response

import "github.com/codecrafters-io/http-server-starter-go/app/protocol"

type Header struct {
	version protocol.Protocol
	status  int
}

func NewResponseHeader(protocolVer protocol.Protocol, status int) *Header {
	return &Header{
		version: protocolVer,
		status:  status,
	}
}
