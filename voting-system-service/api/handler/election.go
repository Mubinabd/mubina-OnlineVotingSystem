package handler

import (
	"context"
	"net/http"
	pb "project/genproto/public"

	"github.com/gin-gonic/gin"  
)

func (h *HTTPHandler) CreateElection(c *gin.Context) {
	var election pb.CreateCandidateReq
	if err := c.ShouldBindJSON(&election); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		return
	}
	_, err := h.Service.EM.Create(context.Background(), &pb.CreateElectionReq{})
	if err != nil {
		c.JSON(400, gin.H{"Error when creating election": err.Error()})
		return
	}
	c.String(201, "Election created")
}

func (h *HTTPHandler) GetElectionByID(c *gin.Context) {
	id := c.Param("id")
	election, err := h.Service.EM.Get(context.Background(), &pb.ElectionReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting election": err.Error()})
		return
	}
	c.JSON(200, election)
}

func (h *HTTPHandler) GetAllElections(c *gin.Context) {
	elections, err := h.Service.EM.GetAll(context.Background(), &pb.Filter{})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting elections": err.Error()})
		return
	}
	c.JSON(200, elections)
}

func (h *HTTPHandler) UpdateElection(c *gin.Context) {
	var election pb.ElectionUpdate
	if err := c.BindJSON(&election); err != nil {
		c.JSON(400, gin.H{"Error when binding json": err.Error()})
		return
	}
	_, err := h.Service.EM.Update(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating election": err.Error()})
		return
	}
	c.String(200, "Election updated")
}

func (h *HTTPHandler) DeleteElection(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Service.EM.Delete(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Error when deleting election": err.Error()})
		return
	}
	c.String(200, "Election deleted")
}
