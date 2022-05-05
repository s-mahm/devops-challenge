package main

import (
	"github.com/badkaktus/gorocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Client *gorocket.Client
	Logger *log.Logger
}

func (a *App) Initialize(url string) {
	a.Logger = log.New(os.Stdout, "", log.LstdFlags)
	a.Client = gorocket.NewClient(url)
	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) Run(addr string) {
	loggedRouter := a.createLoggingRouter(a.Logger.Writer())
	a.Logger.Fatal(http.ListenAndServe(addr, loggedRouter))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/commit/{hash}", a.sendCommitHash).Methods("POST")
}
