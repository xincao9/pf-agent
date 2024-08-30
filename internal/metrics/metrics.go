package metrics

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func Use(engine *gin.Engine) {
	p := ginprometheus.NewPrometheus("pf_agent")
	p.Use(engine)
}
