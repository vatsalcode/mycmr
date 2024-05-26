package controllers

import (
    "net/http"
    "mycrm/models"
    "mycrm/services"
    "github.com/gin-gonic/gin"
)

type InteractionController struct {
    interactionService *services.InteractionService
}

func NewInteractionController(interactionService *services.InteractionService) *InteractionController {
    return &InteractionController{interactionService: interactionService}
}

func (controller *InteractionController) CreateInteraction(c *gin.Context) {
    var interaction models.Interaction
    if err := c.ShouldBindJSON(&interaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := controller.interactionService.CreateInteraction(&interaction)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, interaction)
}

func (controller *InteractionController) GetInteractions(c *gin.Context) {
    customerID := c.Param("id")
    interactions, err := controller.interactionService.GetInteractionsByCustomer(customerID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, interactions)
}

func (controller *InteractionController) GetInteractionStats(c *gin.Context) {
    stats, err := controller.interactionService.CalculateStats()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, stats)
}
