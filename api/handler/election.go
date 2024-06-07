package handler

import (
	"fmt"
	"net/http"
	pb "project/genproto/public"

	"github.com/gin-gonic/gin"
)

// CreateElection godoc
// @Summary Create a new election
// @Description Endpoint for creating a new election
// @Name create_election
// @Date create_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body pb.CreateElectionReq true "Election creation request payload"
// @Success 200 {object} pb.Void "Successfully created election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create election"
// @Router /election [POST]
func (h *HandlerStruct) CreateElection(c *gin.Context) {
	var (
		election pb.CreateElectionReq
		err      error
	)

	if err = c.BindJSON(&election); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON: " + err.Error()})
		return
	}

	h.Election.Create(c.Request.Context(), &election)
	fmt.Println(election.Name)

	c.String(http.StatusOK, "Election created successfully")
}

// GetElection godoc
// @Summary Get Election
// @Description Endpoint for getting election
// @Tags Election
// @Accept json
// @Produce json
// @Param  id query pb.GetByIdReq true "ID"
// @Success 200 {object} pb.ElectionRes "Successfully getted election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get election"
// @Router /election/:id [GET]
func (h *HandlerStruct) GetElectionByID(c *gin.Context) {
	id := c.Param("id")
	election, err := h.Election.Get(c.Request.Context(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting election": err.Error()})
		return
	}
	c.JSON(200, election)
}

// GetAllElections godoc
// @Summary Get All Elections
// @Description Endpoint for getting all elections
// @Tags Election
// @Accept json
// @Produce json
// @Param election query pb.Filter true "Elections geting request payload"
// @Success 200 {object} pb.ElectionsGetAllResp "Successfully getted elections"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get elections"
// @Router /election/all [GET]
func (h *HandlerStruct) GetAllElections(c *gin.Context) {
	elections, err := h.Election.GetAll(c.Request.Context(), &pb.Filter{})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting elections": err.Error()})
		return
	}
	c.JSON(200, elections)
}

// UpdateElection godoc
// @Summary Update election
// @Description Endpoint for updating election
// @Name update_election
// @Date update_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body pb.GetByIdReq true "Election updaing request payload"
// @Success 200 {object} pb.Void "Successfully updated election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update election"
// @Router /election/update [PUT]
func (h *HandlerStruct) UpdateElection(c *gin.Context) {
	var election pb.ElectionUpdate
	if err := c.BindJSON(&election); err != nil {
		c.JSON(400, gin.H{"Error when binding json": err.Error()})
		return
	}
	_, err := h.Election.Update(c.Request.Context(), &election)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating election": err.Error()})
		return
	}
	c.JSON(200, "Election updated")
}

// DeleteElection godoc
// @Summary Delete election
// @Description Endpoint for deleting election
// @Name delete_election
// @Date delete_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body pb.GetByIdReq true "Election deleting request payload"
// @Success 200 {object} pb.Void "Successfully deleted election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete election"
// @Router /election/:id [DELETE]
func (h *HandlerStruct) DeleteElection(c *gin.Context) {
	electionReq := pb.GetByIdReq{Id: c.Param("id")}

	_, err := h.Election.Delete(c.Request.Context(), &electionReq)
	if err != nil {
		c.JSON(400, gin.H{"Error when deleting election": err.Error()})
		return
	}
	c.JSON(200, "Election deleted")
}
