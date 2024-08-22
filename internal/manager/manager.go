package manager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"pf-agent/internal/config"
	"pf-agent/internal/constant"
	"pf-agent/internal/logger"
	"pf-agent/internal/metrics"
)

func init() {
	gin.SetMode(config.C.GetString(constant.ServerMode))
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
	config.Route(engine) // 配置服务接口
	metrics.Use(engine)
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ManagerServerPort))
	logger.L.Infof("Management listening and serving HTTP on : %s", addr)
	go func() {
		if err := engine.Run(addr); err != nil {
			logger.L.Fatalf("Fatal error app: %v\n", err)
		}
	}()
}
