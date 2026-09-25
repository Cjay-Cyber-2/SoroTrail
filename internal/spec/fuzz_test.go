package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzParseWasmCustomSections(f *testing.F) {
	f.Add([]byte("\x00asm\x01\x00\x00\x00"))
	f.Add([]byte("random untrusted bytes"))
	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := parseSpecFromWasm(data)
		if err != nil {
			assert.Error(t, err)
		}
	})
}
