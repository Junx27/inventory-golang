package inventory

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

	r.group.GET("/inventory", r.handler.GetAllInventoriesHandler)
	r.group.GET("/inventory/:id", r.handler.GetInventoryByIDHandler)
	r.group.POST("/inventory", r.handler.StoreInventoryHandler)
	r.group.PUT("/inventory/:id", r.handler.UpdateInventoryHandler)
	r.group.DELETE("/inventory/:id", r.handler.DeleteInventoryHandler)

}
