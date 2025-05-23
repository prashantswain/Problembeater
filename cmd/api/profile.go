package main

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	user "github.com/prashantswain/problem-beater/internal/data"
	"github.com/prashantswain/problem-beater/internal/validator"
)

func (app *Application) createProfileHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailID"`
		MobileNumber string `json:"mobileNumber"`
		Gender       string `json:"gender"`
		Age          int    `json:"age"`
		Password     string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.EmailAddress != "", "Email Address", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Type", "application/json")

	studentId, err := app.db.CreateStudent(input.Name, input.EmailAddress, input.MobileNumber, input.Age, input.Gender, input.Password, time.Now())
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
	}
	slog.Info("Profile Created Successfully!", slog.String("userId", strconv.FormatInt(studentId, 10)))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": studentId, "mesaage": "Profile Created Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *Application) viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	student, err := app.db.GetStudentById(id)
	user := user.Student{
		Id:           student.Id,
		CreatedAt:    student.CreatedAt,
		Name:         student.Name,
		MobileNumber: student.MobileNumber,
		Email_Id:     student.Email_Id,
		Gender:       student.Gender,
		Age:          student.Age,
	}
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
	}
	slog.Info("User Reterived Successfully!", slog.String("user", student.Name))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": user}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
