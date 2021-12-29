package Models

import (
    "errors"
    "time"
)

type PushDeerKey struct {
    Id        int64      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
    Name      string     `gorm:"NOT NULL" json:"name" form:"name" binding:"required"`
    Uid       string     `gorm:"NOT NULL" json:"uid"`
    Key       string     `gorm:"NOT NULL" json:"key"`
    CreatedAt *time.Time `json:"omitempty"`
    UpdatedAt *time.Time `json:"omitempty"`
}

//生成推送KEY
func (pdk *PushDeerKey) GeneratePushKey() error {
    return db.Create(pdk).Error
}

//获取所有的推送KEY
func GetMyPushKeys(uid string) (pdks []PushDeerKey, err error) {
    err = db.Where(&PushDeerKey{Uid: uid}).Find(&pdks).Error
    return
}

//移除已绑定设备
func (pdk PushDeerKey) RemovePushKey() error {
    return db.Delete(&pdk).Error
}

//设备重命名
func (pdk PushDeerKey) RenamePushKey() error {
    return db.Save(&pdk).Error
}

//根据设备id查询设备
func (bdf *BasePushKeyFormTpl) GetPushKeyById() (key PushDeerKey, err error) {
    err = db.Where(&PushDeerDevice{Id: bdf.Id}).First(&key).Error
    if err == nil && key.Id == 0 {
        err = errors.New("设备不存在或已注销")
    }
    return
}

type BasePushKeyFormTpl struct {
    Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" form:"id" binding:"required"`
}

type RenamePushKeyFormTpl struct {
    BasePushKeyFormTpl
    Name string `gorm:"NOT NULL" form:"name" binding:"required"`
}
