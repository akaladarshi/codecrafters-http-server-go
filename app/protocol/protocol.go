package protocol

import "errors"

type Protocol string

var (
	errInvalidVersion = errors.New("invalid protocol")
)

const (
	v1 = "HTTP/1.1"
	v2 = "HTTP/2.1"
)

func CreateVersion(s string) (Protocol, error) {
	v := Protocol(s)
	if !v.IsValid() {
		return "", errInvalidVersion
	}

	return v, nil
}

func (v Protocol) IsValid() bool {
	switch v {
	case v1, v2:
		return true
	default:
		return false
	}
}
