package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price string `json:"price" form:"price" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`
}
