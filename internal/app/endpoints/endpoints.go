package endpoints

import (
	city "cityService/internal/app/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service interface {
	GetCities() city.Cities
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e Endpoint) Cities(ctx *gin.Context) {
	c := e.s.GetCities()
	ctx.JSON(http.StatusOK, c)
}
