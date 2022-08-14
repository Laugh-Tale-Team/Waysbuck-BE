package routes

//import gorilla mux
import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	ProductRoutes(r)
	ToppingRoutes(r)
	// UserRoutes(r)
}
