package middlewares

import (
	"errors"
	"net/http"

	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/fullstack/api/auth"
	// "github.com/agungpambudi55/Learn-Golang/tree/master/CRUD-RESTful-API/fullstack/api/responses"
	"CRUD-RESTful-API/fullstack/api/auth"
	"CRUD-RESTful-API/fullstack/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}