package main

import (
    "github.com/gin-gonic/gin"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func init()  {
    // 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()
    // 记录到文件
    f, err := os.Create("runtime/access_"+time.Now().Format("2006-01-02")+".log")
    if err !=nil{
        log.Fatalln(err)
    }
    gin.DefaultWriter = io.MultiWriter(f)
}
func main() {


    r := gin.Default()
    gin.SetMode(gin.DebugMode)
    // r.Use(middleware.AccessLog())

    // 加载模板
    r.LoadHTMLGlob("app/template/**/*")
    r.Delims("{{", "}}")
    r.GET("/posts/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "post/index.html", gin.H{
            "title": "Posts",
        })
    })

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK,"%s","pong")
    })

    s := &http.Server{
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    if err := s.ListenAndServe(); err != nil {
        log.Fatalln("s.ListenAndServe Err:", err)
    }
}
