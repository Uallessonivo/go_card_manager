package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/enums"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/xuri/excelize/v2"
)

// WIP: extract data from file
func ExtractDataFromExcelFile(file *excelize.File) ([]*model.CardRequest, error) {
	var cards []*model.CardRequest

	rows, err := file.GetRows("Results")
	if err != nil {
		return nil, err
	}

	cardTypes := map[string]enums.CardType{
		"MATRIZ": enums.DespesasMatriz,
		"FILIAL": enums.DespesasFilial,
	}

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
