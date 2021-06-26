package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/kelindar/hashbench/hasher"
	"github.com/olekukonko/tablewriter"
)

func main() {
	algorithms := []string{
		"crc",
		"fnv1",
		"fnv1a",
		"maphash",
		"dgryski/go-marvin32",
		"mtchavez/jenkins",
		"cespare/xxhash",
		"dgryski/go-metro",
		"twmb/murmur3",
		"minio/highwayhash",
		"dgryski/go-sip13",
		"dgryski/tsip",
		"dgryski/go-farm",
	}

	// Generate test data
	sizes := []int{10, 25, 100, 4e3, 10e6, 50e6}
	//sizes := []int{10, 100, 4e3}
	input := make([][]byte, 0, len(sizes))
	for _, s := range sizes {
		buffer := make([]byte, s)
		input = append(input, buffer)
		rand.Read(buffer)
	}

	// Create the output
	var out [][]string

	// Run 32-bit hashes
	fmt.Printf("running 32-bit")
	for _, algo := range algorithms {
		for _, buffer := range input {
			out = append(out, run32(algo, buffer))
		}
	}

	// Run 64-bit hashes
	fmt.Printf("\nrunning 64-bit")
	for _, algo := range algorithms {
		for _, buffer := range input {
			out = append(out, run64(algo, buffer))
		}
	}

	// Print the result
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.SetBorder(false)
	table.AppendBulk(out)
	table.Render()

}

func run32(name string, buffer []byte) []string {
	size := humanize.Bytes(uint64(len(buffer)))
	hasher, ok := hasher.For32(name)
	if !ok {
		return nil
	}

	result := testing.Benchmark(func(b *testing.B) {
		b.SetBytes(int64(len(buffer)))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hasher(buffer)
		}
	})

	fmt.Printf(".")
	return []string{"32-bit", name, size, nsPerOp(result), mbPerSec(result)}
}

func run64(name string, buffer []byte) []string {
	size := humanize.Bytes(uint64(len(buffer)))
	hasher, ok := hasher.For64(name)
	if !ok {
		return nil
	}

	result := testing.Benchmark(func(b *testing.B) {
		b.SetBytes(int64(len(buffer)))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hasher(buffer)
		}
	})

	fmt.Printf(".")
	return []string{"64-bit", name, size, nsPerOp(result), mbPerSec(result)}
}

// mbPerSec returns the "MB/s" metric.
func mbPerSec(r testing.BenchmarkResult) string {
	return fmt.Sprintf("%7.2f MB/s",
		(float64(r.Bytes)*float64(r.N)/1e6)/r.T.Seconds(),
	)
}

// NsPerOp returns the "ns/op" metric.
func nsPerOp(r testing.BenchmarkResult) string {
	return prettyPrint(float64(r.T.Nanoseconds())/float64(r.N), "ns/op")
}

func prettyPrint(x float64, unit string) string {
	var format string
	switch y := math.Abs(x); {
	case y == 0 || y >= 999.95:
		format = "%10.0f %s"
	case y >= 99.995:
		format = "%12.1f %s"
	case y >= 9.9995:
		format = "%13.2f %s"
	case y >= 0.99995:
		format = "%14.3f %s"
	case y >= 0.099995:
		format = "%15.4f %s"
	case y >= 0.0099995:
		format = "%16.5f %s"
	case y >= 0.00099995:
		format = "%17.6f %s"
	default:
		format = "%18.7f %s"
	}
	return fmt.Sprintf(format, x, unit)
}
