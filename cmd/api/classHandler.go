package main

import (
	"log/slog"
	"net/http"

	models "github.com/prashantswain/problem-beater/internal/models"
	"github.com/prashantswain/problem-beater/internal/validator"
)

func (app *Application) createClassHandler(w http.ResponseWriter, r *http.Request) {

	var classesRequest models.ClassRequest

	err := app.readJSON(w, r, &classesRequest)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	for _, class := range classesRequest.Classes {
		v.Check(class.Name != "", "class_name", "must be provided")
	}
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	err = app.db.CreateClasses(classesRequest.Classes)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mesaage": "Classes created Successfully"}, nil)

}

// Handler for Get Student
func (app *Application) getAllClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := app.db.GetAllClasses()

	if err != nil {
		app.errorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"data": classes, "message": "All classes"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) deleteClass(w http.ResponseWriter, r *http.Request) {
	classId, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = app.db.DeleteClass(classId)

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

func (app *Application) updateClass(w http.ResponseWriter, r *http.Request) {

	var classRequest models.Class

	err := app.readJSON(w, r, &classRequest)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(classRequest.Id != 0, "class_name", "must be provided")
	v.Check(classRequest.Name != "", "class_name", "must be provided")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	class, err := app.db.UpdateClass(classRequest.Id, classRequest.Name)

	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	slog.Info("Class Updated Successfully!", slog.String("user", class.Name))
	err = app.writeJSON(w, http.StatusOK, envelope{"data": class, "message": "Class Updated Successfully!."}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
