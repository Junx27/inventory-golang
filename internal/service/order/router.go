package order

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
	r.group.GET("/orders", r.handler.GetAllOrdersHandler)
	r.group.GET("/orders/:id", r.handler.GetOrderByIDHandler)
	r.group.POST("orders", r.handler.CreateOrderHandler)
	r.group.PUT("/orders/:id", r.handler.UpdateOrderHandler)
	r.group.DELETE("/orders/:id", r.handler.DeleteOrderHandler)
}
