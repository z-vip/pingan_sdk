package pkg

/**
请求参数
*/
type ArgsQueryCustAcctIdBalance struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type QueryCustAcctIdBalance struct {
	BaseResp
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码"`
	SubAcctAvailBal   string `json:"SubAcctAvailBal" description:"见证子账户可用余额"`
	SubAcctAssureAmt  string `json:"SubAcctAssureAmt" description:"见证子账户担保金额"`
}
