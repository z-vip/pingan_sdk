package pkg

/**
请求参数
*/
type ArgsReconciliationDocumentQuery struct {
	FundSummaryAcctNo string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	FileType          string `json:"FileType" description:"文件类型" required:"Y"`
	FileDate          string `json:"FileDate" description:"文件日期" required:"Y"`
	ReservedMsg       string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type ReconciliationDocumentQuery struct {
	BaseResp
	ResultNum     string                            `json:"ResultNum" description:"本次交易返回查询结果记录数"`
	TranItemArray []ReconciliationDocumentQueryItem `json:"TranItemArray" description:"交易信息数组"`
}
type ReconciliationDocumentQueryItem struct {
	FileName       string `json:"FileName" description:"文件名称"`
	RandomPassword string `json:"RandomPassword" description:"随机密码"`
	FilePath       string `json:"FilePath" description:"文件路径"`
	DrawCode       string `json:"DrawCode" description:"提取码"`
}
