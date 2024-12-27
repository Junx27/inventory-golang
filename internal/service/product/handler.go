package product

import (
	"log"
	"net/http"
	"os"
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

func (h *Handler) GetProductByIDHandler(c *gin.Context) {
	id := c.Param("id")
	product, notFound, err := GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch product"})
		return
	}

	if notFound != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *Handler) GetAllProductsHandler(c *gin.Context) {
	products, err := GetAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (h *Handler) StoreProductHandler(c *gin.Context) {
	var req Product

	req.Name = c.PostForm("name")
	req.Description = c.PostForm("description")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid price format"})
		return
	}
	req.Price = price
	req.Category = c.PostForm("category")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no image file provided"})
		return
	}

	uploadPath := "./pkg/uploads/"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create upload directory"})
		return
	}

	filePath := uploadPath + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to upload image"})
		return
	}
	req.ImagePath = filePath

	err = StoreProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
		"data":    req,
	})
}

func (h *Handler) UpdateProductHandler(c *gin.Context) {
	var req Product

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	productID := c.Param("id")
	existingProduct, _, err := GetProductByID(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	file, _ := c.FormFile("image")
	if file != nil {
		uploadPath := "./pkg/uploads/" + file.Filename
		if err := os.MkdirAll("./pkg/uploads/", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create upload directory"})
			return
		}

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to upload image"})
			return
		}
		if existingProduct.ImagePath != "" {
			if err := os.Remove(existingProduct.ImagePath); err != nil {
				log.Println("Failed to delete old image:", err)
			}
		}

		req.ImagePath = uploadPath
	} else {

		req.ImagePath = existingProduct.ImagePath
	}

	req.ID = existingProduct.ID

	err = UpdateProduct(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

func (h *Handler) DeleteProductHandler(c *gin.Context) {

	productID := c.Param("id")

	existingProduct, notFound, err := GetProductByID(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch product"})
		return
	}

	if notFound != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound.Message})
		return
	}

	err = DeleteProduct(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete product"})
		return
	}

	if existingProduct.ImagePath != "" {
		if err := os.Remove(existingProduct.ImagePath); err != nil {
			log.Println("Failed to delete image file:", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
