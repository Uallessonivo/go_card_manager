package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/enums"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/xuri/excelize/v2"
)

func ExtractDataFromExcelFile(file *multipart.FileHeader) ([]*entities.CardRequest, error) {
	if filepath.Ext(file.Filename) != ".xlsx" {
		return nil, errors.FileExtension
	}

	fl, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}

	defer fl.Close()

	rows, err := fl.GetRows("Results")
	if err != nil {
		return nil, err
	}

	cardTypes := map[string]enums.CardType{
		"MATRIZ": enums.DespesasMatriz,
		"FILIAL": enums.DespesasFilial,
	}

	var cards []*entities.CardRequest

	for i, row := range rows {
		if i < 1 {
			continue
		}

		owner := strings.Replace(row[2], ".", "", -1)
		owner = strings.Replace(owner, "-", "", -1)

		cards = append(cards, &entities.CardRequest{
			Type:   cardTypes[row[4]],
			Serial: row[1],
			Owner:  owner,
		})
	}

	return cards, nil
}
