package pingan_pkg

/**
请求参数
*/
type ArgsBindSmallAmountWithCheckCorp struct {
	FundSummaryAcctNo      string `json:"FundSummaryAcctNo" description:"资金汇总账号" required:"Y"`
	SubAcctNo              string `json:"SubAcctNo" description:"子账户账号" required:"Y"`
	TranNetMemberCode      string `json:"TranNetMemberCode" description:"交易网会员代码" required:"Y"`
	MemberName             string `json:"MemberName" description:"会员名称" required:"Y"`
	MemberGlobalType       string `json:"MemberGlobalType" description:"会员证件类型" required:"Y"`
	MemberGlobalId         string `json:"MemberGlobalId" description:"会员证件号码" required:"Y"`
	MemberAcctNo           string `json:"MemberAcctNo" description:"会员账号" required:"Y"`
	BankType               string `json:"BankType" description:"银行类型" required:"Y"`
	AcctOpenBranchName     string `json:"AcctOpenBranchName" description:"开户行名称" required:"Y"`
	CnapsBranchId          string `json:"CnapsBranchId" description:"大小额行号" required:"N"`
	EiconBankBranchId      string `json:"EiconBankBranchId" description:"超级网银行号" required:"N"`
	Mobile                 string `json:"Mobile" description:"手机号码" required:"Y"`
	IndivBusinessFlag      string `json:"IndivBusinessFlag" description:"是否个体工商户" required:"Y"`
	CompanyName            string `json:"CompanyName" description:"公司名称" required:"N"`
	CompanyGlobalType      string `json:"CompanyGlobalType" description:"公司证件类型" required:"N"`
	CompanyGlobalId        string `json:"CompanyGlobalId" description:"公司证件号码" required:"N"`
	ShopId                 string `json:"ShopId" description:"店铺id" required:"N"`
	ShopName               string `json:"ShopName" description:"店铺名称" required:"N"`
	AgencyClientFlag       string `json:"AgencyClientFlag" description:"是否存在经办人" required:"Y"`
	AgencyClientName       string `json:"AgencyClientName" description:"经办人姓名" required:"N"`
	AgencyClientGlobalType string `json:"AgencyClientGlobalType" description:"经办人证件类型" required:"N"`
	AgencyClientGlobalId   string `json:"AgencyClientGlobalId" description:"经办人证件号" required:"N"`
	AgencyClientMobile     string `json:"AgencyClientMobile" description:"经办人手机号" required:"N"`
	RepFlag                string `json:"RepFlag" description:"会员名称是否是法人" required:"Y"`
	ReprName               string `json:"ReprName" description:"法人名称" required:"N"`
	ReprGlobalType         string `json:"ReprGlobalType" description:"法人证件类型" required:"N"`
	ReprGlobalId           string `json:"ReprGlobalId" description:"法人证件号码" required:"N"`
	ReservedMsg            string `json:"ReservedMsg" description:"保留域" required:"N"`
}

/*
返回数据
*/
type BindSmallAmountWithCheckCorp struct {
	BaseResp
}
