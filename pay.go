package wxpay

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/rivettio/logs"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	NonceStringLength     = 32
	XcxTRADE              = "JSAPI"
	WechatUnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	Env                   = "dev"
)

type wechatPay struct {
	WechatPayNotifyUrl string
	OrderDesc          string
	SpbilCreateIp      string
	WechatAppId        string
	WechatMchId        string
	WechatPayKey       string
}

func (w *wechatPay) Pay(totalFee uint32, orderSn, attach string, openId string) (*Sextuple, error) {

	nonce := GenRandStr(NonceStringLength)
	nonce = strings.ToUpper(nonce)
	unifiedOrderReq := WechatUnifiedOrderReq{

		NonceStr:       nonce,
		Sign:           "",
		Body:           w.OrderDesc,
		OutTradeNo:     orderSn,
		TotalFee:       strconv.Itoa(int(totalFee)),
		SpbillCreateIp: w.SpbilCreateIp,
		NotifyUrl:      w.WechatPayNotifyUrl,
		TradeType:      XcxTRADE,
		AppId:          w.WechatAppId,
		MchId:          w.WechatMchId,
		OpenId:         openId,
		Attach:         attach,
	}

	m, err := struct2Map(unifiedOrderReq)
	if err != nil {
		return nil, err
	}

	sign, err := GenWechatPaySign(m, w.WechatPayKey)
	if err != nil {
		return nil, err
	}
	unifiedOrderReq.Sign = strings.ToUpper(sign)

	unifiedOrderResp, err := WechatUnifiedOrder(unifiedOrderReq)
	if err != nil {
		logs.Errorf("pay unified order error %#v", err)
		return nil, err
	}

	var sextuple Sextuple
	sextuple.NonceStr = unifiedOrderResp.NonceStr
	sextuple.AppId = unifiedOrderResp.AppId
	sextuple.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	sextuple.PrepayId = "prepay_id=" + unifiedOrderResp.PrepayId
	sextuple.SignType = "MD5"
	logs.Debugf("pay sextuple: %#v", sextuple)

	m, err = struct2Map(sextuple)
	if err != nil {
		logs.Errorf("pay struct to map error %#v", err)
		return nil, err
	}

	sign, err = GenWechatPaySign(m, w.WechatPayKey)
	if err != nil {
		logs.Errorf("pay gen response sign error: %#v", err)
		return nil, err
	}

	sextuple.Sign = strings.ToUpper(sign)

	return &sextuple, nil
}

func PayCallBackHandle(data []byte, payKey string) (*WechatNotifyInfo, string, error) {
	var notify WechatNotifyInfo
	err := xml.Unmarshal(data, &notify)
	if err != nil {
		return nil, "", err
	}
	m, err := struct2Map(notify)
	if err != nil {
		return nil, "", err
	}
	sign, err := GenWechatPaySign(m, payKey)
	if err != nil {
		return nil, "", err
	}
	sign = strings.ToUpper(sign)
	return &notify, sign, err
}

func struct2Map(r interface{}) (s map[string]string, err error) {
	var temp map[string]interface{}
	var result = make(map[string]string)

	bin, err := json.Marshal(r)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(bin, &temp); err != nil {
		return nil, err
	}
	for k, v := range temp {
		switch v2 := v.(type) {
		case string:
			result[k] = v2
		case uint, int8, uint8, int, int16, uint16, int32, uint32, int64, uint64:
			result[k] = fmt.Sprintf("%d", v2)
		case float32, float64:
			result[k] = fmt.Sprintf("%.0f", v2)
		}
	}

	return result, nil
}

func GenWechatPaySign(m map[string]string, payKey string) (string, error) {
	delete(m, "sign")
	var signData []string
	for k, v := range m {
		if v != "" && v != "0" {
			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
		}
	}

	sort.Strings(signData)
	signStr := strings.Join(signData, "&")
	signStr = signStr + "&key=" + payKey
	c := md5.New()
	_, err := c.Write([]byte(signStr))
	if err != nil {
		return "", err
	}

	signByte := c.Sum(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", signByte), nil
}

func WechatUnifiedOrder(r WechatUnifiedOrderReq) (UnifiedOrderResp, error) {

	var payResp UnifiedOrderResp

	data, err := xml.Marshal(r)
	if err != nil {
		return payResp, err
	}
	fmt.Println("xml data:", string(data))
	logs.Error(" xml data: ", string(data))

	client := http.Client{}
	req, err := http.NewRequest("POST", WechatUnifiedOrderUrl, bytes.NewBuffer(data))
	if err != nil {
		return payResp, err
	}
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")

	resp, err := client.Do(req)

	if err != nil {
		return payResp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return payResp, err
	}
	defer resp.Body.Close()

	err = xml.Unmarshal(body, &payResp)
	if err != nil {
		return payResp, err
	}

	if payResp.ReturnCode != "SUCCESS" {
		return payResp, errors.New(payResp.ReturnMsg)
	}

	return payResp, nil
}

func New(wechatPayNotifyUrl, orderDesc, spbilCreateIp, wechatAppId, wechatMchId, wechatPayKey string) *wechatPay {

	return &wechatPay{
		WechatPayNotifyUrl: wechatPayNotifyUrl,
		OrderDesc:          orderDesc,
		SpbilCreateIp:      spbilCreateIp,
		WechatAppId:        wechatAppId,
		WechatMchId:        wechatMchId,
		WechatPayKey:       wechatPayKey,
	}
}
