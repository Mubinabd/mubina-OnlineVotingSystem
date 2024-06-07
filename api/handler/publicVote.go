package handler

import (
	"fmt"
	"net/http"
	public "project/genproto/public"

	"github.com/gin-gonic/gin"
)

// CreatePublicVote godoc
// @Summary Create a new publicvote
// @Description Endpoint for creating a new publicvote
// @Name create_publicvote
// @Date create_publicvote
// @Tags Public Vote
// @Accept json
// @Produce json
// @Param publicvote body public.CreatePublicVoteReq true "publicVote creation request payload"
// @Success 200 {object} public.Void "Successfully created publicvote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create publicvote"
// @Router /publicvote [POST]
func (h *HandlerStruct) CreatePublicVote(c *gin.Context) {
	var (
		publicvote public.CreatePublicVoteReq
		err        error
	)
	if err = c.BindJSON(&publicvote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON: " + err.Error()})
		return
	}

	h.PublicVote.Create(c.Request.Context(), &publicvote)
	fmt.Println(publicvote.ElectionId)

	c.JSON(http.StatusOK, "Public Vote created successfully")
}

// GetPublicVote godoc
// @Summary Get All Public Votes
// @Description Endpoint for getting all public votes
// @Tags Public Vote
// @Accept json
// @Produce json
// @Param publicVote query public.Filter true "Public Votes geting request payload"
// @Success 200 {object} public.PublicVotesGetAllResp "Successfully getted public Votes"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public votes"
// @Router /election/all [GET]
func (h *HandlerStruct) GetAllPublicVotes(c *gin.Context) {
	publicVote, err := h.PublicVote.GetAll(c.Request.Context(), &public.Filter{})
	if err != nil {
		c.JSON(400, gin.H{"Error when getting public Vote": err.Error()})
		return
	}
	c.JSON(200, publicVote)
}
