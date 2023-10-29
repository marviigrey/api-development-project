package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v")
	app.DB = sql.Open("mysql", connectionString)
	var err error
	if err != nil {
		return err
	}

	app.Router := mux.NewRouter().StrictSlash(true)
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_message := map[string]string{"error": err}
	sendResponse(w, statusCode, error_message)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(app.Db)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", getProducts).Methods("GET")
}
