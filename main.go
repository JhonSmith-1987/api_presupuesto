package main

import (
	"github.com/JhonSmith-1987/presupuesto/db"
	"github.com/JhonSmith-1987/presupuesto/models"
	"github.com/JhonSmith-1987/presupuesto/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Users{})
	db.DB.AutoMigrate(models.Budget{})
	db.DB.AutoMigrate(models.Spent{})

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	// route index
	router.HandleFunc("/", routes.HomeRoutes)

	// routes users
	api.HandleFunc("/users", routes.GetUser).Methods("GET")
	api.HandleFunc("/user", routes.PostUser).Methods("POST")
	api.HandleFunc("/userLogin", routes.PostLoginUser).Methods("POST")

	// routes budget
	api.HandleFunc("/budget", routes.GetBudget).Methods("GET")
	api.HandleFunc("/budgets", routes.PostBudget).Methods("POST")

	// routes spent
	api.HandleFunc("/spent", routes.GetSpent).Methods("GET")
	api.HandleFunc("/spentPost", routes.PostSpent).Methods("POST")

	// port init
	http.ListenAndServe(":3000", router)

}
