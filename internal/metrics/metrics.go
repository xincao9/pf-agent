package metrics

import (
	"github.com/gin-gonic/gin"
)

func Use(engine *gin.Engine) {
	p := ginprometheus.NewPrometheus("pf-agent")
	p.Use(engine)
}
