package main

import (
	_ "github.com/lufeipeng/AndroidReceiptVerifyServer/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
    "os"
    "path/filepath"
)

func TestFunct(){
	var err error
	currentPath, err := os.Getwd();
	beego.Info("TestFunct", currentPath);
	
	if err != nil {
		beego.Debug("Get Current Path failed")
	}
	confPath := filepath.Join(currentPath, "conf", "channel.conf")
	beego.Debug("channel.conf path is " + confPath)

	iniconf, err := config.NewConfig("ini", confPath)
	if err != nil {
		beego.Error("channel.conf init is failed")
	}
	beego.Trace("iniConf str is", iniconf.String("BILLDB"));
}

func main() {
	beego.SetLogger("file", `{"filename":"billserver.log"}`)
	beego.SetLevel(beego.LevelDebug);
	//TestFunct();
	beego.Run();
}