package pkg

/**
请求参数
*/
type ArgsMemberTransaction struct {
	FunctionFlag      string `json:"FunctionFlag" description:"功能标志" required:"Y"`
	OutSubAcctNo      string `json:"OutSubAcctNo" description:"转出方的见证子账户的账号" required:"Y"`
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	OutMemberCode     string `json:"OutMemberCode" description:"转出方的交易网会员代码" required:"Y"`
	OutSubAcctName    string `json:"OutSubAcctName" description:"转出方的见证子账户的户名" required:"N"`
	InSubAcctNo       string `json:"InSubAcctNo" description:"转入方的见证子账户的账号" required:"Y"`
	InMemberCode      string `json:"InMemberCode" description:"转入方的交易网会员代码" required:"Y"`
	InSubAcctName     string `json:"InSubAcctName" description:"转入方的见证子账户的户名" required:"N"`
	TranAmt           string `json:"TranAmt" description:"交易金额" required:"Y"`
	TranFee           string `json:"TranFee" description:"交易费用" required:"Y"`
	TranType          string `json:"TranType" description:"交易类型" required:"Y"`
	Ccy               string `json:"Ccy" description:"币种" required:"Y"`
	OrderNo           string `json:"OrderNo" description:"订单号" required:"Y"`
	OrderContent      string `json:"OrderContent" description:"订单内容" required:"N"`
	Remark            string `json:"Remark" description:"备注" required:"N"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
	WebSign           string `json:"WebSign" description:"网银签名" required:"N"`
}

/*
返回数据
*/
type MemberTransaction struct {
	BaseResp
	FrontSeqNo string `json:"FrontSeqNo" description:"前置流水号"`
}
