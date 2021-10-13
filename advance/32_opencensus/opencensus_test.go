package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"

	"go.opencensus.io/stats"
)

var (
	// The latency in milliseconds
	MLatencyMs = stats.Float64("repl/latency", "The latency in milliseconds per REPL loop", "ms")

	// Counts/groups the lengths of lines read in.
	MLineLengths = stats.Int64("repl/line_lengths", "The distribution of line lengths", "By")
)

func TestOpenCensus(t *testing.T) {
	// In a REPL:
	//   1. Read input
	//   2. process input
	br := bufio.NewReader(os.Stdin)

	// repl is the read, evaluate, print, loop
	// for {
	if err := readEvaluateProcess(br); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
	// }
}

func readEvaluateProcess(r io.Reader) error {
	return nil
}
