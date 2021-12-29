package Controllers

import (
    "fmt"
    "pushdeer/Config"
    "pushdeer/Helpers"
    "pushdeer/Models"

    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func UserLogin(c *gin.Context) {
    lp := Config.LoginProfile{
        Uid:   "theid999",
        Email: "easychen+new@gmail.com",
    }
    if lp.Uid == "" {
        Helpers.ReplyTo(c, Helpers.ErrorCode("ARGS"), "id_token解析错误")
        return
    }
    user, err := Models.FindOutLoginUser(lp)
    if err != nil {
        Helpers.ReplyTo(c, 500, err.Error())
        return
    }
    fmt.Println(user)
    token := Config.Token{Token: uuid.New().String()}

    session := sessions.Default(c)
    session.Set(token, user)
    _ = session.Save()

    //session.Delete("tizi365")
    //_ = session.Save()

    Helpers.ReplyTo(c, token)
}

func UserInfo(c *gin.Context) {
    Helpers.ReplyTo(c, 200)
}
