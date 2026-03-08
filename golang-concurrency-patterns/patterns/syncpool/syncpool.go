package syncpool

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// ProcessWithPool uses sync.Pool to reuse byte buffers.
func ProcessWithPool(data []byte) []byte {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	// Simulate processing: uppercase transformation
	for _, b := range data {
		if b >= 'a' && b <= 'z' {
			buf.WriteByte(b - 32)
		} else {
			buf.WriteByte(b)
		}
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	bufferPool.Put(buf)
	return result
}

// ProcessWithoutPool allocates a new buffer each time.
func ProcessWithoutPool(data []byte) []byte {
	buf := new(bytes.Buffer)

	for _, b := range data {
		if b >= 'a' && b <= 'z' {
			buf.WriteByte(b - 32)
		} else {
			buf.WriteByte(b)
		}
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	return result
}
