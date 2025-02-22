package utils

import (
	"encoding/binary"
	"math"
	"unicode/utf16"

	"github.com/hectorgimenez/d2go/pkg/data"
)

func DistanceFromPoint(from data.Position, to data.Position) int {
	first := math.Pow(float64(to.X-from.X), 2)
	second := math.Pow(float64(to.Y-from.Y), 2)

	return int(math.Sqrt(first + second))
}
func DecodeText(value uint32) string {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)

	// Try ASCII
	if isASCIIPrintable(bytes) {
		return string(bytes)
	}

	// Try UTF-16
	if len(bytes) >= 2 {
		utf16Bytes := make([]uint16, 2)
		utf16Bytes[0] = uint16(bytes[0]) | (uint16(bytes[1]) << 8)
		if bytes[2] != 0 {
			utf16Bytes[1] = uint16(bytes[2]) | (uint16(bytes[3]) << 8)
		}
		return string(utf16.Decode(utf16Bytes))
	}

	return ""
}

func detectEncoding(value uint32) string {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)

	// Check for UTF-16 pattern (every other byte zero)
	if bytes[1] == 0 && bytes[3] == 0 {
		return "UTF-16LE"
	}

	// Check for ASCII pattern
	if isASCIIPrintable(bytes) {
		return "ASCII"
	}

	return "unknown"
}

func isASCIIPrintable(bytes []byte) bool {
	for _, b := range bytes {
		if b != 0 && (b < 32 || b > 126) {
			return false
		}
	}
	return true
}
