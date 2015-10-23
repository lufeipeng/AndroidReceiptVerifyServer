package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lufeipeng/AndroidReceiptVerifyServer/models"
)

type Android360Controller struct {
	beego.Controller
}

//Channel 360
//
func (this *Android360Controller) Post() {
	beego.Debug("Android360 is called");
	
	if(models.ProcAndroidCallback(this.Ctx.Input)) {
		this.Ctx.WriteString("ok");	
	} else {
		this.Ctx.WriteString("fail");
	}
}
