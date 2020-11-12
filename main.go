package main

import (
	"github.com/gin-gonic/gin"
	"zx/rout"
)

func main() {
	r := gin.Default()
	//r.StaticFS("/static",http.Dir("./html/static"))
	r.Static("static", "./static") //加载静态资源
	r.LoadHTMLGlob("./html/*")
	rout.Rout(r)
	r.Run(":9000")
}
