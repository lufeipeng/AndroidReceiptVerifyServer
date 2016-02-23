package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lufeipeng/AndroidReceiptVerifyServer/models"
)

type ChinaMobileBillingController struct {
	beego.Controller
}

//Channel mobile
func (this *ChinaMobileBillingController) Post() {
	beego.Debug("ChinaMobile is called");
	
	if(models.ProcChinaMobileBillingCallback(this.Ctx.Input)) {
		this.Ctx.WriteString("ok");	
	} else {
		this.Ctx.WriteString("fail");
	}
}
