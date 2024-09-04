package handler

import (
	"fmt"
	"net/http"

	pb "api_gateway/genproto/doccs"
	"api_gateway/models"

	"github.com/gin-gonic/gin"
)

// @Summary      Get All Document Version
// @Description  This endpoint gets all documents version.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.GetAllVersions true "Request body for getting all versions of document"
// @Success      200    {object}  doccs.GetAllVersionsRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/version/GetAllVersions [get]
func (h Handler) GetAllVersions(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.GetAllVersions

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.GetAllVersionsReq{AuthorId: id, Title: doc.Title}

	res, err := h.DocsService.GetAllVersions(c, &req)
	if err != nil {
		h.Log.Error("CreateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Restore Document's Version
// @Description  This endpoint restores a document's version.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for adding document"
// @Success      200    {object}  doccs.RestoreVersionRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/vesion/RestoreVersion [put]
func (h Handler) RestoreVersion(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.RestoreVersion

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.RestoreVersionReq{AuthorId: id, Title: doc.Title}

	res, err := h.DocsService.RestoreVersion(c, &req)
	if err != nil {
		h.Log.Error("CreateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}