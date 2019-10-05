### wxpay

Lightweight payment tools for WeChat.

### Installation

`go get github.com/rivettio/wxpay`

### Quick Start

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rivettio/logs"
	"github.com/rivettio/wxpay"
	"io/ioutil"
)

func main() {
	var sextuple *wxpay.Sextuple
	wxp := wxpay.New("pay notify url", "order description", "113.97.33.19", "app id", "mchId", "key")
	sextuple, err := wxp.Pay(0, "orderSn", "WX pay type", "WX openID")
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Errorf("pay unified order error %#v", sextuple)
}

func PushWXPayNotice(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logs.Error(err)
	}
	notify, sign, err := wxpay.PayCallBackHandle(data, "WX pay type")
	logs.Info("pay call back handle info:", notify, sign)
}
```

