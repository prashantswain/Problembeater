package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/prashantswain/problem-beater/internal/models"
	"github.com/prashantswain/problem-beater/internal/validator"
)

func (app *Application) loginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials models.LoginRequest

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
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	// Generate JWT
	time, token, err := app.SetToken(userObject.Id)

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	userObject.UpdatedAt = time

	slog.Info("Logged in Successfully!", slog.String("authToken", *token))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": userObject, "authToken": &token, "mesaage": "Logged in Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// Forgot Password
func (app *Application) logoutUser(w http.ResponseWriter, r *http.Request) {
	id, err := app.authenticate(r, w)
	if err != nil {
		return
	}

	err = app.db.Logout(id)
	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mesaage": "Logout Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

// Forgot Password
func (app *Application) forgotPassword(w http.ResponseWriter, r *http.Request) {
	var request models.ForgotPasswordRequest

	err := app.readJSON(w, r, &request)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(request.Email != "", "email", "must be provided")
	v.Check(request.Password != "", "password", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.db.ForgotPassword(request.Email, request.Password)
	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mesaage": "Password changed Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *Application) SetToken(userId int64) (*time.Time, *string, error) {
	// Generate JWT
	token, err := GenerateToken(strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, nil, err
	}

	time, err := app.db.UpdateToken(userId, token)

	if err != nil {
		return nil, nil, err
	}

	return time, &token, nil
}

func (app *Application) ValidateSession(token string) error {
	userIdString, err := ValidateToken(token)
	if err != nil {
		return err
	}
	// Convert string to int64
	userID, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return err
	}
	err = app.db.ValidateSession(userID, token)

	if err != nil {
		return err
	}
	return nil
}
