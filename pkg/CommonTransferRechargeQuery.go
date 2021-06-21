package pkg

/**
请求参数
*/
type ArgsCommonTransferRechargeQuery struct {
	FunctionFlag      string `json:"FunctionFlag" description:"功能标志" required:"Y"`
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	StartDate         string `json:"StartDate" description:"开始日期" required:"Y"`
	EndDate           string `json:"EndDate" description:"终止日期" required:"Y"`
	PageNum           string `json:"PageNum" description:"页码" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type CommonTransferRechargeQuery struct {
	BaseResp
	ResultNum     string                            `json:"ResultNum" description:"本次交易返回查询结果记录数"`
	StartRecordNo string                            `json:"StartRecordNo" description:"起始记录号"`
	EndFlag       string                            `json:"EndFlag" description:"结束标志"`
	TotalNum      string                            `json:"TotalNum" description:"符合业务查询条件的记录总数"`
	TranItemArray []CommonTransferRechargeQueryItem `json:"TranItemArray" description:"交易信息数组"`
}

type CommonTransferRechargeQueryItem struct {
	InAcctType        string `json:"InAcctType" description:"入账类型"`
	TranNetMemberCode string `json:"TranNetMemberCode" description:"交易网会员代码"`
	SubAcctNo         string `json:"SubAcctNo" description:"见证子帐户的帐号"`
	TranAmt           string `json:"TranAmt" description:"入金金额"`
	InAcctNo          string `json:"InAcctNo" description:"入金账号"`
	InAcctName        string `json:"InAcctName" description:"入金账户名称"`
	Ccy               string `json:"Ccy" description:"币种"`
	AccountingDate    string `json:"AccountingDate" description:"会计日期"`
	BankName          string `json:"BankName" description:"银行名称"`
	Remark            string `json:"Remark" description:"转账备注"`
	FrontSeqNo        string `json:"FrontSeqNo" description:"见证系统流水号"`
}
