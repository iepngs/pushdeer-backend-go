package Router

import (
    "pushdeer/Config"
    "pushdeer/Controllers"
    "pushdeer/Helpers"
    "pushdeer/Middlewares"

    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
)

func InitRouter() {
    router := gin.Default()
    // 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
    //router.Use(Middlewares.Cors())

    store := cookie.NewStore([]byte("PushDeer@2022"))
    router.Use(sessions.Sessions("PushDeer", store))

    router.Use(Middlewares.AuthMiddleware())

    router.GET("/login/fake", func(c *gin.Context) {
        c.Redirect(302, "/user/login")
    })

    user := router.Group("/user")
    {
        user.GET("/login", Controllers.UserLogin)
        user.POST("/info", Controllers.UserInfo)
    }

    device := router.Group("/device")
    {
        device.POST("/reg", Controllers.RegisterDevice)
        device.POST("/list", Controllers.GetMyDeviceList)
        device.POST("/remove", Controllers.RemoveDevice)
        device.POST("/rename", Controllers.RenameDevice)
    }

    message := router.Group("/message")
    {
        message.POST("/push", Controllers.PushMessage)
        message.POST("/list", Controllers.GetMyMessages)
        message.POST("/remove", Controllers.RemovePushMessage)
    }

    key := router.Group("/key")
    {
        key.POST("/gen", Controllers.GeneratePushKey)
        key.POST("/list", Controllers.GetMyPushKeys)
        key.POST("/remove", Controllers.RemovePushKey)
        key.POST("/regen", Controllers.RenamePushKey)
    }

    router.NoRoute(func(c *gin.Context) {
        Helpers.ReplyTo(c, 404)
    })

    err := router.Run(Config.Configuration.ServerConfig.Host)
    if err != nil {
        panic(err)
    }
}
