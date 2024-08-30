package Groupie_tracker

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type ArtistWithLocation struct {
	JsonData interface{}
}

var (
	tmpl   *template.Template
	errors AllMessageErrors
)

// Initialize the global template variable
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	errors = ErrorsMessage()
}

// This function is responsible for handling the root path ("/") of the application.
func GetDataFromJson(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	HandleErrors(w, errors.MethodNotAllowed, errors.DescriptionMethodNotAllowed, http.StatusMethodNotAllowed)
	// 	return
	// }
	if r.URL.Path != "/" {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}

	artisData, errs := GetArtistsDataStruct()
	if errs != nil {
		HandleErrors(w, errors.BadRequest, errors.DescriptionBadRequest, http.StatusBadRequest)
		return
	}
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "index.html", artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	_, erro := buf.WriteTo(w)
	if erro != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
}

// This function is responsible for handling the individual artist's information page.
func HandlerShowRelation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErrors(w, errors.MethodNotAllowed, errors.DescriptionMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	idParam := r.PathValue("id")
	artist, err := FetchDataRelationFromId(idParam)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	if artist.ID == 0 {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	errs := tmpl.ExecuteTemplate(&buf, "InforArtis.html", artist)
	if errs != nil {

		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	_, erro := buf.WriteTo(w)
	if erro != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
}

// This function is responsible for serving the CSS files for the application.
func HandleStyle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/styles"):]

	fullpath := filepath.Join("src", path)
	fileinfo, err := os.Stat(fullpath)
	if err != nil {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}

	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.StripPrefix("/styles", http.FileServer(http.Dir("src"))).ServeHTTP(w, r)
	} else {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}
}

// This function is responsible for handling and displaying error messages.
func HandleErrors(w http.ResponseWriter, message, description string, code int) {
	errorsMessage := Errors{
		Message:     message,
		Description: description,
		Code:        code,
	}
	w.WriteHeader(code)
	err := tmpl.ExecuteTemplate(w, "errors.html", errorsMessage)
	if err != nil {
		http.Error(w, "Error 500 Internal Server Error", http.StatusInternalServerError)
	}
}
