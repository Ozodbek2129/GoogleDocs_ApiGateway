package handler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	pb "api_gateway/genproto/doccs"
	"api_gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// @Summary      Create Document
// @Description  This endpoint creates a new document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        title  path  string  false "Title of the document"
// @Success      200   {object}  doccs.CreateDocumentRes
// @Failure      400   {object}  models.Error
// @Failure      500   {object}  models.Error
// @Router       /api/docs/createDocument/{title} [post]
func (h Handler) CreateDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	title := c.Param("title")

	req := pb.CreateDocumentReq{AuthorId: id, Title: title}

	res, err := h.DocsService.CreateDocument(c, &req)
	if err != nil {
		h.Log.Error("CreateDocument function have problems.", "error", err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(title)
	fmt.Println(userId)

	c.JSON(200, res)
}

// @Summary      Search Document
// @Description  This endpoint searches for a document by title and document ID.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        title  query  string  false "Title of the document"
// @Param        docsId query  string  false "Document ID"
// @Success      200    {object}  doccs.SearchDocumentRes
// @Failure      400    {object}  models.Error
// @Failure      500    {object}  models.Error
// @Router       /api/docs/searchDocument [get]
func (h Handler) SearchDocument(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	title := c.Query("title")
	docsId := c.Query("docsId")

	req := pb.SearchDocumentReq{AuthorId: id, Title: title, DocsId: docsId}

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
// @Security     ApiKeyAuth
// @Param        get_all body models.CreateDoc true "Request body for getting all documents"
// @Success      200    {object}  doccs.GetAllDocumentsRes
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

const saveDebounceTime = 5 * time.Second

var clients = make(map[*websocket.Conn]bool)
var wsMutex sync.Mutex
var lastSavedContent = map[string]string{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketMessage struct {
	Action   string `json:"action"`
	Content  string `json:"content"`
	DocsId   string `json:"docs_id"`
	UserId   string `json:"user_id"`
}


func broadcastChanges(message WebSocketMessage) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}


func (h *Handler) WebSocketEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.Log.Error("Failed to upgrade to WebSocket", "error", err)
		return
	}
	defer conn.Close()

	wsMutex.Lock()
	clients[conn] = true
	wsMutex.Unlock()

	for {
		var message WebSocketMessage
		err := conn.ReadJSON(&message)
		if err != nil {
			wsMutex.Lock()
			delete(clients, conn)
			wsMutex.Unlock()
			break
		}
	}
}

func saveDocumentWithDebounce(docId, content string, h *Handler, c *gin.Context, req *pb.UpdateDocumentReq) {

	if lastSavedContent[docId] == content {
		return
	}

	time.Sleep(saveDebounceTime)

	if lastSavedContent[docId] != content {
		res, err := h.DocsService.UpdateDocument(c, req)
		if err != nil {
			h.Log.Error("Error saving document", "error", err.Error())
			return
		}
		lastSavedContent[docId] = content

		fmt.Println("Document saved:", res)
	}
}

// @Summary      Update Document
// @Description  Updates document and broadcasts changes via WebSocket.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body models.UpdateDocument true "Request body for updating document"
// @Success      200    {object}  doccs.UpdateDocumentRes
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

	var doc models.UpdateDocument

	if err := c.ShouldBindJSON(&doc); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, models.Error{Message: err.Error()})
		return
	}

	req := pb.UpdateDocumentReq{AuthorId: id, Title: doc.Title, Content: doc.Content, DocsId: doc.DocsId}

	message := WebSocketMessage{
		Action:  "update",
		Content: doc.Content,
		DocsId:  doc.DocsId,
		UserId:  id,
	}
	broadcastChanges(message)

	go saveDocumentWithDebounce(doc.DocsId, doc.Content, &h, c, &req)

	c.JSON(200, gin.H{"message": "Document update in progress"})
}

// @Summary      Delete Document
// @Description  This endpoint deletes a document.
// @Tags         docs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        delete body models.CreateDoc true "Request body for deleting document"
// @Success      200    {object}  doccs.DeleteDocumentRes
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
// @Security     ApiKeyAuth
// @Param        share body models.CreateDoc true "Request body for sharing document"
// @Success      200    {object}  doccs.ShareDocumentRes
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
