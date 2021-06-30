package pingan_pkg

/**
请求参数
*/
type ArgsSupAcctIdBalanceQuery struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type SupAcctIdBalanceQuery struct {
	BaseResp
	LastBalance  string `json:"LastBalance" description:"上日余额"`
	CurBalabce   string `json:"CurBalabce" description:"当前余额"`
	Balance      string `json:"Balance" description:"账户余额"`
	AddedBalance string `json:"AddedBalance" description:"增值余额"`
}
