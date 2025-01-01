package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func SaveFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(
		fileName,
		[]byte(valueText),
		0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil { // we got error, nil mean absence of anything useful
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
