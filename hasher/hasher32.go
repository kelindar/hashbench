package hasher

import (
	"hash"
	"hash/crc32"

	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-marvin32"
	"github.com/mtchavez/jenkins"
	"github.com/twmb/murmur3"
)

type hasher32 = (func([]byte) uint32)

// For32 loads a 32-bit hasher by name
func For32(name string) (hasher32, bool) {
	switch name {
	case "crc":
		return crc32.ChecksumIEEE, true
	case "twmb/murmur3":
		return murmur3.Sum32, true
	case "mtchavez/jenkins":
		h := jenkins.New()
		return hasher32For(h), true
	case "dgryski/go-marvin32":
		return func(b []byte) uint32 {
			return marvin32.Sum32(0, b)
		}, true
	case "dgryski/go-farm":
		return farm.Hash32, true
	}

	return nil, false
}

// hasher32For returns a hasher for the default interface
func hasher32For(h hash.Hash32) hasher32 {
	return func(b []byte) uint32 {
		h.Reset()
		h.Write(b)
		return h.Sum32()
	}
}
