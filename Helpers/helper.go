package Helpers

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "pushdeer/Config"
    "pushdeer/Models"

    "github.com/gin-gonic/gin"
)

const ServerCrash = "系统异常，请稍后重试"

//TODO. 合并到ReplyTo
func ErrorCode(kind string) (code int) {
    switch kind {
    case "AUTH":
        code = 80403
    case "ARGS":
        code = 80501
    case "REMOTE":
        code = 80502
    default:
        code = 80999
    }
    return code
}

func ReplyTo(c *gin.Context, a ...interface{}) {
    rct := Config.ResponseContentTpl{}
    for inx, arg := range a {
        if inx > 2 {
            break
        }
        switch data := arg.(type) {
        case int:
            rct.Code = data
        case string:
            rct.Error = data
        case error:
            rct.Error = data.Error()
        default:
            rct.Content = data
        }
    }
    if rct.Content == nil {
        rct.Content = struct{}{}
    }
    if rct.Error == "" {
        switch rct.Code {
        case 400:
            rct.Error = "400 - bad request"
        case 403:
            rct.Error = "403 - Forbidden"
        case 404:
            rct.Error = "404 - 请求地址不存在"
        case 500:
            rct.Error = ServerCrash
        }
    } else {
        if rct.Code == 0 {
            rct.Code = 500
        }
    }
    c.JSON(http.StatusOK, rct)
}

func SendCurlRequest(requestUrl string, payload interface{}) bool {
    jsonData, _ := json.Marshal(payload)
    resp, err := http.Post(requestUrl, "application/json", bytes.NewReader(jsonData))
    if err != nil {
        log.Printf("get request failed, err:[%s]", err.Error())
        return false
    }
    defer func() {
        _ = resp.Body.Close()
    }()
    bodyContent, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
    fmt.Printf("resp body data:[%s]\n", string(bodyContent))
    return true
}

func GetUserProfileFromSession(c *gin.Context)(user Models.PushDeerUser, err error){
    sessionCtx, _ := c.Get("token")
    var ok bool
    user,ok = sessionCtx.(Models.PushDeerUser)
    if !ok {
        err = errors.New("token is not valid")
        return
    }
    return
}
