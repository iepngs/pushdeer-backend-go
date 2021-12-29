package Models

import (
    "errors"
    "time"
)

type PushDeerDevice struct {
    Id        int64      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
    Uid       string     `gorm:"NOT NULL" json:"uid"`
    DeviceId  string     `gorm:"NOT NULL" json:"device_id" form:"device_id" binding:"required"`
    Type      string     `gorm:"NOT NULL" json:"type"`
    IsClip    int        `gorm:"NOT NULL" json:"is_clip" form:"is_clip" binding:"required"`
    Name      string     `gorm:"NOT NULL" json:"name" form:"name" binding:"required"`
    CreatedAt *time.Time `json:"omitempty"`
    UpdatedAt *time.Time `json:"omitempty"`
}

func (pdd *PushDeerDevice) RegisterDevice() error {
    pdd.Type = "all"
    return db.Create(pdd).Error
}

//获取设备列表
func GetMyDeviceList(uid string) (pdds []PushDeerDevice, err error) {
    err = db.Where(&PushDeerDevice{Uid: uid}).Find(&pdds).Error
    return
}

//移除已绑定设备
func (pdd PushDeerDevice) RemoveDevice() error {
    return db.Delete(&pdd).Error
}

//设备重命名
func (pdd PushDeerDevice) RenameDevice() error {
    return db.Save(&pdd).Error
}

//根据设备id查询设备
func (bdf *BaseDeviceFormTpl) GetDeviceById() (device PushDeerDevice, err error) {
    err = db.Where(&PushDeerDevice{Id: bdf.Id}).First(&device).Error
    if err == nil && device.Id == 0 {
        err = errors.New("设备不存在或已注销")
    }
    return
}

func GetDeviceByPushKey(uid string) (device []PushDeerDevice, err error) {
    err = db.Where(&PushDeerDevice{Uid: uid}).Find(&device).Error
    return
}

type BaseDeviceFormTpl struct {
    Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" form:"id" binding:"required"`
}

type RenameDeviceFormTpl struct {
    BaseDeviceFormTpl
    Name string `gorm:"NOT NULL" form:"name" binding:"required"`
}
