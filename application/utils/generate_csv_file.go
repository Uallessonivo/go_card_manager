package utils

import (
	"bytes"
	"encoding/csv"
	"strings"
)

func GenerateCSVFile(header string, rows [][]string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	columns := strings.Split(header, ",")

	err := writer.Write(columns)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		_ = writer.Write(row)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	defer writer.Flush()

	return buf, nil
}
