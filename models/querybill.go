package models

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type BillStatus struct {
	OrderId string `json:"oi"`
	Status int `json:"s"`
	Cost int `json:"c"`
}

func QueryBill(orderId, platform string) string {
	var billStatus BillStatus;
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("order_id", "status", "cost").From("billinghistory").Where("order_id = ? and platform_type = ?");
	sql := qb.String();
	
	beego.Error("query bill sql is" + sql );
	o := orm.NewOrm();
	o.Raw(sql, orderId, platform).QueryRow(&billStatus);
	
	beego.Error("query bill result is" + billStatus.OrderId );
	
	b, err := json.Marshal(billStatus);
	if err != nil {
		beego.Error("parse to json is failed");
	}
	return string(b);
}