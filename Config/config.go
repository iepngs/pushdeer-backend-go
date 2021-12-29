package Config

import (
    "fmt"
    "io/ioutil"
    "os"
    "sync"

    "gopkg.in/yaml.v2"
)

// config.yaml
type Server struct {
    Host string `yaml:"host"`
    Port int    `yaml:"port"`
}
type Mysql struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
    Charset  string `yaml:"charset"`
}
type Yaml struct {
    ServerConfig Server `yaml:"server"`
    MysqlConfig  Mysql  `yaml:"mysql"`
}

//login
type LoginProfile struct {
    Uid   string
    Email string
}
type Token struct {
    Token string `json:"token"`
}

//type LoginRespTpl struct{}

// user
type User struct {
    Id              int    `db:"id"`
    Name            string `db:"name"`
    Email           string `db:"email"`
    EmailVerifiedAt int64  `db:"email_verified_at"`
    Password        int64  `db:"password"`
}
type PasswordReset struct {
    Email     string `db:"email"`
    Token     string `db:"token"`
    CreatedAt int64  `db:"created_at"`
}
type FailedJobs struct {
    Id         int    `db:"id"`
    Uuid       string `db:"uuid"`
    Connection string `db:"connection"`
    Queue      string `db:"queue"`
    Payload    string `db:"payload"`
    Exception  string `db:"exception"`
    FailedAt   string `db:"failed_at"`
}
type PersonalAccessTokens struct {
    Id         int    `db:"id"`
    Tokenable  int    `db:"tokenable"`
    Name       string `db:"name"`
    Token      string `db:"token"`
    Abilities  string `db:"abilities"`
    LastUsedAt string `db:"last_used_at"`
}

var Configuration = new(Yaml)

func loadConfigs() {
    filename := "config.yaml"
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            initYamlFileWhenNotExist(filename)
            os.Exit(0)
        }
        panic(err)
    }
    err = yaml.Unmarshal(yamlFile, Configuration)
    if err != nil {
        panic(err)
    }
}

func initYamlFileWhenNotExist(filename string) {
    tpl := Yaml{
        ServerConfig: Server{
            Host: "http://127.0.0.1:8800",
            Port: 8800,
        },
        MysqlConfig: Mysql{
            Host:     "127.0.0.1",
            Port:     3306,
            Username: "root",
            Password: "root",
            Database: "pushdeer",
            Charset:  "UTF8MB4",
        },
    }
    out, err := yaml.Marshal(&tpl)
    if err != nil {
        panic(err)
    }
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer func() {
        _ = file.Close()
    }()
    if _, err = file.Write(out); err != nil {
        panic(err)
    }
    fmt.Printf("配置文件【%s】初始化完成，请编辑该文件配置好运行参数后再次运行！\n", filename)
}

func init() {
    var lc sync.Once
    lc.Do(loadConfigs)
}
