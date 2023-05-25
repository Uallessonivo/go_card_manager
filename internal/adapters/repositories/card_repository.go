package repositories

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"gorm.io/gorm"
)

type CardRepository struct {
	Db *gorm.DB
}

func NewCardRepository(Db *gorm.DB) ports.CardRepository {
	return &CardRepository{Db}
}

func (c *CardRepository) Create(input *models.Card) error {
	err := c.Db.Create(&input).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CardRepository) List() ([]*models.Card, error) {
	var cards []*models.Card

	err := c.Db.Find(&cards).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) ListByTYpe(input string) ([]*models.Card, error) {
	var cards []*models.Card

	err := c.Db.Find(&cards, "type = ?", input).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) ListByOwner(input string) ([]*models.Card, error) {
	var cards []*models.Card

	err := c.Db.Find(&cards, "owner = ?", input).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) GetByOwner(input string) (*models.Card, error) {
	var card *models.Card

	err := c.Db.First(&card, "owner = ?", input).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c *CardRepository) GetById(input string) (*models.Card, error) {
	var card *models.Card

	err := c.Db.First(&card, "id = ?", input).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (c *CardRepository) Delete(id string) error {
	err := c.Db.Delete(&models.Card{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
