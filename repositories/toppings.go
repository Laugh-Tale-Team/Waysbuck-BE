package repositories

import (
	"time"
	"waysbuck/models"

	"gorm.io/gorm"
)

type ToppingRepository interface {
	FindToppings() ([]models.Topping, error)
	GetTopping(ID int) (models.Topping, error)
	CreateTopping(topping models.Topping) (models.Topping, error)
	UpdateTopping(topping models.Topping, ID int) (models.Topping, error)
	DeleteTopping(topping models.Topping, ID int) (models.Topping, error)
}

// type repository struct {
// 	db *gorm.DB
// }

func RepositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindToppings() ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Raw("SELECT * FROM toppings").Scan(&toppings).Error

	return toppings, err
}

func (r *repository) GetTopping(ID int) (models.Topping, error) {
	var topping models.Topping
	err := r.db.Raw("SELECT * FROM toppings WHERE id=?", ID).Scan(&topping).Error

	return topping, err
}

func (r *repository) CreateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Exec("INSERT INTO toppings(title,price,created_at,updated_at) VALUES (?,?,?,?)", topping.Title, topping.Price, time.Now(), time.Now()).Error

	return topping, err
}

// Write this code
func (r *repository) UpdateTopping(topping models.Topping, ID int) (models.Topping, error) {
	err := r.db.Raw("UPDATE toppings SET title=?, Price=? WHERE id=?", topping.Title, topping.Price, ID).Scan(&topping).Error

	return topping, err
}

func (r *repository) DeleteTopping(topping models.Topping, ID int) (models.Topping, error) {
	err := r.db.Raw("DELETE FROM toppings WHERE id=?", ID).Scan(&topping).Error

	return topping, err
}
