package productsdto

type CreateProductRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
	//   Password string `json:"password" form:"password" validate:"required"`
}

type UpdateProductRequest struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	//   Password string `json:"password" form:"password"`
}
