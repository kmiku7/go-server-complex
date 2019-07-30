package main

import (
	"common"
	"flag"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"

	"config"
	"logger"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(origLogClient *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := uuid.New().String()
		startTime := time.Now()
		requestLogClient := origLogClient.WithField(common.ContextRequestID, reqId)
		c.Set(common.ContextLogClient, requestLogClient)

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		bytesSend := c.Writer.Size()

		requestLogClient.Infof(
			"path: %v, client: %v, status: %v, bytes: %v, cost: %v",
			path, clientIP, statusCode, bytesSend, latency)
	}
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "", "Configure file path.")
	flag.Parse()

	if err := config.InitGlobalConfig(configPath); err != nil {
		panic(err)
	}

	logClient, closer, err := logger.OpenLogger(config.GetGlobalConfig().Log)
	if err != nil {
		panic(err)
	}
	defer closer()

	r := gin.New()
	r.Use(LoggerMiddleware(logClient))
	r.GET("/ping/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"pong"})
	})
	r.Run(config.GetGlobalConfig().Core.ListenAddress)
}
