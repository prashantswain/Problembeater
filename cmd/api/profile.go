package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prashantswain/problem-beater/internal/validator"
)

// Handler for Create Student
func (app *Application) createProfileHandler(w http.ResponseWriter, r *http.Request) {

	// Parse multipart form with 10MB limit
	err := r.ParseMultipartForm(10 << 20) // 10 MB

	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("could not parse multipart form: %w", err))
		return
	}

	// Parse form fields
	name := r.FormValue("name")
	email := r.FormValue("emailID")
	mobileStr := r.FormValue("mobileNumber")
	gender := r.FormValue("gender")
	ageStr := r.FormValue("age")
	password := r.FormValue("password")
	classIDStr := r.FormValue("classId")

	// Convert types
	mobileNumber, _ := strconv.Atoi(mobileStr)
	age, _ := strconv.Atoi(ageStr)
	classId, _ := strconv.ParseInt(classIDStr, 10, 64)

	v := validator.New()
	v.Check(email != "", "emailID", "must be provided")
	v.Check(name != "", "name", "must be provided")
	v.Check(mobileNumber != 0, "mobileNumber", "must be provided")
	v.Check(password != "", "password", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Handle profile picture upload
	var imagePath string
	file, handler, err := r.FormFile("profile_picture")
	if err == nil {
		defer file.Close()

		// Optional: Check file type, size, etc.
		if handler.Size > 5<<20 {
			app.badRequestResponse(w, r, fmt.Errorf("file too large"))
			return
		}

		uploadDir := "uploads"

		// Create uploads directory if it doesn't exist
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err := os.MkdirAll(uploadDir, 0755)
			if err != nil {
				app.serverErrorResponse(w, r, fmt.Errorf("failed to create upload directory: %w", err))
				return
			}
		}

		// Save file with a unique name
		safeFilename := strings.ReplaceAll(handler.Filename, " ", "_")
		imagePath = fmt.Sprintf("%s/%d_%s", uploadDir, time.Now().UnixNano(), safeFilename)
		dst, err := os.Create(imagePath)
		if err != nil {
			app.serverErrorResponse(w, r, fmt.Errorf("unable to save file: %w", err))
			return
		}
		defer dst.Close()
		io.Copy(dst, file)
	}

	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Type", "application/json")
	slog.Info("Saved image", "imagePath", imagePath)
	studentId, err := app.db.CreateStudent(name, email, mobileNumber, age, gender, password, time.Now(), classId, imagePath)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	student, err := app.db.GetStudentById(studentId)

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Profile Created Successfully!", slog.String("userId", strconv.FormatInt(studentId, 10)))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": student, "mesaage": "Profile Created Successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

// Handler for Get Student
func (app *Application) viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	student, err := app.db.GetStudentById(id)

	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	slog.Info("User Reterived Successfully!", slog.String("user", student.Name))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": student, "message": "User read successfully."}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// Handler for Update Student
func (app *Application) updateProfileHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		MobileNumber int    `json:"mobileNumber"`
		Gender       string `json:"gender"`
		Age          int    `json:"age"`
		ClassId      int64  `json:"classId"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	v := validator.New()
	v.Check(input.Id != 0, "id", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	err = app.db.UpdateStudent(input.Id, input.Name, input.MobileNumber, input.Age, input.Gender, input.ClassId, "")

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Student object

	student, err := app.db.GetStudentById(input.Id)

	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	slog.Info("User Reterived Successfully!", slog.String("user", student.Name))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": student, "message": "User read successfully."}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// Handler for DeleteStudent
func (app *Application) deleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = app.db.DeleteStudentById(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Deleted successfully"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
