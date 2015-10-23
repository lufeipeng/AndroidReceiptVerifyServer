本工程用于Android 360App计费回调
Web框架采用beego, 地址在http://beego.me/

1： 初始化配置
	A: 创建数据库, 例如billdb
	B: 创建数据库表格， sql语句: sql/billinghistory.sql
	C: 修改channel.conf中的BILLDB的配置
	D: 修改产品对应的APPID, APPKEY, APPSECRET

2: 编译并运行
	A: bee pack
	B: 把生成的AndroidReceiptVerifyServer.tar.gz放到服务器上运行即可
	C: 也可以直接go run main.go

3: Q&A
	QQ: 395088998
	Email: lufeipeng@163.com