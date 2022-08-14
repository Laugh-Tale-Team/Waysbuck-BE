package toppingsdto

type CreateToppingRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price string `json:"price" form:"price" validate:"required"`
	//   Password string `json:"password" form:"password" validate:"required"`
}

type UpdateToppingRequest struct {
	Title string `json:"title" form:"title"`
	Price string `json:"price" form:"price"`
	//   Password string `json:"password" form:"password"`
}
