package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"quebrada_api/docs"
	"quebrada_api/internal/app/middlewares"
	"quebrada_api/internal/app/router"
	"quebrada_api/internal/config"
	"quebrada_api/pkg/database"
	"quebrada_api/pkg/httpserver"
	logger "quebrada_api/pkg/logger"
	"syscall"
)

const configsDir = "configs"

func main() {
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Error(err)

		return
	}

	db := database.InitDB(cfg.Database.ConnectionString)

	r := gin.Default()
	handler := r.Handler()

	docs.SwaggerInfo.BasePath = "/api/v1"

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.GET("/alive", func(c *gin.Context) {
		c.String(http.StatusOK, "Its Alive and Kicking!")
	})

	m := middlewares.NewMonitoringMiddleware("quebrada_api", "/metrics")
	m.Use(r)

	routerManager := router.Router{
		AuthController: InitAuthController(db, cfg.SMPT),
	}

	api := r.Group("/api")
	{
		routerManager.Init(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	httpServer := httpserver.New(handler, httpserver.Port("9090"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatal("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	// Shutdown
	//goland:noinspection ALL
	httpServer.Shutdown()

}
