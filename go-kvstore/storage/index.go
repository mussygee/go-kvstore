package storage

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type Index struct {
	mu      sync.RWMutex
	offsets map[string]int64
	file    *os.File
}

// loadIndex reads the db file once and builds the offsets map
func LoadIndex(filename string) (*Index, error) {
	file, err := os.Open(filename)
	if err != nil {
		// if file doesn't exist yet, create empty index with nil file
		if os.IsNotExist(err) {
			return &Index{
				offsets: make(map[string]int64),
				file:    nil,
			}, nil
		}
		return nil, err
	}

	idx := &Index{
		offsets: make(map[string]int64),
		file:    file,
	}

	scanner := bufio.NewScanner(file)
	var offset int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "="); i >= 0 {
			key := line[:i]
			idx.offsets[key] = offset
		}
		offset += int64(len(line)) + 1 // +1 for new line
	}

	return idx, scanner.Err()
}
