package crossutting

import (
	"../ddd_interfaces"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	//logs.SetLogger(logs.AdapterConsole, `{"level":8,"color":true}`)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"kendoddd.log","level":8,"maxsize":5,"daily":true,"maxdays":10}`)
	beego.BeeLogger.SetLogger(logs.AdapterFile, `{"filename":"kendoddd.log","level":5,"maxsize":2097152,"daily":true,"maxdays":10}`)

	//ORM注册Model
	orm.RegisterModel(new(ddd_interfaces.ProtocolOrm))

	//ORM注册驱动和数据库地址
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "assets/kendopay.db")

	//仓储层初始化
	RepoFac.Init()
}
