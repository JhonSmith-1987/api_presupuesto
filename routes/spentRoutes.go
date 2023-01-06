package routes

import (
	"encoding/json"
	"github.com/JhonSmith-1987/presupuesto/Dto"
	"github.com/JhonSmith-1987/presupuesto/db"
	"github.com/JhonSmith-1987/presupuesto/models"
	"net/http"
)

func GetSpent(response http.ResponseWriter, reques *http.Request) {
	var spent []models.Spent
	var budget []models.Budget
	var suma_budget uint64 = 0
	var suma_spent uint64 = 0
	db.DB.Find(&spent)
	for i := 0; i < len(spent); i++ {
		suma_spent += spent[i].Spent
	}
	db.DB.Find(&budget)
	for i := 0; i < len(budget); i++ {
		suma_budget += budget[i].Budget
	}
	var diferencia uint64 = suma_budget - suma_spent
	res := Dto.DtoResponseTotal{
		SumSpent:   suma_spent,
		SumBudget:  suma_budget,
		Diferencia: diferencia,
		ResBudget:  budget,
		ResSpent:   spent,
	}
	json.NewEncoder(response).Encode(&res)
}

func PostSpent(response http.ResponseWriter, reques *http.Request) {
	var spent_post models.Spent
	json.NewDecoder(reques.Body).Decode(&spent_post)
	create_spent := db.DB.Create(&spent_post)
	error := create_spent.Error
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Hubo un error no se guardo en la base de datos"))
	}
	var spent []models.Spent
	var budget []models.Budget
	var suma_budget uint64 = 0
	var suma_spent uint64 = 0
	db.DB.Find(&spent)
	for i := 0; i < len(spent); i++ {
		suma_spent += spent[i].Spent
	}
	db.DB.Find(&budget)
	for i := 0; i < len(budget); i++ {
		suma_budget += budget[i].Budget
	}
	var diferencia uint64 = suma_budget - suma_spent
	res := Dto.DtoResponseTotal{
		Message:    "Gasto guardado con éxito",
		SumSpent:   suma_spent,
		SumBudget:  suma_budget,
		Diferencia: diferencia,
		ResBudget:  budget,
		ResSpent:   spent,
	}
	json.NewEncoder(response).Encode(&res)
}
