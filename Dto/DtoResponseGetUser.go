package Dto

import "github.com/JhonSmith-1987/presupuesto/models"

type ResponseGetUserDto struct {
	Message  string         `json:"message"`
	Response []models.Users `json:"response"`
}

type ResponseDtoUser struct {
	Message  string       `json:"message"`
	Response models.Users `json:"response"`
}
