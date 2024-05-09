package memory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkDataReader(b *testing.B) {
	process, err := NewProcess()
	require.NoError(b, err)

	gr := NewGameReader(process)
	gr.GetData()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gr.GetData()
	}
}
