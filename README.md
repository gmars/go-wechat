# go-wechat
开箱即用的微信公众平台sdk集成了微信开放平台第三方平台api及微信支付商户平台功能

# 安装使用
## 安装

```go
go get github.com/gmars/go-wechat
```

## 使用

## 微信公众号使用

### 事件处理器
```go
// 初始化客户端
// 如只调用api则第一个配置即可
// 如只处理消息第二个配置即可
cli, err := official.NewOfficialClient(
    official.WithMessageHandlerConfig("", "", ""),
    official.WithBaseAutoAccessTokenConfig("", "", nil))

if err != nil {
    panic(err)
}

r := gin.Default()
r.POST("/message", func(c *gin.Context) {
    res, err := cli.MessageHandler().Handler(c.Request)
    if err != nil {
        //TODO 记录日志
    }

    //普通消息
    if res.ReplayType == message.ReplyTypeStandard {
        switch res.TypeValue {
        case message.MsgText:
            msg, ok := res.Data.(message.TextMsg)
            if !ok {
                //TODO 记录日志
            }
            rep := message.GenerateText(msg.ToUserName, msg.FromUserName, "测试")
            c.Writer.Write([]byte(rep))
            break
        }
    }
})
```

### 接口调用

```go

menus, err := cli.Menu().GetMenu()
if err != nil {
    panic(err)
}
fmt.Println(menus)
```

## 开放平台代公众号实现

### 代处理消息
```go
//初始化客户端和公众号一样
cli, err := openplatform.NewOpenPlatformClient(
    openplatform.WithMessageHandlerConfig("", "", ""),
    openplatform.WithAutoAccessTokenConfig("", "", nil))

//获取消息处理器
handler := cli.MessageHandler()
```

### 代调用接口

```go
// 次回调允许开发者自行保存刷新后的授权凭证
func tokenCallBack(authorizerAppId string, tokenRes *authorizer_token.AuthorizerTokenRes) {
	fmt.Println(authorizerAppId)
	fmt.Println(tokenRes.AuthorizerRefreshToken)
}

func main() {
	cli, err := openplatform.NewOpenPlatformClient(
		openplatform.WithMessageHandlerConfig("", "", ""),
		openplatform.WithAutoAccessTokenConfig("", "", nil))
	if err != nil {
		panic(err)
	}

	//初始化授权方access tokener
	authorizerToken := cli.AuthorizerAccessToken("authorizer app id", "refresh token", tokenCallBack)

	//初始化公众号客户端
	officialClient, err := official.NewOfficialClient(
		official.WithAccessToken(authorizerToken))
	if err != nil {
		panic(err)
	}

	menus, err := officialClient.Menu().GetMenu()
	if err != nil {
		panic(err)
	}
	fmt.Println(menus)
}
```

## 关于缓存

如果使用时配置中cache传入nil则该包会自行使用文件缓存。如果要使用类似redis的缓存请实现core中的cache接口

## 关于access token

关于access token可以直接配置在客户端中，如果需要自行实现来管理access token则实现core中的accesstoken接口即可