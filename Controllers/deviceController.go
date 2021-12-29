package Controllers

import (
    "fmt"
    "pushdeer/Helpers"
    "pushdeer/Models"

    "github.com/gin-gonic/gin"
)

//注册设备
func RegisterDevice(c *gin.Context) {
    pdd := Models.PushDeerDevice{}
    if err := c.ShouldBind(&pdd); err != nil {
        Helpers.ReplyTo(c, 500, err.Error())
        return
    }
    user, _ := Helpers.GetUserProfileFromSession(c)
    pdd.Uid =  fmt.Sprintf("%d", user.Id)
    if err := pdd.RegisterDevice(); err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    GetMyDeviceList(c)
}

//获取已注册的设备
func GetMyDeviceList(c *gin.Context) {

    user, _ := Helpers.GetUserProfileFromSession(c)
    uid := fmt.Sprintf("%d", user.Id)
    pdd, err := Models.GetMyDeviceList(uid)
    if err != nil {
        Helpers.ReplyTo(c, 500)
        return
    }
    Helpers.ReplyTo(c, pdd)
}

//移除设备
func RemoveDevice(c *gin.Context) {
    d := Models.BaseDeviceFormTpl{}
    if err := c.Bind(&d); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var (
        device Models.PushDeerDevice
        err    error
    )
    if device, err = d.GetDeviceById(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    if err := device.RemoveDevice(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    Helpers.ReplyTo(c, "done")
}

//设备重命名
func RenameDevice(c *gin.Context) {
    d := Models.RenameDeviceFormTpl{}
    if err := c.Bind(&d); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    var (
        device Models.PushDeerDevice
        err    error
    )
    if device, err = d.GetDeviceById(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    device.Name = d.Name
    if err := device.RenameDevice(); err != nil {
        Helpers.ReplyTo(c, err)
        return
    }
    Helpers.ReplyTo(c, "done")
}
