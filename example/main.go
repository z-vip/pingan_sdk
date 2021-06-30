package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/z-vip/pingan_sdk"
	"github.com/z-vip/pingan_sdk/pingan_pkg"
	"time"
)

func main() {
	var str = "1234567890"
	fmt.Println(str[3 : len(str)-1])

	conf := pingan_sdk.Config{
		"app_id":               "",
		"app_type":             "",
		"api_url":              "",
		"public_url":           "",
		"public_key_path":      "",
		"private_key_path":     "",
		"private_key_pwd":      "",
		"ecif":                 "",
		"fund_summary_acct_no": "",
		"user_min_name":        "",
		"user_pwd":             "",
		"mrch_code":            "",
	}

	var myRedis = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   5000,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			_, err = c.Do("SELECT", 0)
			if err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	//fmt.Println(myRedis)
	app, err := pingan_sdk.NewApp(conf, myRedis)
	//fmt.Println("@@", app, err)
	//6000 开户
	var userId = 2005
	fmt.Println("UserId:", userId)

	//var params = pingan_sdk.Params{
	//	"FunctionFlag":      "1",
	//	"FundSummaryAcctNo": "15000099027617",
	//	"ChangeParam":       "SH",
	//	"TranNetMemberCode": fmt.Sprintf("Z%014d", userId),
	//}
	//fmt.Println(params)
	/*var args = pingan_pkg.ArgsOpenCustAcctId{
		FunctionFlag:      "1",
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		MemberProperty:    "SH",
		TranNetMemberCode: fmt.Sprintf("Z%014d", userId),
	}

	res, err := app.OpenCustAcctId(args)
	fmt.Println(err, ";;;;;", res)*/

	//6238 绑卡
	/*var args1 = pingan_pkg.ArgsBindUnionPayWithCheckCorp{
		FundSummaryAcctNo:  app.FundSummaryAcctNo,
		SubAcctNo:          "3620000000151087",
		TranNetMemberCode:  "Z00000000002005",
		MemberName:         "刘春晓",
		MemberGlobalType:   "1",
		MemberGlobalId:     "37063019800706501X",
		MemberAcctNo:       "6222023602087840336",
		BankType:           "2",
		AcctOpenBranchName: "中国工商银行",
		EiconBankBranchId:  "102100099996",
		Mobile:             "15820213936",
		IndivBusinessFlag:  "2",
	}
	res1, err := app.BindUnionPayWithCheckCorp(args1)
	fmt.Println(err, ";;;;;", res1)*/

	//6239
	/*var arg2 = pingan_pkg.ArgsCheckMsgCodeWithCorp{
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		SubAcctNo:         "3620000000151087",
		TranNetMemberCode: fmt.Sprintf("Z%014d", userId),
		MemberAcctNo:      "6222023602087840336",
		MessageCheckCode:  "1234",
	}
	res2, err := app.CheckMsgCodeWithCorp(arg2)
	fmt.Println(err, ";;;;;", res2)*/

	//余额
	/*var arg3 = pingan_pkg.ArgsQueryCustAcctId{
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		TranNetMemberCode: fmt.Sprintf("Z%014d", userId),
	}
	res3, err := app.QueryCustAcctId(arg3)
	fmt.Println(err, ";;;;;", res3)*/

	//对账

	var arg4 = pingan_pkg.ArgsReconciliationDocumentQuery{
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		FileType:          pingan_pkg.FileTypeCZ,
		FileDate:          time.Now().AddDate(0, 0, -1).Format("20060102"),
	}
	res4, err := app.ReconciliationDocumentQuery(arg4)
	fmt.Println(arg4, "=", err, "----", res4)

}
