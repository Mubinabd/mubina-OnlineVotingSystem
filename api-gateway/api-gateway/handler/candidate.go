package handler

import (
	"context"
	"fmt"
	"net/http"

	pb "project/genproto/public"

	"github.com/gin-gonic/gin"
)

// CreateCandidate godoc
// @Summary Create a new candidate
// @Description Endpoint for creating a new candidate
// @Name create_candidate
// @Date create_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body pb.CreateCandidateReq true "Candidate creation request payload"
// @Success 200 {object} pb.Void "Successfully created candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate [POST]
func (h *HandlerStruct) CreateCandidate(c *gin.Context) {
	var (
		candidate pb.CreateCandidateReq
		err       error
	)
	if err = c.BindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON: " + err.Error()})
		return
	}
	h.Candidate.Create(context.Background(), &candidate)
	c.JSON(http.StatusOK, "Candidate created successfully")
}

// GetCandidate godoc
// @Summary Get Candidate
// @Description Endpoint for getting candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param  id query pb.GetByIdReq true "ID"
// @Success 200 {object} pb.CandidateRes "Successfully getted candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get candidate"
// @Router /candidate/:id [GET]
func (h *HandlerStruct) GetCandidateByID(c *gin.Context) {
	id := c.Param("id")
	candidate, err := h.Candidate.Get(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting candidate: " + err.Error()})
		return
	}
	fmt.Println(1)

	c.IndentedJSON(http.StatusOK, candidate)
}

// GetAllCandidate godoc
// @Summary Get All Candidates
// @Description Endpoint for getting all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate query pb.Filter true "Candidates geting request payload"
// @Success 200 {object} pb.CandidatiesGetAllResp "Successfully getted candidates"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get candidates"
// @Router /candidate/all [GET]
func (h *HandlerStruct) GetAllCandidaties(c *gin.Context) {
	var filter pb.Filter

	if err := c.BindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding query: " + err.Error()})
		return
	}

	candidates, err := h.Candidate.GetAll(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting candidates: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, candidates)
}

// DeleteCandidate godoc
// @Summary Delete candidate
// @Description Endpoint for deleting candidate
// @Name delete_candidate
// @Date delete_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body pb.GetByIdReq true "candidate deleting request payload"
// @Success 200 {object} pb.Void "Successfully deleted candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete candidate"
// @Router /candidate/:id [DELETE]

func (h *HandlerStruct) DeleteCandidate(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Candidate.Delete(c.Request.Context(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when deleting candidate: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Candidate deleted")
}
