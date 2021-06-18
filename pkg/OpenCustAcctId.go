package pkg

const (
	MemberPropertyDefault  = "00"
	MemberPropertyBusiness = "SH"
)

/**
请求参数
*/
type ArgsOpenCustAcctId struct {
	FunctionFlag      string `json:"FunctionFlag" description:"功能标志"`
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号"`
	MemberProperty    string `json:"MemberProperty" description:"会员属性"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码"`
	UserNickname      string `json:"UserNickname" description:"用户昵称" required:"N"`
	Mobile            string `json:"Mobile" description:"手机号码" required:"N"`
	Email             string `json:"Email" description:"邮箱" required:"N"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type OpenCustAcctId struct {
	BaseResp
	SubAcctNo   string `json:"SubAcctNo" description:"见证子账户的账号"`
	ReservedMsg string `json:"ReservedMsg" description:"保留域"`
}
