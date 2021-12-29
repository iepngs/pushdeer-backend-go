#### 项目说明
这只是 [pushdeer](https://github.com/easychen/pushdeer) [后端项目](https://github.com/easychen/pushdeer/tree/main/api) 一个粗糙的Go版本。

#### 运行环境
- go 1.13+
- mysql 5.7+

#### 运行方式
直接运行或者编译后运行，首次运行会生成`config.yaml`配置模板文件，更改它，之后再次运行，如果一切正常的话即可正常启动程序。

#### 功能模块
可在ApiPost下查看项目 [API文档](https://docs.apipost.cn/preview/05869c50a8c093c0/b5980324865db3c7)

- [ ] 设备相关
    - [x] 设备注册
    - [x] 设备列表
    - [x] 移除设备
    - [x] 设备重命名
- [ ] 用户相关
    - [ ] 用户注册
    - [x] 用户信息
- [ ] 消息相关
    - [x] 推送消息
    - [x] 删除消息
    - [x] 消息列表
- [ ] 推送KEY
    - [x] 删除KEY
    - [x] 重置KEY
    - [x] 重命名KEY
    - [x] 新增KEY
    - [x] KEY列表

#### 迭代功能
- [ ] 响应状态码统一
- [ ] 推送结果处理

#### 项目测试
- [X] P0事故惯犯