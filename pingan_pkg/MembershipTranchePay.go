package pingan_pkg

/**
请求参数
*/
type ArgsMembershipTranchePay struct {
	FundSummaryAcctNo    string         `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	OutSubAcctNo         string         `json:"OutSubAcctNo" description:"转出见证子账户的账号" required:"Y"`
	OutTranNetMemberCode string         `json:"OutTranNetMemberCode" description:"转出交易网会员代码" required:"Y"`
	Commission           string         `json:"Commission" description:"手续费" required:"Y"`
	Ccy                  string         `json:"Ccy" description:"币种" required:"Y"`
	OrderNo              string         `json:"OrderNo" description:"订单号" required:"Y"`
	InSubAcctNum         string         `json:"InSubAcctNum" description:"转入见证子账户数" required:"Y"`
	TranItemArray        []ArgsTranItem `json:"TranItemArray" description:"交易信息数组" required:"Y"`
	Remark               string         `json:"Remark" description:"备注" required:"N"`
	ReservedMsg          string         `json:"ReservedMsg" description:"保留域" required:"N"`
}

type ArgsTranItem struct {
	TranSeqNo           string `json:"TranSeqNo" description:"交易流水号" required:"Y"`
	InSubAcctNo         string `json:"InSubAcctNo" description:"转入见证子账户的账号" required:"Y"`
	InTranNetMemberCode string `json:"InTranNetMemberCode" description:"转入交易网会员代码" required:"Y"`
	TranAmt             string `json:"TranAmt" description:"交易金额" required:"Y"`
}

/*
返回数据
*/
type MembershipTranchePay struct {
	BaseResp
	FrontSeqNo string `json:"FrontSeqNo" description:"前置流水号"`
}
