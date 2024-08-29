package account

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pf-agent/internal/constant"
	"pf-agent/internal/util"
	"pf-agent/model/account"
	accountService "pf-agent/service/account"
	"time"
)

func Route(engine *gin.Engine) {
	save := func(c *gin.Context) {
		a := &account.Account{}
		if err := c.ShouldBindJSON(a); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		err := accountService.A.Login(a) // 登录校验
		if err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		a.Token = uuid.New().String()
		a.Expire = time.Now().Add(time.Hour * time.Duration(constant.SessionExpireHour))
		err = accountService.A.Save(a) // 更新登录信息
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.SetCookie(constant.Token, a.Token, (int)((time.Hour*time.Duration(constant.SessionExpireHour))/time.Second), "/", "*", false, false)
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
	}
	// 登录
	engine.PUT("/session", save)
	engine.POST("/session", save)
}

func AuthenticationRoute(engine *gin.RouterGroup) {
	// 注销
	engine.DELETE("/session/:id", func(c *gin.Context) {
		sa, ok := c.Get(constant.SessionAccount)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		a := sa.(*account.Account)
		a, err := accountService.A.GetAccountByMobile(a.Mobile)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		a.Expire = time.Now()
		err = accountService.A.Save(a)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSON(c, http.StatusOK, constant.Success)
	})
	engine.GET("/account", func(c *gin.Context) {
		sa, ok := c.Get(constant.SessionAccount)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		a := sa.(*account.Account)
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
	})
}
