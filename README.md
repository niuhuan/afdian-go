AFDIAN-GO 爱发电
===============

爱发电SDK

golang 1.18 以上

## 使用方法

```go
import (
    "github.com/niuhuan/afdian-go"
)

// 创建客户端
client := &afdian.Client{
    UserId: "uid",
    Token: "token",
}
// 查看配置是否正常(成功时err为nil)
err := client.Ping()
// 查询订单
orderRsp, err := client.QueryOrder(1)
// 查询赞助者
sponsorRsp, err := client.QuerySponsor(1)

// 解析爱发电调用服务器
orderRsp, err := client.ParseOrder(body)

// 给爱发电的调用返回json
client.CallResponseString()

```