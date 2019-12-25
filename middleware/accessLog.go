package middleware

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/satori/go.uuid"
    "time"
)

func AccessLog() gin.HandlerFunc {
    return func(c *gin.Context) {

        var respBody string
        /*// form方式
        jPostForm, _ := json.Marshal(c.Request.PostForm)
        postForm := string(jPostForm)

        // json方式的请求
        jBody, _ := ioutil.ReadAll(c.Request.Body)
        reqBody := string(jBody)*/

        // 开始时间
        startTime := time.Now()

        // 处理请求
        c.Next()

        // 结束时间
        endTime := time.Now()

        // 请求id
        requestId := uuid.NewV4().String()

        // 请求方式
        method := c.Request.Method
        // 请求路由
        url := c.Request.URL.String()
        // 请求类型
        contentType := c.ContentType()
        // 请求地址
        remoteAddr := c.Request.RemoteAddr
        // 请求IP
        clientIP := c.ClientIP()
        // 执行时间
        latencyTime := endTime.Sub(startTime)

        // 状态码
        statusCode := c.Writer.Status()
        // 响应大小
        statusSize := c.Writer.Size()

        respJson, err := json.Marshal(c.Keys["resp"])
        if err == nil {
            respBody = string(respJson)
        } else {
            respBody = " "
        }
        fmt.Println(
            startTime.Format("2006-01-02 15:04:05"),"|",
            requestId,"|",
            method, "|",
            url, "|",
            contentType, "|",
            remoteAddr, "|",
            clientIP,"|",
            statusCode, "|",
            statusSize,  "|",
            latencyTime,"|",
            respBody,
        )

    }
}
