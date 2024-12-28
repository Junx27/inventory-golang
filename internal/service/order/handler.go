package order

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Junx27/inventory-golang/internal/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg config.Config
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		cfg: cfg,
	}
}

func (h *Handler) GetAllOrdersHandler(c *gin.Context) {
	orders, err := GetAllOrders(c.Request.Context())
	if err != nil {
		log.Printf("Failed to fetch orders: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *Handler) GetOrderByIDHandler(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid order ID"})
		return
	}

	order, err := GetOrderByID(c.Request.Context(), strconv.Itoa(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch order"})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h *Handler) CreateOrderHandler(c *gin.Context) {
	var req Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	err := CreateOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully", "data": req})
}

func (h *Handler) UpdateOrderHandler(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid order ID"})
		return
	}

	var req Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	err = UpdateOrder(c.Request.Context(), orderID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

func (h *Handler) DeleteOrderHandler(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid order ID"})
		return
	}

	err = DeleteOrder(c.Request.Context(), strconv.Itoa(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
