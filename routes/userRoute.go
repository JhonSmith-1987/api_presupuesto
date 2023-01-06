package routes

import (
	"encoding/json"
	"github.com/JhonSmith-1987/presupuesto/Dto"
	"github.com/JhonSmith-1987/presupuesto/db"
	"github.com/JhonSmith-1987/presupuesto/models"
	"net/http"
)

func GetUser(response http.ResponseWriter, reques *http.Request) {
	var users []models.Users
	db.DB.Find(&users)
	res := Dto.ResponseGetUserDto{
		Message:  "Esta son los usuarios",
		Response: users,
	}
	json.NewEncoder(response).Encode(&res)
}

func PostUser(response http.ResponseWriter, reques *http.Request) {
	var user models.Users
	json.NewDecoder(reques.Body).Decode(&user)
	create_user := db.DB.Create(&user)
	error := create_user.Error
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Hubo un error no se guardo en la base de datos"))
	}
	res := Dto.ResponseGetUserDto{
		Message: "Usuario creado con éxito",
	}
	json.NewEncoder(response).Encode(&res)
}

func PostLoginUser(response http.ResponseWriter, reques *http.Request) {
	var user_login Dto.LoginUser
	json.NewDecoder(reques.Body).Decode(&user_login)
	var user models.Users
	db.DB.Raw(`SELECT * FROM users WHERE "user" = ?`, user_login.User).Scan(&user)
	if user.ID == 0 {
		res := Dto.ResponseDtoUser{
			Message: "Usuario no éxiste",
		}
		json.NewEncoder(response).Encode(&res)
		return
	}
	if user.Password != user_login.Password {
		res := Dto.ResponseDtoUser{
			Message: "Contraseña incorrecta",
		}
		json.NewEncoder(response).Encode(&res)
		return
	}
	res := Dto.ResponseDtoUser{
		Message:  "Bienvenido!...",
		Response: user,
	}
	json.NewEncoder(response).Encode(&res)
}
