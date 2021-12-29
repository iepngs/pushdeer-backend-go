package Config

type ResponseContentTpl struct {
    Code    int         `json:"code"`
    Content interface{} `json:"content"`
    Error   string      `json:"error"`
}
