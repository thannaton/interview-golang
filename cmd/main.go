package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thannaton/interview-golang/internal/core/services"
	"github.com/thannaton/interview-golang/internal/handler"
	logUtils "github.com/thannaton/interview-golang/internal/pkg/logs"
	"github.com/thannaton/interview-golang/internal/pkg/mdw"
)

func main() {
	r := gin.Default()

	// initialize env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%v: %v", "Error loading .env file", err)
	}

	// initailize log
	enableInfo := os.Getenv("ENABLE_INFO_LOG") == "true"
	enableWarning := os.Getenv("ENABLE_WARNING_LOG") == "true"
	enableDebug := os.Getenv("ENABLE_DEBUG_LOG") == "true"
	enableError := os.Getenv("ENABLE_ERROR_LOG") == "true"
	logUtils.InitLogs(enableInfo, enableWarning, enableDebug, enableError)

	// set up router verison and middleware
	apiVersion := os.Getenv("API_VERSION")
	apiContext := os.Getenv("API_CONTEXT")
	v1Router := r.Group(apiContext + "/" + apiVersion)

	// initialize middleware
	v1Router.Use(
		mdw.Logger(),
		gin.Recovery(),
	)

	// set up service and handler
	svr := services.NewService()
	hdr := handler.NewHandler(svr)
	handler.NewRouter(v1Router, hdr)

	port := os.Getenv("PORT")
	if err := r.Run(port); err != nil {
		panic(err)
	}
}
