package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PathName struct {
	pathName string
}

func NewPathName(name string) *PathName {
	return &PathName{
		pathName: name,
	}
}

func (pn *PathName) ReadMyFile() ([]byte, error) {
	data, err := os.ReadFile(pn.pathName)
	if err != nil {
		return nil, fmt.Errorf("Не могу прочитать файл: %w", err)
	}
	return data, nil
}

func (pn *PathName) IsJSONPath() bool {
	return strings.EqualFold(filepath.Ext(strings.TrimSpace(pn.pathName)), ".json")
}
