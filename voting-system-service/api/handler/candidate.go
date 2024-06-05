package handler

import (
	"context"
	pb "project/genproto/public"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreateCandidate(c *gin.Context) {
	var candidate pb.CreateCandidateReq
	if err := c.BindJSON(&candidate); err != nil {
		c.JSON(400, gin.H{"Error when binding json": err.Error()})
		return
	}
	_,err := h.Service.CM.Create(context.Background(),&candidate)
	if  err != nil {
		c.JSON(400, gin.H{"Error when creating candidate": err.Error()})
		return
	}
	c.String(201, "candidate created")
}

func (h *HTTPHandler) GetCandidateByID(c *gin.Context) {
	id := c.Query("id")
	election_id := c.Query("election_id")
	public_id := c.Query("public_id")
	candidate, err := h.Service.CM.Get(context.Background(),&pb.GetCandidate{Id: id,ElectionId: election_id,PublicId: public_id})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting candidate": err.Error()})
		return
	}
	c.JSON(200, candidate)
}

func (h *HTTPHandler) GetAllCandidaties(c *gin.Context) {
	candidaties, err := h.Service.CM.GetAll(context.Background(),&pb.Filter{})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting candidaties": err.Error()})
		return
	}
	c.JSON(200, candidaties)
}

func (h *HTTPHandler) DeleteCandidate(c *gin.Context) {
	id := c.Param("id")
	_,err := h.Service.CM.Delete(context.Background(),&pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Error when deleting candidate": err.Error()})
		return
	}
	c.String(200, "candidate deleted")
}
