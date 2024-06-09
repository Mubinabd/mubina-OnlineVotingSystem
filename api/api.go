package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"project/api/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(h *handler.HandlerStruct) *gin.Engine {

	r := gin.Default()
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	cand := r.Group("/candidate")
	{
		cand.POST("", h.CreateCandidate)
		cand.GET("/:id", h.GetCandidateByID)
		cand.GET("/all", h.GetAllCandidaties)
		cand.DELETE("/:id", h.DeleteCandidate)
	}

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

	party := r.Group("/party")
	{
		party.POST("", h.CreateParty)
		party.GET("/:id", h.GetByIdParty)
		party.GET("/", h.GetAllParty)
		party.PUT("/:id", h.UpdateParty)
		party.DELETE("/:id", h.DeleteParty)
	}

	public := r.Group("/public")
	{
		public.POST("", h.CreatePublic)
		public.PUT("/:id", h.UpdatePublic)
		public.GET("/id", h.GetByIdPublic)
		public.GET("", h.GetAllPublic)
		public.DELETE("/:id", h.DeletePublic)
	}
	return r
}
