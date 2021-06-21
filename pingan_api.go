package pingan_sdk

import (
	"github.com/z-vip/pingan_sdk/pkg"
)

/**
常用到的API
*/
var serverId string

//KFEJZB6000	会员子账户开立
func (a *App) OpenCustAcctId(params interface{}) (*pkg.OpenCustAcctId, error) {
	serverId = "OpenCustAcctId"
	var output = &pkg.OpenCustAcctId{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6238	会员绑定提现账户银联鉴权-校验法人
func (a *App) BindUnionPayWithCheckCorp(params interface{}) (*pkg.BindUnionPayWithCheckCorp, error) {
	serverId = "BindUnionPayWithCheckCorp"
	var output = &pkg.BindUnionPayWithCheckCorp{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6239	银联鉴权回填短信码-校验法人
func (a *App) CheckMsgCodeWithCorp(params interface{}) (*pkg.CheckMsgCodeWithCorp, error) {
	serverId = "CheckMsgCodeWithCorp"
	var output = &pkg.CheckMsgCodeWithCorp{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6240	会员绑定提现账户小额鉴权-校验法人
func (a *App) BindSmallAmountWithCheckCorp(params interface{}) (*pkg.BindSmallAmountWithCheckCorp, error) {
	serverId = "BindSmallAmountWithCheckCorp"
	var output = &pkg.BindSmallAmountWithCheckCorp{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6241	小额鉴权回填金额-校验法人
func (a *App) CheckAmountWithCorp(params interface{}) (*pkg.CheckAmountWithCorp, error) {
	serverId = "CheckAmountWithCorp"
	var output = &pkg.CheckAmountWithCorp{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6034	会员间交易-不验证
func (a *App) MemberTransaction(params interface{}) (*pkg.MemberTransaction, error) {
	serverId = "MemberTransaction"
	var output = &pkg.MemberTransaction{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6033	会员提现-不验证
func (a *App) MembershipWithdrawCash(params interface{}) (*pkg.MembershipWithdrawCash, error) {
	serverId = "MembershipWithdrawCash"
	var output = &pkg.MembershipWithdrawCash{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6007	会员资金冻结-不验证
func (a *App) MembershipTrancheFreeze(params interface{}) (*pkg.MembershipTrancheFreeze, error) {
	serverId = "MembershipTrancheFreeze"
	var output = &pkg.MembershipTrancheFreeze{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6163	会员资金支付-不验证
func (a *App) MembershipTranchePay(params interface{}) (*pkg.MembershipTranchePay, error) {
	serverId = "MembershipTranchePay"
	var output = &pkg.MembershipTranchePay{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6050	查询普通转账充值明细
func (a *App) CommonTransferRechargeQuery(params interface{}) (*pkg.CommonTransferRechargeQuery, error) {
	serverId = "CommonTransferRechargeQuery"
	var output = &pkg.CommonTransferRechargeQuery{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6103	查询对账文件信息
func (a *App) ReconciliationDocumentQuery(params interface{}) (*pkg.ReconciliationDocumentQuery, error) {
	serverId = "ReconciliationDocumentQuery"
	var output = &pkg.ReconciliationDocumentQuery{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6011	查询资金汇总账户余额
func (a *App) SupAcctIdBalanceQuery(params interface{}) (*pkg.SupAcctIdBalanceQuery, error) {
	serverId = "SupAcctIdBalanceQuery"
	var output = &pkg.SupAcctIdBalanceQuery{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6092	根据会员代码查询会员子账号
func (a *App) QueryCustAcctIdByThirdCustId(params interface{}) (*pkg.QueryCustAcctIdByThirdCustId, error) {
	serverId = "QueryCustAcctIdByThirdCustId"
	var output = &pkg.QueryCustAcctIdByThirdCustId{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6093	查询会员子账号余额
func (a *App) QueryCustAcctIdBalance(params interface{}) (*pkg.QueryCustAcctIdBalance, error) {
	serverId = "QueryCustAcctIdBalance"
	var output = &pkg.QueryCustAcctIdBalance{}
	err := a.Execute(serverId, params, &output)
	return output, err
}

//KFEJZB6037	查询会员子账号
func (a *App) QueryCustAcctId(params interface{}) (*pkg.QueryCustAcctId, error) {
	serverId = "QueryCustAcctId"
	var output = &pkg.QueryCustAcctId{}
	err := a.Execute(serverId, params, &output)
	return output, err
}
