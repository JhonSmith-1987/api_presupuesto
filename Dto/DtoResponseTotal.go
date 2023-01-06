package Dto

import "github.com/JhonSmith-1987/presupuesto/models"

type DtoResponseTotal struct {
	Message    string          `json:"message"`
	SumSpent   uint64          `json:"sumSpent"`
	SumBudget  uint64          `json:"sumBudget"`
	Diferencia uint64          `json:"diferencia"`
	ResSpent   []models.Spent  `json:"resSpent"`
	ResBudget  []models.Budget `json:"resBudget"`
}
