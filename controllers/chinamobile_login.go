package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lufeipeng/AndroidReceiptVerifyServer/models"
)

type ChinaMobileLoginController struct {
	beego.Controller
}

//Channel mobile Login
//
func (this *ChinaMobileLoginController) Get() {
	beego.Debug("ChinaMobile login is called");
	
	if(models.ProcChinaMobileLoginCallback(this.Ctx.Input)) {
		this.Ctx.WriteString("1");	
	} else {
		this.Ctx.WriteString("0");
	}
}
