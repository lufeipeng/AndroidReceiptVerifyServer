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
	"io/ioutil"
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

type BillResponse struct {
	XMLName xml.Name `xml:"response"`
	Hret    string `xml:"hRet"`
	Message string `xml:"message"`
}

func ProcChinaMobileBillingCallback(input *context.BeegoInput) string {
	var billResponse BillResponse
	billResponse.Hret  = "1";
	billResponse.Message  = "failure";
		
    request := BillRequest{}
    data := input.CopyBody(1024);
    err := xml.Unmarshal(data, &request)
    if err != nil {
        beego.Error("ProcChinaMobileBillingCallback para xml is error", string(data[:]))
        resultBytes, _ := xml.MarshalIndent(billResponse, "  ", "  ");
		return string(resultBytes);
    }
    billResponse.Hret  = "0";
	billResponse.Message  = "successful";
	
    beego.Debug("ProcChinaMobileBillingCallback para xml is ", string(data[:]))
	resultBytes, _ := xml.MarshalIndent(billResponse, "  ", "  ");
	return string(resultBytes);
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
	resultBytes, _ := ioutil.ReadAll(input.Context.Request.Body)
	beego.Debug("ValidLoginCallback,success", string(resultBytes));
	return true
}

func ProcChinaMobileLoginCallback(input *context.BeegoInput) bool {
	if(ValidLoginCallback(input)) {
		return true;
	}
	return false;
}