package pkg

/**
请求参数
*/
type ArgsMembershipWithdrawCash struct {
	TranWebName       string `json:"TranWebName" description:"交易网名称" required:"Y"`
	SubAcctNo         string `json:"SubAcctNo" description:"见证子账户的账号" required:"Y"`
	MemberGlobalType  string `json:"MemberGlobalType" description:"会员证件类型" required:"Y"`
	MemberGlobalId    string `json:"MemberGlobalId" description:"会员证件号码" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	MemberName        string `json:"MemberName" description:"会员名称" required:"Y"`
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	TakeCashAcctNo    string `json:"TakeCashAcctNo" description:"提现账号" required:"Y"`
	OutAmtAcctName    string `json:"OutAmtAcctName" description:"出金账户名称" required:"Y"`
	Ccy               string `json:"Ccy" description:"币种" required:"Y"`
	CashAmt           string `json:"CashAmt" description:"可提现金额" required:"Y"`
	Remark            string `json:"Remark" description:"备注" required:"N"`
	ReservedMsg       string `json:"ReservedMsg" description:"手续费" required:"N"`
	WebSign           string `json:"WebSign" description:"网银签名" required:"N"`
}

/*
返回数据
*/
type MembershipWithdrawCash struct {
	BaseResp
	FrontSeqNo  string `json:"FrontSeqNo" description:"前置流水号"`
	TransferFee string `json:"TransferFee" description:"转账手续费"`
}
