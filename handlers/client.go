package handlers

import (
	"lms/models"
	"lms/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{
		service: service.NewClientService(service.NewDB()),
	}
}

// @BasePath /
// GetClients godoc
// @Summary Get all clients
// @Description Returns all clients
// @Produce json
// @Success 200
// @Router /clients [get]
func (h *ClientHandler) GetClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.service.GetClients())
}

func (h *ClientHandler) GetClient(c *gin.Context) {
	raw_id := c.Param("id")

	id, err := strconv.Atoi(raw_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	client := h.service.GetClient(uint64(id))

	if client == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var createClientRequst models.CreateClientRequest

	if err := c.BindJSON(&createClientRequst); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newClient := h.service.CreateClient(
		createClientRequst.Name,
		createClientRequst.WorkingHours,
		createClientRequst.Priority,
		createClientRequst.LeadCapacity,
	)

	c.IndentedJSON(http.StatusCreated, newClient)
}

func (h *ClientHandler) AssignLeadToClient(c *gin.Context) {
	var assignLeadRequest models.AssignLeadRequest

	if err := c.BindJSON(&assignLeadRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availableClient := h.service.FindLead(assignLeadRequest.MeetingStartTime, assignLeadRequest.MeetingEndTime)

	if availableClient == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no client available"})
		return
	}

	availableClient = h.service.IncrementLeadCount(availableClient.ID)
	c.IndentedJSON(http.StatusOK, availableClient)
}
