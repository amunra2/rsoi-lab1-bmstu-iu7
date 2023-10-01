package handler

import (
	"persserv/docs"
	"persserv/internal/usecase"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	API_TAG    = "/api/v1"
	PERSON_TAG = "/persons"
)

type Handler struct {
	useCases *usecase.UseCase
}

func NewHandler(useCases *usecase.UseCase) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// SWAGGER
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group(API_TAG)
	{
		person := api.Group(PERSON_TAG)
		{
			person.GET("/", h.getAllPersons)
			person.GET("/:id", h.getByIdPerson)
			person.POST("/", h.createPerson)
			person.PATCH("/:id", h.updatePerson)
			person.DELETE("/:id", h.deletePerson)
		}
	}

	return router
}
