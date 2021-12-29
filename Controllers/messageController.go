package Controllers

import (
    "fmt"
    "pushdeer/Helpers"
    "pushdeer/Models"
    "pushdeer/Services"
    "strconv"
    "sync"

    "github.com/gin-gonic/gin"
)

//获取所有的推送消息
func GetMyMessages(c *gin.Context) {
    user, _ := Helpers.GetUserProfileFromSession(c)
    uid := fmt.Sprintf("%d", user.Id)
    var (
        keys []Models.PushDeerMessage
        err  error
    )
    num := 100
    limit, ok := c.GetPostForm("limit")
    if ok {
        num, _ = strconv.Atoi(limit)
    }
    if keys, err = Models.GetMyMessages(uid, num); err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    Helpers.ReplyTo(c, keys)
}

//推送消息
func PushMessage(c *gin.Context) {
    var (
        msg Models.PushDeerMessage
        err error
    )
    if err := c.ShouldBind(&msg); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    msg, err = msg.GetUidByPushKey(msg.ReadKey)
    if err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var ms []Models.MessageTpl
    ms, err = msg.PushMessage()
    if err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    // TODO. 推送结果处理
    var (
        wg  sync.WaitGroup
        ret map[string]bool
    )
    for _, message := range ms {
        wg.Add(1)
        go func(msg Models.MessageTpl) {
            ret[msg.DeviceId] = Services.PushMessage(msg, &wg)
        }(message)
    }
    wg.Wait()
    Helpers.ReplyTo(c, ret)
}

//删除消息
func RemovePushMessage(c *gin.Context) {
    k := Models.BasePushMessageFormTpl{}
    if err := c.Bind(&k); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var (
        key Models.PushDeerMessage
        err error
    )
    if key, err = k.GetPushMessageById(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    if err := key.RemoveMessage(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    Helpers.ReplyTo(c, "done")
}
