package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Raw("SELECT * FROM products").Scan(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Raw("SELECT * FROM products WHERE id=?", ID).Scan(&product).Error

	return product, err
}
