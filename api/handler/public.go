package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	pb "project/genproto"
)

// CreatePublic godoc
// @Summary Create a new public
// @Description Endpoint for creating a new public
// @Name create_public
// @Date create_public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body pb.PublicCreate true "Public creation request payload"
// @Success 200 {object} pb.Void "Successfully created public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create public"
// @Router /public [POST]
func (h *HandlerStruct) CreatePublic(c *gin.Context) {
	var body pb.PublicCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Public.Create(context.Background(), &body)
	if err != nil {
		h.Logger.ERROR.Printf("could not create public: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("public created")
	c.JSON(http.StatusOK, gin.H{"status": "public created"})
}

// GetById godoc
// @Summary Get By Id Public
// @Description Endpoint for getting public
// @Tags Public
// @Accept json
// @Produce json
// @Param  id query pb.GetByIdReq true "ID"
// @Success 200 {object} pb.CandidateRes "Successfully getted public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public"
// @Router /public/:id [GET]
func (h *HandlerStruct) GetByIdPublic(c *gin.Context) {
	id := c.Param("id")

	public, err := h.Public.Get(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.ERROR.Printf("could not get public: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"public": public})
	h.Logger.INFO.Printf("public retrieved")
}

// GetAllPublic godoc
// @Summary Get All Publics
// @Description Endpoint for getting all publics
// @Tags Public
// @Accept json
// @Produce json
// @Param candidate query pb.GetAllPublicsRequest true "Publics geting request payload"
// @Success 200 {object} pb.GetAllPublicsResponse "Successfully getted publics"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get publics"
// @Router /public [GET]
func (h *HandlerStruct) GetAllPublic(c *gin.Context) {
	var pbs pb.GetAllPublicsRequest

	if err := c.ShouldBindJSON(&pbs); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publics, err := h.Public.GetAll(context.Background(), &pbs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.ERROR.Printf("could not get publics: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"publics": publics})
}

// UpdatePublic godoc
// @Summary Update public
// @Description Endpoint for updating public
// @Name update_public
// @Date update_public
// @Tags Public
// @Accept json
// @Param id path string true "User ID"
// @Param public body pb.PublicCreate true "Public updaing request payload"
// @Success 200 {object} pb.Void "Successfully updated public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update public"
// @Router /public/:id [PUT]
func (h *HandlerStruct) UpdatePublic(c *gin.Context) {
	id := c.Param("id")

	var body pb.PublicCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Printf("could not bind JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Public.Update(context.Background(), &pb.PublicUpdate{
		UpdatePublic: &body,
		Id:           &pb.GetByIdReq{Id: id},
	})
	if err != nil {
		h.Logger.ERROR.Printf("could not update public: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("public updated")
	c.JSON(http.StatusOK, gin.H{"status": "public updated"})
}

// DeletePublic godoc
// @Summary Delete public
// @Description Endpoint for deleting public
// @Name delete_public
// @Date delete_public
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Success 200 {object} pb.Void "Successfully deleted public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete public"
// @Router /public/:id [DELETE]
func (h *HandlerStruct) DeletePublic(c *gin.Context) {
	id := c.Param("id")

	_, err := h.Public.Delete(context.Background(), &pb.GetByIdReq{Id: id})
	if err != nil {
		h.Logger.ERROR.Printf("could not delete public: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.INFO.Printf("public deleted")
	c.JSON(http.StatusOK, gin.H{"status": "public deleted"})
}
