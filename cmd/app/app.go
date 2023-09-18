package app

import (
	"log"
	"os"
)

type App struct {
	lg *log.Logger
}

func (a *App) Init() error {
	var err error = nil

	l := log.New(os.Stdout, "", log.LstdFlags)
	a.lg = l

	a.lg.Println("Hello RESTFUL")
	return err
}

func (a *App) Run() {
	defer a.lg.Println("Server is closed")
	a.lg.Println("Server started")
}
