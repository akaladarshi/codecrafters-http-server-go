package response

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/protocol"
	"io"
	"net"
)

const (
	CRLF = "\r\n"
)

type Response struct {
	header      *Header
	content     *Content
	skipContent bool
}

func NewResponse(ver protocol.Protocol, code int, data string) (*Response, error) {
	header := NewResponseHeader(ver, code)
	var (
		content *Content
		err     error
	)

	content, err = CreateContent(data)
	if err != nil {
		return nil, fmt.Errorf("failed to create content: %w", err)
	}

	return &Response{
		header:  header,
		content: content,
	}, nil
}

func (r *Response) writeHeader(w io.Writer, isHeaderWithoutContent bool) error {
	var endLine string
	if isHeaderWithoutContent {
		endLine = fmt.Sprintf("%s%s", CRLF, CRLF)
	} else {
		endLine = fmt.Sprintf("%s", CRLF)
	}

	_, err := io.WriteString(w, fmt.Sprintf("%s %d%s", r.header.version, r.header.status, endLine))
	if err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	return nil
}

func (r *Response) WriteResponse(conn net.Conn) error {
	var buf bytes.Buffer

	if r.content.length == 0 {
		// only write the header
		fmt.Println("Header without content")
		err := r.writeHeader(&buf, true)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Header with Content")
		err := r.writeHeader(&buf, false)
		if err != nil {
			return err
		}

		err = r.content.writeData(&buf)
		if err != nil {
			return fmt.Errorf("failed to write content data")
		}
	}

	_, err := conn.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write response to connection: %w", err)
	}

	return nil
}
