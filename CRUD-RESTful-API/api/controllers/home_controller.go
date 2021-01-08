package controllers

import (
	"net/http"

	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/api/responses"
	"CRUD-RESTful-API/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}