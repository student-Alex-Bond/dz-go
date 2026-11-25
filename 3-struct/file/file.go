package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadMyFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Не могу прочитать файл: %w", err)
	}
	return data, nil
}

func IsJSONPath(path string) bool {
	return strings.EqualFold(filepath.Ext(strings.TrimSpace(path)), ".json")
}
