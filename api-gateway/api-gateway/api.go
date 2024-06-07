package api

import (
	"project/api-gateway/handler"
	_ "project/docs"

	// "project/api-gateway/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name role

func NewGin(h *handler.HandlerStruct) *gin.Engine {
	r := gin.Default()

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	cand := r.Group("/candidate")
	{
		cand.POST("", h.CreateCandidate)
		cand.GET("/:id", h.GetCandidateByID)
		cand.GET("/all", h.GetAllCandidaties)
		cand.DELETE("/:id", h.DeleteCandidate)
	}

	// r.Use(middleware.AuthMiddleware)

	election := r.Group("/election")
	{
		election.POST("", h.CreateElection)
		election.GET("/:id", h.GetElectionByID)
		election.GET("/all", h.GetAllElections)
		election.PUT("/update", h.UpdateElection)
		election.DELETE("/:id", h.DeleteElection)
	}

	publicVote := r.Group("/publicvote")
	{
		publicVote.POST("", h.CreatePublicVote)
		publicVote.GET("/all", h.GetAllPublicVotes)
	}
	return r
}
