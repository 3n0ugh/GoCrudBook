package config

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/3n0ugh/GoCrudBook/pkg/models/mysql"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Books    *mysql.BookModel
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
