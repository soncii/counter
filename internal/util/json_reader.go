package util

import (
	"counter/internal/models"
	"encoding/json"
	"fmt"
	"os"
)

func ReadJSONFile(filePath string) ([]models.Numbers, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var numbers []models.Numbers
	err = json.Unmarshal(file, &numbers)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return numbers, nil
}
