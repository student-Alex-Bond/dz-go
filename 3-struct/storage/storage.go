package storage

import (
	"dz/bin-cli/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveJSONToLocal(data *bins.BinList) error {
	if data == nil {
		return fmt.Errorf("нечего сохранять: список равен nil")
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать в JSON: %w", err)
	}
	err = os.WriteFile("data.json", buf, 0644)
	if err != nil {
		return fmt.Errorf("Не удалось записать: %w", err)
	}
	return nil
}

func ReadJSONFile(path string) (*bins.BinList, error) {
	var binList bins.BinList
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Не удалось прочитать файл: %w", err)
	}
	err = json.Unmarshal(buf, &binList)
	if err != nil {
		return nil, fmt.Errorf("Не удалось преобразовать в JSON: %w", err)
	}
	if binList.Bins == nil {
		binList.Bins = make([]bins.Bin, 0)
	}
	return &binList, nil
}
