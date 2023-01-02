package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type CardRepository struct {
	Db *gorm.DB
}

func NewCardRepository(Db *gorm.DB) interfaces.CardRepositoryInterface {
	return &CardRepository{Db}
}

func (c *CardRepository) Create(input *model.Card) error {
	err := c.Db.Create(&input).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CardRepository) List() ([]*model.Card, error) {
	var cards []*model.Card

	err := c.Db.Find(&cards).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) GetByCpf(cpf string) (*model.Card, error) {
	var cards *model.Card

	err := c.Db.First(&cards, "cpf = ?", cpf).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) ListByType(cardType string) ([]*model.Card, error) {
	var cards []*model.Card

	err := c.Db.Find(&cards, "type = ?", cardType).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) Delete(id string) error {
	err := c.Db.Delete(&model.Card{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
