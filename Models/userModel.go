package Models

import (
    "pushdeer/Config"
    "strings"
    "time"
)

//type User struct {
//    Id              int       `gorm:"primary_key"`
//    Name            string    `gorm:"type:varchar(256);not null;`
//    Email           string    `gorm:"unique_index:hash_idx;"`
//    EmailVerifiedAt time.Time `db:"email_verified_at"`
//    Password        int64     `gorm:"type:varchar(256);not null;`
//}
type PushDeerUser struct {
    Id        int64      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
    Name      string     `json:"name"`
    Email     string     `json:"email"`
    AppleId   string     `gorm:"UNIQUE_INDEX" json:"apple_id"`
    WechatId  string     `json:"wechat_id"`
    Level     int        `gorm:"NOT NULL" json:"level"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

// 如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
//db.SingularTable(true)

// 设置User的表名为`user`
//func (User) TableName() string {
//    return "user"
//}
//
//func (u *User) CreateUser() (user User, err error) {
//    u.EmailVerifiedAt = time.Now()
//    if err = db.Create(u).Error; err != nil {
//        return
//    }
//    return *u, nil
//}

func GetPushDeerUserByAppleId(uid string) (pdu PushDeerUser, err error) {
    db.Where(&PushDeerUser{AppleId: uid}).First(&pdu)
    return pdu, db.Error
}

func FindOutLoginUser(l Config.LoginProfile) (pdu PushDeerUser, err error) {
    if pdu, err = GetPushDeerUserByAppleId(l.Uid); err != nil {
        return
    }
    if pdu.AppleId == "" {
        pdu = PushDeerUser{
            AppleId: l.Uid,
            Email:   l.Email,
            Name:    strings.Split(l.Email, "@")[0],
            Level:   1,
        }
        err = db.Create(&pdu).Error
    }
    return
}

//func (u *User) GetUserByToken(token string) string {
//    result := db.Create(&u)
//    if result.Error != nil {
//        return result.Error.Error()
//    }
//    return ""
//}
