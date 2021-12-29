package Mysql

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "pushdeer/Config"
)

var DB *gorm.DB

func init() {
    var err error
    mc := Config.Configuration.MysqlConfig
    // root:root123@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
        mc.Username, mc.Password, mc.Host, mc.Port, mc.Database, mc.Charset)
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if DB.Error != nil {
        fmt.Printf("database error %v", DB.Error)
    }
}
