package util

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/*
转成字符串
*/
func Interface2String(val interface{}) string {
	fv := reflect.ValueOf(val)
	kind := fv.Kind()
	if kind == reflect.String {
		return fv.String()
	} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		return strconv.FormatInt(fv.Int(), 10)
	} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
		return strconv.FormatUint(fv.Uint(), 10)
	} else if kind == reflect.Float32 || kind == reflect.Float64 {
		return strconv.FormatFloat(fv.Float(), 'f', -1, 64)
	} else if kind == reflect.Slice || kind == reflect.Map {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		_ = encoder.Encode(val)
		bytes := buffer.Bytes()
		return string(bytes[0 : len(bytes)-1])
	} else {
		return ""
	}
}

// 将字符串转成值
func String2Interface(str string, args ...string) (value interface{}) {
	if len(args) == 0 {
		return str
	}
	kind := strings.ToLower(args[0])

	switch kind {
	case "string":
		if len(args) > 1 && args[1] != "" {
			r, _ := regexp.Compile(`%([dfsv]{1})`)
			if r.MatchString(args[1]) {
				value = fmt.Sprintf(args[1], str)
			} else {
				value = args[1]
			}
		} else {
			value = str
		}
	case "int":
		value, _ = strconv.Atoi(str)
	case "int8", "int16", "int32", "int64":
		bit, _ := strconv.Atoi(kind[3:])
		value, _ = strconv.ParseInt(str, 10, bit)
	case "uint8", "uint16", "uint32", "uint64":
		bit, _ := strconv.Atoi(kind[4:])
		value, _ = strconv.ParseUint(str, 10, bit)
	case "float", "float32", "float64":
		bit, _ := strconv.Atoi(kind[5:])
		value, _ = strconv.ParseFloat(str, bit)
	case "slice", "array":
		value = []string{str}
	default:
		if kind == "" {
			value = str
		} else {
			value = args[0]
		}
	}
	// TODO support more types
	return
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//判断输出
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

//反解析像
func Convert(in interface{}, out interface{}) error {
	inByte, err := json.Marshal(in)
	//fmt.Println("json:", string(inByte))
	err = json.Unmarshal(inByte, &out)
	return err
}

//生成json string
func JsonEncode(in interface{}) string {
	if inByte, err := json.Marshal(in); err == nil {
		return string(inByte)
	}
	return ""
}

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(t))
		} else if t < 36 {
			result = append(result, string(t-10+65))
		} else {
			result = append(result, string(t-36+97))
		}
	}
	return strings.Join(result, "")
}

//生成随机数字
func RandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(result, "")
}

// GenerateCode 通用生成单据编号code
func GenerateCode(args ...string) string {
	var pre string //前辍
	if len(args) > 0 {
		pre = args[0]
	}
	day := time.Now().Format("060102")
	microTime := fmt.Sprintf("%v", time.Now().UnixNano()/10)[5:9]
	return fmt.Sprintf("%s%s%s%s", pre, day, microTime, RandomNumber(4))
}

//生成32位唯一字符串
func UniqueString() string {
	//用时间纳秒 + 随机10位数
	str := fmt.Sprintf("%d_%s", time.Now().UnixNano(), RandomString(10))
	return String2md5(str) //生成md5值返回
}

//判断某一个值是否含在切片之中
func InArray(val interface{}, arr interface{}) (index int, exists bool) {
	exists = false
	index = -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
func ReplaceSpecialChar(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "　", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "+", "%2B")
	str = strings.ReplaceAll(str, "=", "%3D")
	return str
}
