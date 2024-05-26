package controllers

import (
    "net/http"
    "mycrm/models"
    "mycrm/services"
    "github.com/gin-gonic/gin"
)

type CustomerController struct {
    customerService *services.CustomerService
}

func NewCustomerController(customerService *services.CustomerService) *CustomerController {
    return &CustomerController{
        customerService: customerService,
    }
}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := cc.customerService.CreateCustomer(&customer)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, customer)
}

func (cc *CustomerController) GetCustomer(c *gin.Context) {
    id := c.Param("id")
    customer, err := cc.customerService.GetCustomerByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
    id := c.Param("id")
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := cc.customerService.UpdateCustomer(id, &customer)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
    id := c.Param("id")
    err := cc.customerService.DeleteCustomer(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
