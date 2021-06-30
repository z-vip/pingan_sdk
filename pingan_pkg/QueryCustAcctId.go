package pingan_pkg

/**
请求参数
*/
type ArgsQueryCustAcctId struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type QueryCustAcctId struct {
	BaseResp
	SubAcctNo        string `json:"SubAcctNo" description:"见证子账户的账号"`
	SubAcctCashBal   string `json:"SubAcctCashBal" description:"见证子账户可提现余额"`
	SubAcctAvailBal  string `json:"SubAcctAvailBal" description:"见证子账户可用余额"`
	SubAcctFreezeAmt string `json:"SubAcctFreezeAmt" description:"见证子账户冻结金额"`
}
