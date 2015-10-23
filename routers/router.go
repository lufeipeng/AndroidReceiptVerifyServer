package routers

import (
	"github.com/lufeipeng/AndroidReceiptVerifyServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/android360", &controllers.Android360Controller{})
    beego.Router("/query", &controllers.QueryBillController{})
}
