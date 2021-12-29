package Services

import (
    "errors"

    "github.com/gin-gonic/gin"
)

func GetUserAuthToken(c *gin.Context) string {
    return c.DefaultPostForm("token", "")
}

func CheckUserAuth(c *gin.Context) (ok bool, err error) {
    var token string
    if token, ok = c.GetPostForm("token"); !ok {
        token = ""
    }
    if token == "" {
        err = errors.New("auth token is not valid")
    }
    return
}

//func DiscernUserByToken(c *gin.Context) {
//    token, _ := c.Get("token")
//    user := Models.User{}
//    user.GetUserByToken(token.(string))
//}
