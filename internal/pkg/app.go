package app

import (
	"cityService/internal/app/endpoints"
	"cityService/internal/app/repositories"
	"cityService/internal/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	e   *endpoints.Endpoint
	s   *services.Service
	r   *repositories.Repository
	gin *gin.Engine
}

func New() (*App, error) {
	a := &App{}

	a.r = repositories.New()

	a.s = services.New(a.r)

	a.e = endpoints.New(a.s)

	a.gin = gin.New()

	a.gin.GET("/cities", a.e.Cities)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")
	err := a.gin.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
