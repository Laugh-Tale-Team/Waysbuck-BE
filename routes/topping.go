package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/middleware"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", middleware.Auth(h.FindToppings)).Methods("GET")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.GetTopping)).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(h.CreateTopping)).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.UpdateTopping)).Methods("PATCH")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.DeleteTopping)).Methods("DELETE")
}
