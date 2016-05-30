package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lufeipeng/AndroidReceiptVerifyServer/models"
	"encoding/xml"
)

type ChinaMobileBillingController struct {
	beego.Controller
}

//Channel mobile
func (this *ChinaMobileBillingController) Post() {
	beego.Debug("ChinaMobile is called");
	
	resultStr := models.ProcChinaMobileBillingCallback(this.Ctx.Input);
	this.Ctx.WriteString(xml.Header);	
	this.Ctx.WriteString(resultStr);
}
