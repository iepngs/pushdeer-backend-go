package Models

import (
    "errors"
    "math/rand"
    "strconv"
    "time"
)

type PushDeerMessage struct {
    Id        int64      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
    Uid       string     `gorm:"NOT NULL" json:"uid"`
    Text      string     `gorm:"NOT NULL" json:"text" form:"text" binding:"required"`
    Type      string     `gorm:"NOT NULL;DEFAULT:markdown" json:"type" form:"type"`
    Desp      string     `json:"desp" form:"desp"`
    ReadKey   string     `gorm:"NOT NULL" json:"readkey" form:"pushkey" binding:"required"`
    Url       string     `json:"url"`
    CreatedAt *time.Time `json:"omitempty"`
    UpdatedAt *time.Time `json:"omitempty"`
}

type BasePushMessageFormTpl struct {
    Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" form:"id" binding:"required"`
}

type PushMessageResultTpl struct {
    Result []bool `json:"result"`
}

type MessageTpl struct {
    Uid         string
    Text        string
    Desp        string
    ReadKey     string
    DeviceId    string
    Development bool
    IsClip      bool
}

func GetMyMessages(uid string, limit int) (ms []PushDeerMessage, err error) {
    err = db.Where(&PushDeerKey{Uid: uid}).Find(&ms).Limit(limit).Error
    return
}

//根据推送key获取uid
func (bdm *PushDeerMessage) GetUidByPushKey(key string) (m PushDeerMessage, err error) {
    err = db.Where(&PushDeerMessage{ReadKey: key}).First(&key).Error
    if err == nil && m.Id == 0 {
        err = errors.New("key invalid")
    }
    return
}

//根据推送key获取uid
func (bdm *PushDeerMessage) PushMessage() (ms []MessageTpl, err error) {
    var devices []PushDeerDevice
    devices, err = GetDeviceByPushKey(bdm.Uid)
    if err != nil {
        return
    }
    if len(devices) == 0 {
        return
    }
    for _, device := range devices {
        message := MessageTpl{
            IsClip:   device.IsClip == 1,
            Uid:      device.Uid,
            Text:     bdm.Text,
            Desp:     bdm.Desp,
            DeviceId: device.DeviceId,
            ReadKey:  strconv.Itoa(rand.Int()),
        }
        ms = append(ms, message)
    }
    return
}

//func (bdm *PushDeerMessage) PushMessage() (m PushMessageResultTpl, err error) {
//    var devices []PushDeerDevice
//    devices, err = GetDeviceByPushKey(bdm.Uid)
//    if err != nil {
//        return
//    }
//    if len(devices) == 0 {
//        return
//    }
//    var wg sync.WaitGroup
//    for _, device := range devices {
//        message := Services.MessageTpl{
//            IsClip:   device.IsClip == 1,
//            Uid:      device.Uid,
//            Text:     bdm.Text,
//            Desp:     bdm.Desp,
//            DeviceId: device.DeviceId,
//            ReadKey:  strconv.Itoa(rand.Int()),
//        }
//        wg.Add(1)
//        go func(msg Services.MessageTpl) {
//            m.Result = append(m.Result, msg.Push(&wg))
//        }(message)
//    }
//    wg.Wait()
//    return
//}

//根据消息id查询消息
func (bdf *BasePushMessageFormTpl) GetPushMessageById() (key PushDeerMessage, err error) {
    err = db.Where(&PushDeerDevice{Id: bdf.Id}).First(&key).Error
    if err == nil && key.Id == 0 {
        err = errors.New("消息不存在或已删除")
    }
    return
}

//删除消息
func (bdm PushDeerMessage) RemoveMessage() error {
    return db.Delete(&bdm).Error
}
