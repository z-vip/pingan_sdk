package pingan_sdk

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/z-vip/pingan_sdk/pkg"
	"github.com/z-vip/pingan_sdk/util"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"time"
)

const Version = "1.3.7"

type Config map[string]string      //配置
type Params map[string]interface{} //参数
type Result struct {
	ErrCode string      `json:"errorCode"`
	ErrMsg  string      `json:"errorMsg"`
	Data    interface{} `json:"data"`
}

//app基础信息（配置信息）
type AppBase struct {
	AppId             string `json:"app_id" description:"应用ID" required:"true"`
	AppType           string `json:"app_type" description:"应用类型" required:"true"`
	PrivateKeyPath    string `json:"private_key_path" description:"私钥地址" required:"true"`
	PrivateKeyPwd     string `json:"private_key_pwd" description:"私钥密码" required:"true"`
	PublicKeyPath     string `json:"public_key_path" description:"公钥地址" required:"true"`
	PublicUrl         string `json:"public_url" description:"公共地址" required:"true"`
	ApiUrl            string `json:"api_url" description:"API地址" required:"true"`
	ApiVersion        string `json:"app_version" description:"版本号" required:"N"`
	Ecif              string `json:"ecif" description:"ECIF号" required:"true"`
	FundSummaryAcctNo string `json:"fund_summary_acct_no" description:"资金汇总账号" required:"N"`
	MrchCode          string `json:"mrch_code" description:"商户号" required:"N"`
	UserMinName       string `json:"user_min_name" description:"用户短号" required:"N"`
	UserPwd           string `json:"UserPwd" description:"用有密码" required:"N"`
}

type App struct {
	AppBase
	AccessToken string            `json:"access_token" description:"Token验证" required:"N"`
	PrivateCert *x509.Certificate `json:"private_cert" description:"私钥证书" required:"N"`
	PrivateKey  *rsa.PrivateKey   `json:"private_key" description:"私钥" required:"N"`
	PublicCert  *x509.Certificate `json:"public_cert" description:"公钥证书" required:"N"`
	PublicKey   *rsa.PublicKey    `json:"public_key" description:"公钥" required:"N"`
	Redis       *redis.Pool       `json:"redis" description:"redis连接池" required:"N"`
}

/**
初始化app参数是否完整
*/
func (a *App) AppInit() error {
	//检查配置是否定义
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		//非必要项不检查
		if required := v.Type().Field(i).Tag.Get("required"); required == "false" || required == "N" {
			continue
		}
		fieldTag := v.Type().Field(i).Tag.Get("description")
		if fieldValue := fmt.Sprintf("%v", v.Field(i)); fieldValue == "" {
			return errors.New(fieldTag + "未定义")
		}
	}
	if a.ApiVersion == "" {
		a.ApiVersion = Version
	}
	//处理私钥证书
	bytes, err := ioutil.ReadFile(a.PrivateKeyPath)
	if err != nil {
		return errors.New("无法读取私钥文件")
	}
	key, cert, err := pkcs12.Decode(bytes, a.PrivateKeyPwd)
	if err != nil {
		return errors.New("私钥解密失败")
	}
	a.PrivateKey = key.(*rsa.PrivateKey)
	a.PrivateCert = cert
	//fmt.Println(util.JsonEncode(a.PrivateCert))
	//处理公钥
	bytes, err = ioutil.ReadFile(a.PublicKeyPath)
	if err != nil {
		return errors.New("无法读取公钥文件")
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return errors.New("公钥解析失败")
	}
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	a.PublicCert = cert
	//fmt.Println("------------", reflect.TypeOf(cert.PublicKey).String())
	a.PublicKey = cert.PublicKey.(*rsa.PublicKey)
	//fmt.Println(util.JsonEncode(a.PublicCert), "++++", util.JsonEncode(a.PublicKey))
	return nil
}

/**
生成PK，证书的字符串
*/
func (a *App) CreatePK() string {
	if a.PrivateCert == nil {
		return ""
	}
	str := util.JsonEncode(a.PrivateCert.RawSubjectPublicKeyInfo)
	str = strings.ReplaceAll(str, "\"", "")
	return util.ReplaceSpecialChar(str)
}

/**
生成DN信息
*/
func (a *App) CreateDN() string {
	if a.PrivateCert == nil {
		return ""
	}
	str := fmt.Sprintf("CN=%s,OU=%s,OU=%s,O=%s,C=%s", a.PrivateCert.Subject.CommonName, a.PrivateCert.Subject.OrganizationalUnit[1], a.PrivateCert.Subject.OrganizationalUnit[0],
		a.PrivateCert.Subject.Organization[0], a.PrivateCert.Subject.Country[0])
	return util.ReplaceSpecialChar(str)
}

/**
生成签名，token的特殊,请传args[0]:true
*/
func (a *App) CreateSign(params Params, args ...bool) error {
	var keys []string
	var isToken bool
	if len(args) > 0 && args[0] == true {
		isToken = true
	}
	for k, v := range params {
		//空值不保留
		if fmt.Sprintf("%v", v) == "" {
			delete(params, k)
			continue
		}
		if isToken {
			//token只保留有用的 key
			if _, ok := util.InArray(k, []string{"ApplicationID", "RandomNumber", "SDKType", "CerSeqNo", "CertPeriod", "Mac"}); !ok {
				continue
			}
		}
		keys = append(keys, k) //参与签名的key
	}
	sort.Strings(keys) //key排序
	var str string
	for _, v := range keys {
		str += fmt.Sprintf("%s=%v&", v, params[v]) //组合加密字符串
	}
	//fmt.Println("str:", str)
	hashMd5 := md5.Sum([]byte(str))
	hashed := hashMd5[:]

	signByte, err := a.PrivateKey.Sign(rand.Reader, hashed, crypto.MD5) //使用MD5生成签名
	//signature, err := rsa.SignPKCS1v15(rand.Reader, a.PrivateKey, crypto.MD5, hashed)
	if err != nil {
		return err
	}
	//签名赋值到参数中
	sign := base64.StdEncoding.EncodeToString(signByte) //签名base64
	params["RsaSign"] = sign
	if isToken {
		params["RsaSign"] = util.ReplaceSpecialChar(sign)
	}
	return nil
}

/**
验签
*/
func (a *App) VerifySign(input interface{}) error {
	var params = Params{}
	_ = util.Convert(input, &params) //将数据转换到map
	//得到签名
	var sign string
	if v, ok := params["RsaSign"]; ok && reflect.TypeOf(v).String() == "string" {
		sign = v.(string)
	}
	if sign == "" {
		return nil //没有签名字段不处理
	}
	//fmt.Println("####sign:", sign)
	//处理数据生成 验签内容
	var keys []string
	for k, v := range params {
		//errCode,errorMsg,RsaSign不参与 内容校验
		if _, ok := util.InArray(k, []string{"errorCode", "errorMsg", "RsaSign"}); !ok {
			//值为数组，取最后一组值 todo

			//值为空的不参与 内容校验
			if str := fmt.Sprintf("%v", v); str != "" {
				keys = append(keys, k)
			}
		}
	}
	sort.Strings(keys) //排序
	var str string
	for _, k := range keys {
		str += fmt.Sprintf("%s=%v&", k, params[k]) //组合生成 校验字符串
	}
	//fmt.Println("####str:", str)
	if str == "" {
		return nil //没有验签的内容，不处理
	}
	//验签内容生成md5
	hashMd5 := md5.Sum([]byte(str))
	hashed := hashMd5[:]
	//签名base64解码
	sig, _ := base64.StdEncoding.DecodeString(sign)
	//md5校验签名
	err := rsa.VerifyPKCS1v15(a.PublicKey, crypto.MD5, hashed, sig)
	//return rsa.VerifyPSS(pub, hashType, signed, sig, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
	//err = a.PublicCert.CheckSignature(x509.MD5WithRSA, []byte(str), sig)
	if err != nil {
		return errors.New("签名验证效失败")
	}
	return nil
}

/**
获取token
*/
func (a *App) GetToken(args ...bool) error {
	handler := a.Redis.Get()
	defer handler.Close()
	var refreshToken bool
	var cacheKey = "pingan_token_" + a.AppId
	if len(args) > 0 && args[0] == true {
		refreshToken = true
	}
	//不强制刷新token, 从缓存取token
	if !refreshToken {
		data, err := handler.Do("GET", cacheKey)
		if err == nil && data != nil {
			a.AccessToken = string(data.([]byte))
			return nil
		}
	}
	//从接口获取token
	var params = Params{}
	params["ApplicationID"] = a.AppId
	params["PK"] = a.CreatePK()
	params["DN"] = a.CreateDN()
	params["RandomNumber"] = util.RandomNumber(6)
	params["SDKType"] = "api"
	err := a.CreateSign(params, true)
	if err != nil {
		return err
	}
	//fmt.Println("++request:", util.JsonEncode(params))
	resp, err := a.Request(a.PublicUrl, params)
	if err != nil {
		return err
	}
	//toke返回数据结构体
	var appToken struct {
		ErrorCode      string `json:"errorCode"`
		ErrorMsg       string `json:"errorMsg"`
		RsaSign        string `json:"RsaSign"`
		AppAccessToken string `json:"appAccessToken"`
		ValidateDay    string `json:"ValidateDay"`
		UrlList        string `json:"urlList"`
	}
	_ = json.Unmarshal(resp, &appToken)
	fmt.Println("===get Token:===", util.JsonEncode(appToken))
	if appToken.ErrorCode == "OPEN-E-000000" {
		//验证签名
		err = a.VerifySign(appToken)
		if err != nil {
			return err
		}
		//token处理
		if appToken.AppAccessToken != "" {
			a.AccessToken = appToken.AppAccessToken
			//写缓存
			_, err = handler.Do("SETEX", cacheKey, 3600*5, a.AccessToken)
			return err
		}
	}
	if appToken.ErrorMsg != "" {
		return errors.New(appToken.ErrorCode + " " + appToken.ErrorMsg)
	} else {
		return errors.New("生成Token失败")
	}
}

/**
HTTP请求
*/
func (a *App) Request(uri string, params Params) ([]byte, error) {
	data := util.JsonEncode(params)
	header := Config{
		"Content-Type": "application/json",
	}
	resp, err := util.HttpRequest("POST", uri, data, header)
	return resp, err
}

/**
请求API接口
*/
func (a *App) Execute(serverId string, params interface{}, output interface{}) (err error) {
	//获取token
	err = a.GetToken()
	if err != nil {
		return
	}
	//公共参数
	var baseArgs = pkg.BaseArgs{
		ApiVersionNo:   a.ApiVersion,
		AppAccessToken: a.AccessToken,
		ApplicationID:  a.AppId,
		RequestMode:    "json",
		SDKType:        "api",
		SdkSeid:        util.RandomNumber(6) + "-pab" + a.MrchCode,
		SdkVersionNo:   a.ApiVersion,
		TranStatus:     "0",
		TxnTime:        time.Now().Format("20060102150405") + "000",
		ValidTerm:      time.Now().Format("20060102"),
		TxnClientNo:    a.Ecif,              //商户ecif_code
		TxnCode:        serverId,            //接口code
		CnsmrSeqNo:     util.GenerateCode(), //交易流水号
		MrchCode:       a.MrchCode,          //商户号
	}
	var input = Params{}
	_ = util.Convert(baseArgs, &input) //base值并入args
	_ = util.Convert(params, &input)   //传参值并入args
	//fmt.Println("===input:", util.JsonEncode(input))
	err = a.CreateSign(input) //生成签名
	if err != nil {
		return
	}
	var uri = a.ApiUrl + "/" + serverId
	body, err := a.Request(uri, input)
	//fmt.Println("===body:", string(body))
	if err != nil {
		return
	}
	var respBase pkg.BaseResp
	_ = json.Unmarshal(body, &respBase)
	//处理接口基本信息
	if respBase.TxnReturnCode != "000000" {
		return fmt.Errorf("[%s](%s)%s", respBase.CnsmrSeqNo, respBase.TxnReturnCode, respBase.TxnReturnMsg)
	}
	if output == nil {
		//如果不需要定义输出对象，直接使用base
		output = respBase
	} else {
		//反解析JSON到输出内容
		err = json.Unmarshal(body, &output)
	}
	//验证签名
	if respBase.RsaSign != "" {
		err = a.VerifySign(output)
	}
	return
}
