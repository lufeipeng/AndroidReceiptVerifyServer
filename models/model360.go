package models

import (
	"crypto/md5"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"sort"
	"strings"
	"encoding/hex"
	"time"
	"strconv"
)

var requireFields = []string{"app_key", "amount", "product_id", "app_uid",
	"user_id", "order_id", "sign_type", "gateway_flag", "app_order_id", "sign", "sign_return"}

func init() {

}

func calAndroid360Sign(input *context.BeegoInput) string {
	
	processedParams := make(map[string]string)
	for _, paramKey := range requireFields {
		if paramKey != "sign" && paramKey != "sign_return" {
			processedParams[paramKey] = input.Query(paramKey)
		}
	}
	keys := make([]string, len(processedParams))
	values := make([]string, len(processedParams))
	
	i := 0
    for k, _ := range processedParams {
        keys[i] = k
        i++
    }
    
	sort.Strings(keys)
	keyIndex := 0
	for _, key := range keys {
		if _, ok := processedParams[key]; ok {
			values[keyIndex] = processedParams[key];	
		}
		keyIndex++
	}
	signStr := strings.Join(values, "#") + "#" + iniconf.String("android360::APPSECRET")
	
	beego.Debug("calAndroid360Sign cal sign str is " + signStr);
		
	hasher := md5.New()
    hasher.Write([]byte(signStr))
    return hex.EncodeToString(hasher.Sum(nil))
}

func ProcAndroidCallback(input *context.BeegoInput) bool {
	if(ValidBillCallback(input)) {
		//TODO insert to db
		 o := orm.NewOrm()
		 OrderId := input.Query("app_order_id");
		 ExtraOrderId := input.Query("order_id");
		 
		 var PlatformType int
		 var Cost int;
		 var Uid int;
		 if platformType, ok := iniconf.Int("TYPE"); ok == nil {
		 	PlatformType = platformType
		 }
		 AccountId := input.Query("account_id");
		 
		 if value, ok := strconv.Atoi( input.Query("amount") ) ; ok == nil {
		 	Cost =	value;
		 }
		 
		 if value, ok := strconv.Atoi( input.Query("user_id") ) ; ok == nil {
		 	Uid = value
		 }
		 StoreItemId := input.Query("product_id");
		 Status := BILL_STATUS_FINISH;
		 InsertTime := time.Now();
		 r := o.Raw("Insert Ignore into billinghistory values(?,?,?,?,?,?,?,?,?)", OrderId, PlatformType, AccountId, 
		 	ExtraOrderId, Uid, StoreItemId, Cost, Status, InsertTime);
		 _, err := r.Exec();
		 return err == nil;
	}
	return false;
}


func ValidBillCallback(input *context.BeegoInput) bool {
		
	if input.Query("app_key") != iniconf.String("android360::APPKEY") {
		beego.Error("app_key is not matched")
		return false
	}

	for _, paramKey := range requireFields {
		if input.Query(paramKey) == "" {
			beego.Error("param:" + paramKey + " is empty")
			return false
		}
	}
	calMd5 := calAndroid360Sign(input)
	beego.Debug("calMd5 is" + calMd5 + " param md5 is " + input.Query("sign"))

	if calMd5 != input.Query("sign") {
		return false
	}
	return true
}
