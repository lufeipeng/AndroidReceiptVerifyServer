package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("This is android receipt verify server");
	c.Ctx.WriteString("360 callback url: http://localhost:8080/android360");
	c.Ctx.WriteString("query bill url: http://localhost:8080/query");	
}
