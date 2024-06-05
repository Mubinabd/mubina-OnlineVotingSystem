package api

import (
	"project/api/handler"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handler.HTTPHandler) *gin.Engine {
	r := gin.Default()

	cand := r.Group("/candidate")
	cand.POST("/", h.CreateCandidate)
	cand.GET("/get", h.GetCandidateByID)
	cand.DELETE("/:id", h.DeleteCandidate)
	r.GET("/products", h.GetAllCandidaties)

	election := r.Group("/election")
	election.POST("/", h.CreateElection)
	election.GET("/:id", h.GetElectionByID)
	election.PUT("/update", h.UpdateElection)
	election.DELETE("/:id", h.DeleteElection)
	r.GET("/elections", h.GetAllElections)

	publicVote := r.Group("/publicvote")
	publicVote.POST("/", h.CreatePublicVote)
	publicVote.GET("/get", h.GetPublicVoteByID)
	publicVote.DELETE("/:id", h.DeletePublicVOte)
	r.GET("/votes", h.GetAllPublicVotes)

	return r
}
