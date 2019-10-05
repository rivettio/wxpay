package wxpay

import "encoding/xml"

type (
	WechatUnifiedOrderReq struct {
		XMLName        xml.Name `xml:"xml"`                                                //xml标签
		AppId          string   `xml:"appid" json:"appid"`                                 //微信分配的小程序ID，必须
		MchId          string   `xml:"mch_id" json:"mch_id"`                               //微信支付分配的商户号，必须
		DeviceInfo     string   `xml:"device_info" json:"device_info"`                     //微信支付填"WEB"，必须
		NonceStr       string   `xml:"nonce_str" json:"nonce_str"`                         //随机字符串，必须
		Sign           string   `xml:"sign" json:"sign"`                                   //签名，必须
		SignType       string   `xml:"sign_type" json:"sign_type"`                         //"HMAC-SHA256"或者"MD5"，非必须，默认MD5
		Body           string   `xml:"body" json:"body"`                                   //商品简单描述，必须
		Detail         string   `xml:"detail,omitempty" json:"detail,omitempty"`           //商品详细列表，使用json格式
		Attach         string   `xml:"attach" json:"attach"`                               //附加数据，如"贵阳分店"，非必须
		OutTradeNo     string   `xml:"out_trade_no" json:"out_trade_no"`                   //CRS订单号，必须
		FeeType        string   `xml:"fee_type,omitempty" json:"fee_type,omitempty"`       //默认人民币：CNY，非必须
		TotalFee       string   `xml:"total_fee" json:"total_fee"`                         //订单金额，单位分，必须
		SpbillCreateIp string   `xml:"spbill_create_ip" json:"spbill_create_ip"`           //支付提交客户端IP，如“123.123.123.123”，必须
		TimeStart      string   `xml:"time_start,omitempty" json:"time_start,omitempty"`   //订单生成时间，格式为yyyyMMddHHmmss，如20170324094700，非必须
		TimeExpire     string   `xml:"time_expire,omitempty" json:"time_expire,omitempty"` //订单结束时间，格式同上，非必须
		GoodsTag       string   `xml:"goods_tag,omitempty" json:"goods_tag,omitempty"`     //商品标记，代金券或立减优惠功能的参数，非必须
		NotifyUrl      string   `xml:"notify_url" json:"notify_url"`                       //接收微信支付异步通知回调地址，不能携带参数，必须
		TradeType      string   `xml:"trade_type" json:"trade_type"`                       //交易类型，小程序写"JSAPI"，必须
		LimitPay       string   `xml:"limit_pay,omitempty" json:"limit_pay,omitempty"`     //限制某种支付方式，非必须
		OpenId         string   `xml:"openid" json:"openid"`                               //微信用户唯一标识，必须
	}

	UnifiedOrderResp struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
		AppId      string `xml:"appid"`
		MchId      string `xml:"mch_id"`
		DeviceInfo string `xml:"device_info"`
		NonceStr   string `xml:"nonce_str"`
		Sign       string `xml:"sign"`
		ResultCode string `xml:"result_code"`
		ErrCode    string `xml:"err_code"`
		ErrCodeDes string `xml:"err_code_des"`
		TradeType  string `xml:"trade_type"`
		PrepayId   string `xml:"prepay_id"`
	}

	WechatNotifyInfo struct {
		ReturnCode         string `xml:"return_code,CDATA"  json:"return_code"`
		ReturnMsg          string `xml:"return_msg,CDATA"  json:"return_msg"`
		Appid              string `xml:"appid,CDATA" json:"appid"`
		MchId              string `xml:"mch_id,CDATA" json:"mch_id"`
		DeviceInfo         string `xml:"device_info,CDATA" json:"device_info"`
		NonceStr           string `xml:"nonce_str,CDATA" json:"nonce_str"`
		Sign               string `xml:"sign,CDATA" json:"sign"`
		SignType           string `xml:"sign_type,CDATA" json:"sign_type"`
		ResultCode         string `xml:"result_code,CDATA" json:"result_code"`
		ErrCode            string `xml:"err_code,CDATA" json:"err_code"`
		ErrCodeDes         string `xml:"err_code_des,CDATA" json:"err_code_des"`
		Openid             string `xml:"openid,CDATA" json:"openid"`
		IsSubscribe        string `xml:"is_subscribe,CDATA" json:"is_subscribe"`
		TradeType          string `xml:"trade_type,CDATA" json:"trade_type"`
		BankType           string `xml:"bank_type,CDATA" json:"bank_type"`
		TotalFee           uint   `xml:"total_fee,CDATA" json:"total_fee"`
		SettlementTotalFee uint   `xml:"settlement_total_fee" json:"settlement_total_fee"`
		FeeType            string `xml:"fee_type,CDATA" json:"fee_type"`
		CashFee            uint   `xml:"cash_fee,CDATA" json:"cash_fee"`
		CashFeeType        string `xml:"cash_fee_type,CDATA" json:"cash_fee_type"`
		CouponFee          uint   `xml:"coupon_fee,CDATA" json:"coupon_fee"`
		CouponCount        uint   `xml:"coupon_count,CDATA" json:"coupon_count"`
		CouponType0        uint   `xml:"coupon_type_0,CDATA" json:"coupon_type"`
		CouponId0          string `xml:"coupon_id_0,CDATA" json:"coupon_id"`
		CouponFee0         uint   `xml:"coupon_fee_0,CDATA" json:"coupon_fee"`
		TransactionId      string `xml:"transaction_id,CDATA" json:"transaction_id"`
		OutTradeNo         string `xml:"out_trade_no,CDATA" json:"out_trade_no"`
		Attach             string `xml:"attach,CDATA" json:"attach"`
		TimeEnd            string `xml:"time_end,CDATA" json:"time_end"`
	}

	Sextuple struct {
		AppId     string `json:"appId"`
		NonceStr  string `json:"nonceStr,omitempty"`
		Timestamp string `json:"timeStamp,omitempty"`
		PrepayId  string `json:"package,omitempty"`
		SignType  string `json:"signType,omitempty"`
		Sign      string `json:"paySign,omitempty"`
	}
)
