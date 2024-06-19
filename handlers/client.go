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
// @Router /clients [get] []models.Client
func (h *ClientHandler) GetClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.service.GetClients())
}

// @basePath /
// GetClient godoc
// @summary Get client by id
// @description Get client by id
// @param id path int true "id of client"
// @produce json
// @success 200 {object} models.Client
// @router /clients/{id} [get] models.Client
func (h *ClientHandler) GetClient(c *gin.Context) {
	raw_id := c.Param("id")

	id, err := strconv.Atoi(raw_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Error: "invalid client id"})
		return
	}

	client := h.service.GetClient(uint64(id))

	if client == nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Error: "client not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}

// @basePath /
// GetClient godoc
// @summary Create client
// @description Create client
// @produce json
// @param Client body models.CreateClientRequest true "Client data to be created"
// @failure 400 {object} models.Error "Invalid request provided"
// @success 200 {object} models.Client
// @router /clients [post] models.Client
func (h *ClientHandler) CreateClient(c *gin.Context) {
	var createClientRequst models.CreateClientRequest

	if err := c.BindJSON(&createClientRequst); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Error: err.Error()})
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

// @basePath /
// AssignLeadToClient godoc
// @summary Assigns available lead to client
// @description Assigns available lead to client
// @produce json
// @failure 400 {object} models.Error "Invalid request provided"
// @failure 404 {object} models.Error "Client not found"
// @success 200 {object} models.Client "Client object"
// @router /lead [get] models.Client
func (h *ClientHandler) AssignLeadToClient(c *gin.Context) {
	var assignLeadRequest models.AssignLeadRequest

	if err := c.BindJSON(&assignLeadRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Error: err.Error()})
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
