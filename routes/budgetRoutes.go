package routes

import (
	"encoding/json"
	"github.com/JhonSmith-1987/presupuesto/Dto"
	"github.com/JhonSmith-1987/presupuesto/db"
	"github.com/JhonSmith-1987/presupuesto/models"
	"net/http"
)

func GetBudget(response http.ResponseWriter, reques *http.Request) {
	var budget []models.Budget
	db.DB.Find(&budget)
	json.NewEncoder(response).Encode(&budget)
}

func PostBudget(response http.ResponseWriter, reques *http.Request) {
	var budget models.Budget
	json.NewDecoder(reques.Body).Decode(&budget)
	create_budget := db.DB.Create(&budget)
	error := create_budget.Error
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Hubo un error no se guardo en la base de datos"))
	}
	res := Dto.ResponseGetUserDto{
		Message: "Valor presupuesto ingresado con éxito",
	}
	json.NewEncoder(response).Encode(&res)
}
