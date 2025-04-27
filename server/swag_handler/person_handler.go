package swag_handler

import (
	"log/slog"
	"net/http"
	"person-enrichment-service/server/entity"
	"person-enrichment-service/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	service service.PersonService
	logger  *slog.Logger
}

func NewPersonHandler(service service.PersonService, logger *slog.Logger) *PersonHandler {
	return &PersonHandler{
		service: service,
		logger:  logger,
	}
}

// CreatePerson godoc
// @Summary Create a new person
// @Description Create a new person with the input payload
// @Tags persons
// @Accept  json
// @Produce  json
// @Param person body entity.CreatePersonRequest true "Create person"
// @Success 201 {object} entity.PersonResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons [post]
func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var req entity.CreatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := h.service.CreatePerson(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create person", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, person)
}

// GetPerson godoc
// @Summary Get a person by ID
// @Description Get a person by ID
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {object} entity.PersonResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons/{id} [get]
func (h *PersonHandler) GetPerson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Failed to parse ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	person, err := h.service.GetPersonByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Error("Failed to get person", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if person == nil {
		h.logger.Info("Person not found", "id", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

// GetPeople godoc
// @Summary Get people with filtering and pagination
// @Description Get people with filtering and pagination
// @Tags persons
// @Accept  json
// @Produce  json
// @Param name query string false "Name filter"
// @Param surname query string false "Surname filter"
// @Param age query int false "Age filter"
// @Param gender query string false "Gender filter"
// @Param nationality query string false "Nationality filter"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} map[string]interface{} "{"data": [], "total": 0}"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons [get]
func (h *PersonHandler) GetPeople(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	age, _ := strconv.Atoi(c.Query("age"))

	filter := entity.FilterOptions{
		Name:        c.Query("name"),
		Surname:     c.Query("surname"),
		Age:         age,
		Gender:      c.Query("gender"),
		Nationality: c.Query("nationality"),
		Page:        page,
		PageSize:    pageSize,
	}

	people, total, err := h.service.GetAllPersons(c.Request.Context(), filter)
	if err != nil {
		h.logger.Error("Failed to get people", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  people,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// UpdatePerson godoc
// @Summary Update a person
// @Description Update a person by ID
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Param person body entity.UpdatePersonRequest true "Update person"
// @Success 200 {object} entity.PersonResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons/{id} [put]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Failed to parse ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req entity.UpdatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := h.service.UpdatePerson(c.Request.Context(), uint(id), &req)
	if err != nil {
		h.logger.Error("Failed to update person", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

// DeletePerson godoc
// @Summary Delete a person
// @Description Delete a person by ID
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Failed to parse ID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeletePerson(c.Request.Context(), uint(id)); err != nil {
		h.logger.Error("Failed to delete person", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
