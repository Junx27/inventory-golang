package inventory

import (
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

func (h *Handler) GetInventoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	inventory, err := GetInventoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "inventory no exist"})
		return
	}

	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})
}

func (h *Handler) GetAllInventoriesHandler(c *gin.Context) {
	inventories, err := GetAllInventories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch inventories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inventories})
}

func (h *Handler) StoreInventoryHandler(c *gin.Context) {
	var req Inventory

	req.ProductID, _ = strconv.Atoi(c.PostForm("product_id"))
	req.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))
	req.Location = c.PostForm("location")

	err := CreateInventory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create inventory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "inventory created successfully",
		"data":    req,
	})
}

func (h *Handler) UpdateInventoryHandler(c *gin.Context) {
	var req Inventory

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	inventoryID := c.Param("id")
	existingInventory, err := GetInventoryByID(c.Request.Context(), inventoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		return
	}

	req.ID = existingInventory.ID

	err = UpdateInventory(c.Request.Context(), inventoryID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "inventory updated successfully"})
}

func (h *Handler) DeleteInventoryHandler(c *gin.Context) {
	inventoryID := c.Param("id")

	_, err := strconv.Atoi(inventoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid inventory ID"})
		return
	}
	inventory, err := GetInventoryByID(c.Request.Context(), inventoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch inventory"})
		return
	}

	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "inventory not found"})
		return
	}

	err = DeleteInventory(c.Request.Context(), inventoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "inventory deleted successfully"})
}
