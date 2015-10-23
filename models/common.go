package models

import (
	"time"
	"os"
	"path/filepath"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
)

var iniconf config.ConfigContainer

func init() {
	var err error
	currentPath, err := os.Getwd()
	if err != nil {
		beego.Debug("Get Current Path failed")
	}
	confPath := filepath.Join(currentPath, "conf", "channel.conf")
	beego.Debug("channel.conf path is " + confPath)

	iniconf, err = config.NewConfig("ini", confPath)
	if err != nil {
		beego.Error("channel.conf init is failed")
	}
	
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	
	// set default database	
    orm.RegisterDataBase("default", "mysql", iniconf.String("BILLDB"), 30)
    
    // register model
  	//  orm.RegisterModel(new(BillingHistory2))

    // create table
    //orm.RunSyncdb("default", false, true)
}
    
type BillingHistory2 struct {
	OrderId   string `orm:"size(128);column(order_id)"`
    PlatformType int
    AccountId string `orm:"size(80)"`
    ExtraOrderId string `orm:"size(128)"`
    Uid int
    StoreItemId string `orm:"size(32)"`
    Cost int
    Status int
    InsertTime time.Time
}

const (
	BILL_STATUS_NEW = 1
	BILL_STATUS_RETRY = 2
	BILL_STATUS_FINISH = 3
	BILL_STATUS_INVALID = 4
	BILL_STATUS_COSTNOTMATCH = 5
	BILL_STATUS_RETRYOVERTIME = 6
)

type QueryBillStatus struct {
	oi string		//orderid
	s int			//status
	c int			//cost
}