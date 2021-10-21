# wxcom-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/mingzaily/wxcom-sdk)](https://goreportcard.com/report/github.com/mingzaily/wxcom-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mingzaily/wxcom-sdk)

企业微信 SDK

## Feature

- 通讯录管理（暂未计划）
- 应用管理
  - [ ] 获取应用
  - [ ] 设置应用
  - 自定义菜单
    - [ ] 创建菜单
    - [ ] 获取菜单
    - [ ] 删除菜单
- 消息管理
  - [x] 发送应用信息：支持文本、图片、语音、文件、文本卡片、markdown消息
  - 发送消息到群聊会话
    - [ ] 创建群聊会话
    - [ ] 修改群聊会话
    - [ ] 获取群聊会话
    - [ ] 应用推送信息

## 使用方式

```go
package main

import (
  "fmt"
  wxcom "github.com/mingzaily/go-wxcom"
)

func main() {
  client := wxcom.New("corpid", "corpsecret", 0)

  resp, err := client.Message().Text("测试").Send()
  if err != nil {
    panic(err)
  }
  fmt.Println(resp)
}

```
