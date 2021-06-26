package hasher

import (
	"hash"
	"hash/crc64"
	"hash/fnv"
	"hash/maphash"

	"github.com/cespare/xxhash/v2"
	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-metro"
	"github.com/dgryski/go-sip13"
	"github.com/dgryski/tsip/go"
	"github.com/minio/highwayhash"
	md5simd "github.com/minio/md5-simd"
	"github.com/twmb/murmur3"
)

var md5 = md5simd.NewServer()

type hasher64 = (func([]byte) uint64)

// For64 loads a 64-bit hasher by name
func For64(name string) (hasher64, bool) {
	switch name {
	case "cespare/xxhash":
		return xxhash.Sum64, true
	case "crc":
		return hasher64For(crc64.New(crc64.MakeTable(crc64.ISO))), true
	case "fnv1":
		return hasher64For(fnv.New64()), true
	case "fnv1a":
		return hasher64For(fnv.New64a()), true
	case "dgryski/go-metro":
		return func(b []byte) uint64 {
			return metro.Hash64(b, 0)
		}, true
	case "twmb/murmur3":
		return murmur3.Sum64, true
	case "minio/highwayhash":
		key := make([]byte, 32)
		return func(b []byte) uint64 {
			return highwayhash.Sum64(b, key)
		}, true
	case "maphash":
		var h maphash.Hash
		return hasher64For(&h), true
	case "dgryski/go-sip13":
		return func(b []byte) uint64 {
			return sip13.Sum64(0, 0, b)
		}, true
	case "dgryski/tsip":
		return func(b []byte) uint64 {
			return tsip.HashASM(0, 0, b)
		}, true
	case "dgryski/go-farm":
		return farm.Hash64, true
	}

	return nil, false
}

// hasher64For returns a hasher for the default interface
func hasher64For(h hash.Hash64) hasher64 {
	return func(b []byte) uint64 {
		h.Reset()
		h.Write(b)
		return h.Sum64()
	}
}
