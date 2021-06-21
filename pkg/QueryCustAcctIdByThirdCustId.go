package pkg

/**
请求参数
*/
type ArgsQueryCustAcctIdByThirdCustId struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type QueryCustAcctIdByThirdCustId struct {
	BaseResp
	SubAcctNo string `json:"SubAcctNo" description:"见证子账户的账号"`
}
