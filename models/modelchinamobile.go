package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//"github.com/astaxie/beego/orm"
	//"sort"
	//"strings"
	//"encoding/hex"
	//"time"
	//"strconv"
	"encoding/xml"
)

func init() {
}

var loginRequireFields = []string{"userId", "key", "cpId", "cpServiceId",
	"channelId", "p", "Ua"}


var billRequireFields = []string{"userId", "contentId", "consumeCode", "cpId",
	"hRet", "status", "versionId", "cpparam"}

type BillRequest struct {
    Request    xml.Name `xml:"request"`
    UserId string   `xml:"userId"`
    ContentId   string   `xml:"contentId"`
    ConsumeCode   string   `xml:"consumeCode"`
    CpId   string   `xml:"cpid"`
    HRet   string   `xml:"hRet"`
    Status   string   `xml:"status"`
    VersionId   string   `xml:"versionId"`
    Cpparam   string   `xml:"cpparam"`
    PackageId   string   `xml:"packageID"`
}

func ProcChinaMobileBillingCallback(input *context.BeegoInput) bool {
    request := BillRequest{}
    data := input.CopyBody();
    err := xml.Unmarshal(data, &request)
    if err != nil {
        beego.Error("ProcChinaMobileBillingCallback para xml is error", data)
        return false;
    }
    beego.Debug("ProcChinaMobileBillingCallback para xml is ", string(data[:]))
    //TODO
	return true;
}

func ValidLoginCallback(input *context.BeegoInput) bool {
	for _, paramKey := range loginRequireFields {
		value := input.Query(paramKey)
		if  value == "" {
			beego.Error("param:" + paramKey + " is empty")
			return false
		}else{
			beego.Debug("ValidLoginCallback,key:" + paramKey + " value:" + value)
		}
	}
	beego.Debug("ValidLoginCallback,success")
	return true
}

func ProcChinaMobileLoginCallback(input *context.BeegoInput) bool {
	if(ValidLoginCallback(input)) {
		return true;
	}
	return false;
}