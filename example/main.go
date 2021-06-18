package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/z-vip/pingan_sdk"
	"pingan_sdk/pkg"
	"time"
)

func main() {
	conf := pingan_sdk.Config{
		"app_id":               "d908fb6b-1e7e-45cf-b720-0f5c4b98589a",
		"app_type":             "PKCS12",
		"api_url":              "https://my-st1.orangebank.com.cn:567/api/group",
		"public_url":           "https://my-st1.orangebank.com.cn:567/api/approveDev",
		"public_key_path":      "../cert/publickey.cer",
		"private_key_path":     "../cert/2000908886@39.pfx",
		"private_key_pwd":      "1",
		"user_min_name":        "X67917",
		"password":             "${3DES}SHgd2RIwB2BFSPzkaES3Hg=",
		"mrch_code":            "3620",
		"ecif":                 "620104260170",
		"fund_summary_acct_no": "15000099027617",
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

	//6000 开户
	var userId = 2005
	//var params = pingan_sdk.Params{
	//	"FunctionFlag":      "1",
	//	"FundSummaryAcctNo": "15000099027617",
	//	"ChangeParam":       "SH",
	//	"TranNetMemberCode": fmt.Sprintf("Z%014d", userId),
	//}
	//fmt.Println(params)
	/*var args = pkg.ArgsOpenCustAcctId{
		FunctionFlag:      "1",
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		MemberProperty:    "SH",
		TranNetMemberCode: fmt.Sprintf("Z%014d", userId),
	}

	res, err := app.OpenCustAcctId(args)
	fmt.Println(err, ";;;;;", res)*/

	//6238 绑卡
	/*var args1 = pkg.ArgsBindUnionPayWithCheckCorp{
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
	var arg2 = pkg.ArgsCheckMsgCodeWithCorp{
		FundSummaryAcctNo: app.FundSummaryAcctNo,
		SubAcctNo:         "3620000000151087",
		TranNetMemberCode: fmt.Sprintf("Z%014d", userId),
		MemberAcctNo:      "6222023602087840336",
		MessageCheckCode:  "1234",
	}
	res2, err := app.CheckMsgCodeWithCorp(arg2)
	fmt.Println(err, ";;;;;", res2)

}
