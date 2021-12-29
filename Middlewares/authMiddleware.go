package Middlewares

import (
    "pushdeer/Helpers"
    "pushdeer/Models"

    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        whiteList := [...]string{"/", "/error", "/user/login", "/logout"}
        for _, route := range whiteList {
            if route == c.Request.URL.Path {
                c.Next()
                return
            }
        }
        //whiteList = []string{"/static/*"}
        //for _, routeExp := range whiteList {
        //    exp := regexp.MustCompile(routeExp)
        //    if exp.MatchString(c.Request.URL.Path) {
        //        c.Next()
        //        return
        //    }
        //}
        token := c.DefaultPostForm("token", "")
        // 判断是否登录
        if token == "" {
            Helpers.ReplyTo(c, 403, "auth token invalid")
            c.Abort()
            return
        }

        // 初始化session对象
        session := sessions.Default(c)
        sessionCtx := session.Get(token)

        user,ok := sessionCtx.(Models.PushDeerUser)
        if !ok {
            Helpers.ReplyTo(c, 403)
            c.Abort()
            return
        }

        c.Set("token", user)
        c.Next()
    }
}
