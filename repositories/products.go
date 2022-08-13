package repositories

import (
	"time"
	"waysbuck/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product, ID int) (models.Product, error)
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

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Exec("INSERT INTO users(title,price,created_at,updated_at) VALUES (?,?,?,?)", product.Title, product.Price, time.Now(), time.Now()).Error

	return product, err
}

// Write this code
func (r *repository) UpdateProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Raw("UPDATE Products SET title=?, Price=? WHERE id=?", product.Title, product.Price, ID).Scan(&product).Error

	return product, err
}
