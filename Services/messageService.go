package Services

import (
    "fmt"
    "pushdeer/Helpers"
    "pushdeer/Models"
    "sync"
)

type Notification struct {
    Tokens      []string
    Platform    int
    Title       string
    Message     string
    Development bool
    Production  bool
    Topic       string
    Port        int
    Sound       struct {
        Volume float32
    }
}
type Notifications struct {
    Notifications []Notification `json:"notifications"`
}

func PushMessage(m Models.MessageTpl, wg *sync.WaitGroup) bool {
    notify := Notification{
        Tokens:   []string{m.DeviceId},
        Platform: 1,
        Message:  m.Text,
        Topic:    "com.pushdeer.app.ios",
    }
    if len(m.Desp) > 0 {
        notify.Title = m.Text
        notify.Message = m.Desp
    }
    if m.Development {
        notify.Development = true
    } else {
        notify.Production = true
    }
    port := 8888
    if m.IsClip {
        port = 8889
        notify.Topic = "com.pushdeer.app.ios.Clip"
    }
    notify.Sound.Volume = 2.0
    payload := Notifications{
        Notifications: []Notification{notify},
    }
    url := fmt.Sprintf("http://127.0.0.1:%d/api/push", port)
    fmt.Println(url, payload)
    Helpers.SendCurlRequest(url, payload)
    wg.Done()
    return true
}
