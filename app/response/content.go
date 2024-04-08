package response

import (
	"errors"
	"fmt"
	"io"
)

var (
	errTypeNotSupported = errors.New("data type not supported")
)

type DataType string

const (
	PlainText DataType = "text/plain"
	FileData           = "application/octet-stream"
)

func (t DataType) String() string {
	return string(t)
}

type Content struct {
	dataType DataType
	data     []byte
	length   int
}

const (
	contentType = "Content-Type"
	contentLen  = "Content-Length"
)

func CreateContent(data []byte, dataType DataType) *Content {
	return &Content{
		dataType: dataType,
		data:     data,
		length:   len(data),
	}
}

func (c *Content) writeData(w io.Writer) error {
	if c.length == 0 {
		return nil
	}

	if _, err := io.WriteString(w, fmt.Sprintf("%s: %s", contentType, c.dataType+CRLF)); err != nil {
		return fmt.Errorf("failed to write content type: %w", err)
	}

	if _, err := io.WriteString(w, fmt.Sprintf("%s: %d%s%s", contentLen, c.length, CRLF, CRLF)); err != nil {
		return fmt.Errorf("failed to write content length: %w", err)
	}

	var err error
	switch c.dataType {
	case PlainText:
		_, err = io.WriteString(w, string(c.data))
	default:
		_, err = w.Write(c.data)
	}

	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}
