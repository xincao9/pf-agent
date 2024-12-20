package config

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/exec"
	"pf-agent/internal/constant"
	"strings"
)

var (
	C *viper.Viper
)

func init() {
	d := flag.Bool("d", false, "run app as a daemon with -d=true")
	c := flag.String("conf", "app.yaml", "configure file")
	if flag.Parsed() == false {
		flag.Parse()
	}
	if *d {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
	C = viper.New()
	for _, t := range []string{"yaml", "yml"} {
		if strings.HasSuffix(*c, t) {
			i := strings.LastIndex(*c, t)
			*c = string([]byte(*c)[:i-1])
		}
	}
	C.SetConfigName(*c)
	C.SetConfigType("yaml")
	C.AddConfigPath("./conf")
	C.SetDefault(constant.LoggerDir, "/tmp/pf-agent/log")
	C.SetDefault(constant.LoggerLevel, "debug")
	C.SetDefault(constant.ServerMode, "debug")
	C.SetDefault(constant.ServerPort, 8080)
	C.SetDefault(constant.ServerCorsAccessControlAllowOrigin, "http://localhost:8081")
	C.SetDefault(constant.ManagerServerPort, 8090)
	C.SetDefault(constant.DataSource, "root:asdf@tcp(localhost:3306)/golite?charset=utf8&parseTime=true")
	C.SetDefault(constant.AssetsRootDir, "./assets")
	C.SetDefault(constant.AssetsJsDir, "./assets/js")
	C.SetDefault(constant.AssetsCssDir, "./assets/css")
	C.SetDefault(constant.AssetsImgDir, "./assets/img")
	err := C.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error conf : %v\n", err)
	}
}

func Route(engine *gin.Engine) {
	engine.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, C.AllSettings())
	})
}
