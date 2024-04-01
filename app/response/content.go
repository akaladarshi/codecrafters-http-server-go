package response

import (
	"errors"
	"fmt"
	"io"
)

var (
	errTypeNotSupported = errors.New("data type not supported")
)

type Content struct {
	dataType string
	data     string
	length   int64
}

const (
	contentType = "Content-Type"
	contentLen  = "Content-Length"

	// content data type
	text = "text/plain"
)

func CreateContent(data string) (*Content, error) {
	return &Content{
		dataType: text,
		data:     data,
		length:   int64(len(data)),
	}, nil
}

func (c *Content) writeData(w io.Writer) error {
	if _, err := io.WriteString(w, fmt.Sprintf("%s: %s", contentType, c.dataType+CRLF)); err != nil {
		return fmt.Errorf("failed to write content type: %w", err)
	}

	if _, err := io.WriteString(w, fmt.Sprintf("%s: %d%s%s", contentLen, c.length, CRLF, CRLF)); err != nil {
		return fmt.Errorf("failed to write content length: %w", err)
	}

	if _, err := io.WriteString(w, c.data); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}
