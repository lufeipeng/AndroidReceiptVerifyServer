package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lufeipeng/AndroidReceiptVerifyServer/models"
)

type QueryBillController struct {
	beego.Controller
}

func (c *QueryBillController) Post() {
	orderId := c.Ctx.Input.Query("orderId");
	platformType := c.Ctx.Input.Query("platformType");
	result := models.QueryBill(orderId, platformType);
	c.Ctx.WriteString(result);	
}