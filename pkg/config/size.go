package config

import (
	"fmt"
	"strconv"
	"strings"
)

type DataSize struct {
	Bytes uint64
}

func (ds *DataSize) UnmarshalText(text []byte) error {
	str := string(text)
	unit := strings.ToUpper(str[len(str)-2:])                  // Получение единицы измерения
	numPart, err := strconv.ParseInt(str[:len(str)-2], 10, 64) // Получение числовой части
	if err != nil {
		return err
	}

	switch unit {
	case "B":
		ds.Bytes = uint64(numPart)
	case "KB":
		ds.Bytes = uint64(numPart) * 1024
	case "MB":
		ds.Bytes = uint64(numPart) * 1024 * 1024
	case "GB":
		ds.Bytes = uint64(numPart) * 1024 * 1024 * 1024
	default:
		return fmt.Errorf("unknown size unit: %s", unit)
	}

	return nil
}
