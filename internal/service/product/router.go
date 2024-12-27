package product

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler Handler
	group   gin.RouterGroup
}

func NewRouter(handler Handler, group gin.RouterGroup) Router {
	return Router{
		handler: handler,
		group:   group,
	}
}

func (r *Router) Register() {
	r.group.GET("/products", r.handler.GetAllProductsHandler)
	r.group.GET("/products/:id", r.handler.GetProductByIDHandler)
	r.group.POST("/products", r.handler.StoreProductHandler)
	r.group.PUT("/products/:id", r.handler.UpdateProductHandler)
	r.group.DELETE("/products/:id", r.handler.DeleteProductHandler)
}
