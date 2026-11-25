package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PathName struct {
	fileName string
}

func NewPathName(name string) *PathName {
	return &PathName{
		fileName: name,
	}
}

func (fl *PathName) ReadMyFile() ([]byte, error) {
	data, err := os.ReadFile(fl.fileName)
	if err != nil {
		return nil, fmt.Errorf("Не могу прочитать файл: %w", err)
	}
	return data, nil
}

func (fl *PathName) IsJSONPath() bool {
	return strings.EqualFold(filepath.Ext(strings.TrimSpace(fl.fileName)), ".json")
}
