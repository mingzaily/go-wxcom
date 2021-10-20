# wxcom-sdk

企业微信 SDK

## Feature

- 通讯录管理
  - [ ] 成员管理
  - [ ] 部门管理
- 消息管理
  - [x] 发送应用信息
  - 发送消息到群聊会话
    - [ ] 应用推送信息

## 使用方式
```go
package main

import (
  wxcom "github.com/mingzaily/wxcom-sdk"
  message "github.com/mingzaily/wxcom-sdk/message"
)

client := wxcom.New("corpid", "corpsecret", agentid)

// 应用信息推送
ms := message.NewWithClient(client).SendAppMessage(
	message.NewAppTxtMessage(...)
)

```
