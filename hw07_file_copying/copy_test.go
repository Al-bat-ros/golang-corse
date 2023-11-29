package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	text := "Tell me, uncle, it s not for nothing Moscow, burned by fire, Given to the Frenchman?"
	fileIn := "in_file"
	fileOut := "out_file"

	tests := []struct {
		name   string
		offset int64
		limit  int64
		input  string
		result string
	}{
		{"equals", 0, 0, text, text},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			from, err := os.Create(fileIn)
			if err != nil {
				log.Fatal(err)
			}

			defer os.Remove(fileIn)
			_, err = from.WriteString(test.input)
			if err != nil {
				log.Fatal(err)
			}
			defer from.Close()

			err = Copy(fileIn, fileOut, test.offset, test.limit)

			if err != nil {
				log.Fatal(err)
			}

			defer os.Remove(fileOut)
			result, err := os.ReadFile(fileOut)
			if err != nil {
				log.Fatal(err)
			}
			require.Equal(t, string(result), test.result)
		})
	}
}
