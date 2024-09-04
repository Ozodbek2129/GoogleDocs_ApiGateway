package handler

import (
	"fmt"
	"net/http"

	pb "api_gateway/genproto/docs"
	"api_gateway/models"

	"github.com/gin-gonic/gin"
)

// @Summary      Create Document
// @Description  This endpoint creates a new document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for adding document"
// @Success      200    {object}  pb.CreateDocumentRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/createDocument [post]
func (h Handler) CreateDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.CreateDoc

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.CreateDocumentReq{AuthorId: id, Title: doc.Title}

	res, err := h.DocsService.CreateDocument(c, &req)
	if err != nil {
		h.Log.Error("CreateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Search Document
// @Description  This endpoint search a document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.SearchDocument true "Request body for searching document"
// @Success      200    {object}  pb.CreateDocumentRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/SearchDocument [get]
func (h Handler) SearchDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.SearchDocument

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.SearchDocumentReq{AuthorId: id, Title: doc.Title, DocsId: doc.DocsId}

	res, err := h.DocsService.SearchDocument(c, &req)
	if err != nil {
		h.Log.Error("SearchDocument function have problem.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Get all Document
// @Description  This endpoint gets all documents.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for getting all documents"
// @Success      200    {object}  pb.GetAllDocumentsRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/GetAllDocuments [get]
func (h Handler) GetAllDocuments(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.GetAllDocuments

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.GetAllDocumentsReq{AuthorId: id, Limit: doc.Limit, Page: doc.Page, DocsId: doc.DocsId}

	res, err := h.DocsService.GetAllDocuments(c, &req)
	if err != nil {
		h.Log.Error("GetAllDocuments function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Update Document
// @Description  This endpoint updates a document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for adding document"
// @Success      200    {object}  pb.UpdateDocumentRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/UpdateDocument [put]
func (h Handler) UpdateDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.UpdateDocument

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.UpdateDocumentReq{AuthorId: id, Title: doc.Title, Content: doc.Content, DocsId: doc.DocsId}

	res, err := h.DocsService.UpdateDocument(c, &req)
	if err != nil {
		h.Log.Error("UpdateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Delete Document
// @Description  This endpoint deletes a document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for deleting document"
// @Success      200    {object}  pb.DeleteDocumentRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/DeleteDocument [delete]
func (h Handler) DeleteDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.DeleteDocument

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.DeleteDocumentReq{AuthorId: id, Title: doc.Title}

	res, err := h.DocsService.DeleteDocument(c, &req)
	if err != nil {
		h.Log.Error("DeleteDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}


// @Summary      Share Document
// @Description  This endpoint shares a document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param        body models.CreateDoc true "Request body for sharing document"
// @Success      200    {object}  pb.ShareDocumentRes
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /api/docs/ShareDocument [post]
func (h Handler) ShareDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	var doc models.ShareDocument

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.ShareDocumentReq{Title: doc.Title, RecipientEmail: doc.RecipientEmail, Permissions: doc.Permissions, UserId: doc.UserId, Id: doc.Id}

	res, err := h.DocsService.ShareDocument(c, &req)
	if err != nil {
		h.Log.Error("CreateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}