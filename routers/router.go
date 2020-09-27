package routers

import (
	"beego-clothing/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/store", &controllers.StoreController{})
	beego.AutoRouter(&controllers.StoreController{})

}
