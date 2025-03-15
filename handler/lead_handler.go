package handler

import (
	"go-crud-grpc/pb"
	"go-crud-grpc/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LeadHandler struct {
	repo *repository.LeadRepository
}

func NewLeadHandler(repo *repository.LeadRepository) *LeadHandler {
	return &LeadHandler{repo: repo}
}

// ✅ Get all leads
func (h *LeadHandler) GetLeads(c *gin.Context) {
	leads := h.repo.GetAllLeads()
	c.JSON(http.StatusOK, leads)
}

// ✅ Get a single lead by ID
func (h *LeadHandler) GetLead(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	lead, err := h.repo.GetLead(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	c.JSON(http.StatusOK, lead)
}

// ✅ Create a new lead
func (h *LeadHandler) CreateLead(c *gin.Context) {
	var lead pb.Lead
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newLead, err := h.repo.CreateLead(&lead)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLead)
}

// ✅ Update an existing lead
func (h *LeadHandler) UpdateLead(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	var lead pb.Lead
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	lead.Id = int32(id) // Set ID yang benar
	updatedLead, err := h.repo.UpdateLead(&lead)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLead)
}

// ✅ Delete a lead
func (h *LeadHandler) DeleteLead(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lead ID"})
		return
	}

	if err := h.repo.DeleteLead(int32(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lead deleted successfully"})
}
