package main
import (

	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (app *App) Initialize() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DbName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	app.Router = mux.NewRouter.StrictSlash(true)
	return nil
}
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func(app *App) handleRoutes() {
	app.Router.HandleFunc("/products", getProducts).Method("GET")
}

/*
Line 11: Char 26: TreeNode.Left undefined (type precompiled.TreeNode has no method Left) (solution.go)
Line 12: Char 27: TreeNode.Right undefined (type precompiled.TreeNode has no method Right) (solution.go)
Line 13: Char 16: invalid operation: root == left + right (mismatched types *precompiled.TreeNode and int) (solution.go)
*/