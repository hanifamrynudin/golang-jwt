package repositories

import (
	"golang-jwt/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemByItemID(itemID int) (models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
	UpdateItem(itemID int, item models.Item) (models.Item, error)
	DeleteItem(item models.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func (r *itemRepository) GetItemByItemID(itemID int) (models.Item, error) {
	var item models.Item

	err := r.db.Where("item_id = ?", itemID).First(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) CreateItem(item models.Item) (models.Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) UpdateItem(itemID int, item models.Item) (models.Item, error) {
	err := r.db.Where("item_id = ?", itemID).Updates(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) DeleteItem(item models.Item) error {
	err := r.db.Delete(item).Error
	if err != nil {
		return err
	}

	return err
}

func ProvideItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}