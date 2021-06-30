package pingan_pkg

/**
请求参数
*/
type ArgsCheckAmountWithCorp struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	SubAcctNo         string `json:"SubAcctNo" description:"子账户账号" required:"Y"`
	TakeCashAcctNo    string `json:"TakeCashAcctNo" description:"提现账号" required:"Y"`
	AuthAmt           string `json:"AuthAmt" description:"鉴权金额" required:"Y"`
	OrderNo           string `json:"OrderNo" description:"指令号" required:"Y"`
	Ccy               string `json:"Ccy" description:"币种" required:"Y"`
}

/*
返回数据
*/
type CheckAmountWithCorp struct {
	BaseResp
	FrontSeqNo string `json:"FrontSeqNo" description:"前置流水号"`
}
