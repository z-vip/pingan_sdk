package pingan_pkg

/**
请求参数
*/
type ArgsMembershipTrancheFreeze struct {
	FunctionFlag      string `json:"FunctionFlag" description:"功能标志" required:"Y"`
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	SubAcctNo         string `json:"SubAcctNo" description:"见证子账户的账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	TranAmt           string `json:"TranAmt" description:"交易金额" required:"Y"`
	TranCommission    string `json:"TranCommission" description:"交易手续费" required:"Y"`
	Ccy               string `json:"Ccy" description:"币种" required:"Y"`
	OrderNo           string `json:"OrderNo" description:"订单号" required:"Y"`
	OrderContent      string `json:"OrderContent" description:"订单内容" required:"N"`
	Remark            string `json:"Remark" description:"备注" required:"N"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type MembershipTrancheFreeze struct {
	BaseResp
	FrontSeqNo string `json:"FrontSeqNo" description:"前置流水号"`
}
