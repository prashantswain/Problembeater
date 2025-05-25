package main

import (
	"log/slog"
	"net/http"
	"strconv"

	user "github.com/prashantswain/problem-beater/internal/data"
	"github.com/prashantswain/problem-beater/internal/validator"
)

func (app *Application) loginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(credentials.Username != "", "username", "must be provided")
	v.Check(credentials.Password != "", "password", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	userObject, err := app.db.Login(credentials.Username, credentials.Password)
	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err)
		return
	}

	// Generate JWT
	token, err := generateToken(strconv.FormatInt(userObject.Id, 10))

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	err = app.db.UpdateToken(userObject.Id, token)

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	user := user.Student{
		Id:           userObject.Id,
		CreatedAt:    userObject.CreatedAt,
		UpdatedAt:    userObject.UpdatedAt,
		Name:         userObject.Name,
		MobileNumber: userObject.MobileNumber,
		Email_Id:     userObject.Email_Id,
		Gender:       userObject.Gender,
		Age:          userObject.Age,
		ClassId:      userObject.ClassId,
	}

	slog.Info("Logged in Successfully!", slog.String("authToken", token))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": user, "authToken": token, "mesaage": "Logged in Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
