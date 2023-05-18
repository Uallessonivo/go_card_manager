package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/Uallessonivo/go_card_manager/domain/enums"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/xuri/excelize/v2"
)

func ExtractDataFromExcelFile(file *multipart.FileHeader) ([]*model.CardRequest, error) {
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

	var cards []*model.CardRequest

	for i, row := range rows {
		if i < 1 {
			continue
		}

		cards = append(cards, &model.CardRequest{
			Type:   cardTypes[row[4]],
			Serial: row[1],
			Owner:  row[2],
		})
		println(row)
	}

	return cards, nil
}
