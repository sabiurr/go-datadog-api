package datadog

import (
	"bytes"
)

// Compressor is an interace to compress the metrics payload for a HTTP request
type Compressor interface {
	Compress(req reqPostSeries) (*bytes.Buffer, error)
	GetContentEncoding() string
}
