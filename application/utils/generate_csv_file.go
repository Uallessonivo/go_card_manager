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
	writer.Write(columns)

	for _, row := range rows {
		writer.Write(row)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	defer writer.Error()
	defer writer.Flush()

	return buf, nil
}
