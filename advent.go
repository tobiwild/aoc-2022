package advent

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type Test[T any] struct {
	File     string
	Expected T
}

func TestSolve[T any](t *testing.T, solve func(io.Reader) (T, error), tests []Test[T]) {
	for _, test := range tests {
		t.Run(test.File, func(t *testing.T) {
			file, err := os.Open(test.File)
			require.NoError(t, err)
			defer file.Close()
			result, err := solve(file)
			require.NoError(t, err)
			t.Logf("RESULT: %+v", result)
			require.Equal(t, test.Expected, result)
		})
	}
}

func BenchmarkSolve[T any](b *testing.B, solve func(io.Reader) (T, error), files []string) {
	for _, file := range files {
		// Read whole file to not benchmark IO performance
		content, err := os.ReadFile(file)
		require.NoError(b, err)
		b.Run(file, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := bytes.NewReader(content)
				_, err := solve(reader)
				require.NoError(b, err)
			}
		})
	}
}
