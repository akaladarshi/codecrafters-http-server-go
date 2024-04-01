package request

import "errors"

var (
	errFailedToParseHeader = errors.New("failed to parse header data")
	errInvalidHeaderData   = errors.New("invalid header data")

	errInvalidMethod = errors.New("invalid method")
)
