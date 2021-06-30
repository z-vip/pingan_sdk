package pingan_pkg

/**
请求参数
*/
type ArgsCheckMsgCodeWithCorp struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	SubAcctNo         string `json:"SubAcctNo" description:"子账户账号" required:"Y"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	MemberAcctNo      string `json:"MemberAcctNo" description:"会员账号" required:"Y"`
	MessageCheckCode  string `json:"MessageCheckCode" description:"短信验证码" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type CheckMsgCodeWithCorp struct {
	BaseResp
	FrontSeqNo string `json:"FrontSeqNo" description:"前置流水号"`
}
