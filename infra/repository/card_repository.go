package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"gorm.io/gorm"
)

type CardRepository struct {
	Db *gorm.DB
}

func NewCardRepository(Db *gorm.DB) interfaces.CardRepositoryInterface {
	return &CardRepository{Db}
}

func (c *CardRepository) Create(input *entities.Card) error {
	err := c.Db.Create(&input).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CardRepository) List() ([]*entities.Card, error) {
	var cards []*entities.Card

	err := c.Db.Find(&cards).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) ListByTYpe(input string) ([]*entities.Card, error) {
	var cards []*entities.Card

	err := c.Db.Find(&cards, "type = ?", input).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) ListByOwner(input string) ([]*entities.Card, error) {
	var cards []*entities.Card

	err := c.Db.Find(&cards, "owner = ?", input).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) GetByOwner(input string) (*entities.Card, error) {
	var card *entities.Card

	err := c.Db.First(&card, "owner = ?", input).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c *CardRepository) GetById(input string) (*entities.Card, error) {
	var card *entities.Card

	err := c.Db.First(&card, "id = ?", input).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c *CardRepository) Delete(id string) error {
	err := c.Db.Delete(&entities.Card{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
