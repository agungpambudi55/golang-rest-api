package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/api/auth"
	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/api/models"
	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/api/responses"
	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/api/utils/formaterror"
	"CRUD-RESTful-API/api/auth"
	"CRUD-RESTful-API/api/models"
	"CRUD-RESTful-API/api/responses"
	"CRUD-RESTful-API/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}