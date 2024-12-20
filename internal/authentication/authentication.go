package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pf-agent/internal/constant"
	"pf-agent/internal/util"
	accountService "pf-agent/service/account"
	"time"
)

func Authentication(c *gin.Context) {
	t, err := c.Cookie(constant.Token) // 请求必须携带token
	if err != nil {
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, "the request must carry a token")
		return
	}
	a, err := accountService.A.GetAccountByToken(t)
	if err != nil || a == nil || a.Expire.Before(time.Now()) { // 会话对象是否过期
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, "session expired or nonexistent session")
		return
	}
	c.Set(constant.SessionAccount, a) // 设置本地会话
}
