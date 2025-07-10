package storage

import (
	"fmt"
	"os"
	"strings"
)

const dbFileName = "data.db"

// Set writes a key=value pair to the file
func (idx *Index) Set(key, value string) error {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	if idx.file == nil {
		// Create file if it doesn't exist
		var err error
		idx.file, err = os.OpenFile(dbFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
	}

	// Seek to end before writing
	offset, err := idx.file.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	_, err = idx.file.WriteString(fmt.Sprintf("%s=%s\n", key, value))

	if err != nil {
		return err
	}

	idx.offsets[key] = offset
	return nil
}

func (idx *Index) Get(key string) (string, error) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	offset, ok := idx.offsets[key]
	if !ok {
		return "", os.ErrNotExist
	}

	buf := make([]byte, 1024) // enough buffer for one line

	_, err := idx.file.Seek(offset, 0)
	if err != nil {
		return "", err
	}

	n, err := idx.file.Read(buf)
	if err != nil {
		return "", err
	}

	line := string(buf[:n])
	if i := strings.Index(line, "="); i >= 0 {
		return strings.TrimSpace(line[i+1:]), nil
	}
	return "", fmt.Errorf("malformed entry")
}

// Helper: splits string into lines
func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, ch := range s {
		if ch == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

// Helper: parses "key=value" lines
func parseLine(line string) (key, value string, ok bool) {
	for i := 0; i < len(line); i++ {
		if line[i] == '=' {
			return line[:i], line[i+1:], true
		}
	}
	return "", "", false
}
