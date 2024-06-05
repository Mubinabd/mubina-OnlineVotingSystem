package handler

import (
	"context"
	pb "project/genproto/public"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreatePublicVote(c *gin.Context) {
	var publicVote pb.CreatePublicVoteReq
	if err := c.BindJSON(&publicVote); err != nil {
		c.JSON(400, gin.H{"Error when binding json": err.Error()})
		return
	}
	_,err := h.Service.PVM.Create(context.Background(),&publicVote)
	if err != nil {
		c.JSON(400, gin.H{"Error when creating public Vote": err.Error()})
		return
	}
	c.String(201, "public Vote created")
}

func (h *HTTPHandler) GetPublicVoteByID(c *gin.Context) {
	id := c.Query("id")
	election_id := c.Query("election_id")
	public_id := c.Query("public_id")
	candidate_id := c.Query("candidate_id")
	publicVote, err := h.Service.PVM.Get(context.Background(),&pb.PublicVoteReq{Id: id,ElectionId: election_id,PublicId: public_id,CandidateId: candidate_id})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting public Vote": err.Error()})
		return
	}
	c.JSON(200, publicVote)
}

func (h *HTTPHandler) GetAllPublicVotes(c *gin.Context) {
	publicVote, err := h.Service.PVM.GetAll(context.Background(),&pb.Filter{})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting public Vote": err.Error()})
		return
	}
	c.JSON(200, publicVote)
}


func (h *HTTPHandler) DeletePublicVOte(c *gin.Context) {
	id := c.Param("id")
	_,err := h.Service.PVM.Delete(context.Background(),&pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Error when deleting publicVote": err.Error()})
		return
	}
	c.String(200, "publicVote deleted")
}
