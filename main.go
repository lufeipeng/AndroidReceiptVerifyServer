package main

import (
	_ "github.com/lufeipeng/AndroidReceiptVerifyServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}