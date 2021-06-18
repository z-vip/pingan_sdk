package pkg

type BaseArgs struct {
	ApiVersionNo   string `json:"ApiVersionNo" description:"API版本号"`
	AppAccessToken string `json:"AppAccessToken" description:"app访问token"`
	ApplicationID  string `json:"ApplicationID" description:"应用KEY"`
	CnsmrSeqNo     string `json:"CnsmrSeqNo" description:"交易流水号"`
	RequestMode    string `json:"RequestMode" description:"请求模式"`
	SDKType        string `json:"SDKType" description:"接入方式"`
	SdkSeid        string `json:"SdkSeid" description:"SDK识别码"`
	SdkVersionNo   string `json:"SdkVersionNo" description:"SDK版本号"`
	TranStatus     string `json:"TranStatus" description:"RSA签名值"`
	TxnTime        string `json:"TxnTime" description:"发送时间"`
	ValidTerm      string `json:"ValidTerm" description:"有效期"`
	TxnCode        string `json:"TxnCode" description:"交易码"`
	TxnClientNo    string `json:"TxnClientNo" description:"交易客户号" required:"N"`
	MrchCode       string `json:"MrchCode" description:"商户号" required:"N"`
	CerSeqNo       string `json:"CerSeqNo" description:"证书序列号" required:"N"`
	Mac            string `json:"Mac" description:"Mac地址" required:"N"`
	RsaSign        string `json:"RsaSign" description:"加密签名值"`
}

type BaseResp struct {
	TxnReturnCode   string `json:"TxnReturnCode" description:"返回码"`
	TxnReturnMsg    string `json:"TxnReturnMsg" description:"返回信息"`
	CnsmrSeqNo      string `json:"CnsmrSeqNo" description:"交易流水号"`
	RsaSign         string `json:"RsaSign" description:"加密签名值"`
	TokenExpiryFlag string `json:"tokenExpiryFlag" description:"Token过期"`
	ReservedMsg     string `json:"ReservedMsg" description:"保留域"`
}
