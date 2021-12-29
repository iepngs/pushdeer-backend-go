package Controllers

import (
    "fmt"
    "math/rand"
    "pushdeer/Helpers"
    "pushdeer/Models"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

//获取所有的推送KEY
func GetMyPushKeys(c *gin.Context) {
    user, _ := Helpers.GetUserProfileFromSession(c)
    uid := fmt.Sprintf("%d", user.Id)
    var (
        keys []Models.PushDeerKey
        err  error
    )
    if keys, err = Models.GetMyPushKeys(uid); err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    Helpers.ReplyTo(c, keys)
}

//生成推送KEY
func GeneratePushKey(c *gin.Context) {
    rand.Seed(time.Now().UnixNano())
    user, _ := Helpers.GetUserProfileFromSession(c)
    uid := fmt.Sprintf("%d", user.Id)
    pdd := Models.PushDeerKey{
        Name: "key" + strconv.Itoa(rand.Int()),
        Uid:  uid,
        Key:  "PDU" + uid + "T" + strconv.Itoa(rand.Int()),
    }
    var err error
    if err = pdd.GeneratePushKey(); err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    GetMyPushKeys(c)
}

//移除设备
func RemovePushKey(c *gin.Context) {
    k := Models.BasePushKeyFormTpl{}
    if err := c.Bind(&k); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var (
        key Models.PushDeerKey
        err error
    )
    if key, err = k.GetPushKeyById(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    if err := key.RemovePushKey(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    Helpers.ReplyTo(c, "done")
}

//设备重命名
func RenamePushKey(c *gin.Context) {
    k := Models.RenamePushKeyFormTpl{}
    if err := c.Bind(&k); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var (
        key Models.PushDeerKey
        err error
    )
    if key, err = k.GetPushKeyById(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    key.Name = k.Name
    if err := key.RenamePushKey(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    Helpers.ReplyTo(c, "done")
}
