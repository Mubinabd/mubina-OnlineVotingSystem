package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	pb "project/genproto"
)

// CreateParty godoc
// @Summary Create a new Party
// @Description Endpoint for creating a new Party
// @Name create_Party
// @Date create_Party
// @Tags Party
// @Accept json
// @Produce json
// @Param Party body pb.PartyCreate true "Party creation request payload"
// @Success 200 {object} pb.Void "Successfully created Party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create Party"
// @Router /Party [POST]
func (h *HandlerStruct) CreateParty(c *gin.Context) {
	var body pb.PartyCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Party.CreateParty(context.Background(), &body)
	if err != nil {
		h.Logger.ERROR.Printf("could not create Party: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("Party created")
	c.JSON(http.StatusOK, gin.H{"status": "Party created"})
}

// GetById godoc
// @Summary Get By Id Party
// @Description Endpoint for getting Party
// @Tags Party
// @Accept json
// @Produce json
// @Param  id query pb.GetByIdReq true "ID"
// @Success 200 {object} pb.CandidateRes "Successfully getted Party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Party"
// @Router /Party/:id [GET]
func (h *HandlerStruct) GetByIdParty(c *gin.Context) {
	id := c.Param("id")

	Party, err := h.Party.GetParty(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.ERROR.Printf("could not get Party: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Party": Party})
	h.Logger.INFO.Printf("Party retrieved")
}

// GetAllParty godoc
// @Summary Get All Partys
// @Description Endpoint for getting all Partys
// @Tags Party
// @Accept json
// @Produce json
// @Param candidate query pb.GetAllPartysRequest true "Partys geting request payload"
// @Success 200 {object} pb.GetAllPartysResponse "Successfully getted Partys"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Partys"
// @Router /Party [GET]
func (h *HandlerStruct) GetAllParty(c *gin.Context) {
	var pbs pb.GetAllPartysRequest

	if err := c.ShouldBindJSON(&pbs); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Partys, err := h.Party.GetAllParty(context.Background(), &pbs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.ERROR.Printf("could not get Partys: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Partys": Partys})
}

// UpdateParty godoc
// @Summary Update Party
// @Description Endpoint for updating Party
// @Name update_Party
// @Date update_Party
// @Tags Party
// @Accept json
// @Param id path string true "User ID"
// @Param Party body pb.PartyCreate true "Party updaing request payload"
// @Success 200 {object} pb.Void "Successfully updated Party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update Party"
// @Router /Party/:id [PUT]
func (h *HandlerStruct) UpdateParty(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		h.Logger.ERROR.Printf("missing party ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing party ID"})
		return
	}

	var body pb.PartyCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Party.UpdateParty(context.Background(), &pb.PartyUpdate{
		UpdateParty: &body,
		Id:          &pb.GetByIdReq{Id: id},
	})
	if err != nil {
		h.Logger.ERROR.Printf("could not update party with ID %s: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("Party with ID %s updated", id)
	c.JSON(http.StatusOK, gin.H{"status": "Party updated"})
}

// DeleteParty godoc
// @Summary Delete Party
// @Description Endpoint for deleting Party
// @Name delete_Party
// @Date delete_Party
// @Tags Party
// @Accept json
// @Produce json
// @Param id path string true "Party ID"
// @Success 200 {object} pb.Void "Successfully deleted Party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete Party"
// @Router /Party/:id [DELETE]
func (h *HandlerStruct) DeleteParty(c *gin.Context) {
	id := c.Param("id")

	_, err := h.Party.DeleteParty(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		h.Logger.ERROR.Printf("could not delete Party: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("Party deleted")
	c.JSON(http.StatusOK, gin.H{"status": "Party deleted"})
}
