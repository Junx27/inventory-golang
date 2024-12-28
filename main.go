package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Junx27/inventory-golang/internal/config"
	"github.com/Junx27/inventory-golang/internal/service/inventory"
	"github.com/Junx27/inventory-golang/internal/service/product"
	"github.com/Junx27/inventory-golang/pkg/database"
	"github.com/caarlos0/env/v11"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.New(context.Background(), cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.Migrate(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	productHandler := product.NewHandler(cfg)
	inventoryHandler := inventory.NewHandler(cfg)

	var logger *zap.Logger
	var mode string

	switch cfg.Env {
	case "prod":
		mode = gin.ReleaseMode
		l, _ := zap.NewProduction()
		logger = l
	default:
		mode = gin.DebugMode
		l, _ := zap.NewDevelopment()
		logger = l
	}

	gin.SetMode(mode)

	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
	}))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(cors.Default())

	productRouter := product.NewRouter(productHandler, r.RouterGroup)
	productRouter.Register()

	inventoryRouter := inventory.NewRouter(inventoryHandler, r.RouterGroup)
	inventoryRouter.Register()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "server is run")
	})
	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
